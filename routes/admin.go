package routes

import (
	"context"
	"log"
	"mctrialgo/models"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type adminHTMLtemplate struct {
	Title        string
	HospitalName string
	ImageTag     string
	Message      string
	Redirect     string
	CSS          string
}

// GetHospAndUser  Hosp ID, Usernameをcontextから得る (global varの代用)
func GetHospAndUser(c echo.Context) (hospid uint, username string) {
	hospid = c.Get("UserID").(uint)
	if hospid > 0 {
		username = c.Get("UserName").(string)
	}
	return hospid, username
}

// AdminRouter  GET "/admin" を処理
func AdminRouter(c echo.Context) error {
	db := Repository()
	defer db.Close()

	hospid, username := GetHospAndUser(c)
	if hospid != 1 {
		// invalid access, go to logout
		return c.Redirect(http.StatusFound, "/logout")
	}

	htmlvariable := adminHTMLtemplate{
		Title:        "adminコンソール",
		HospitalName: username,
		Redirect:     "",
		Message:      "",
		ImageTag:     "",
		CSS:          "/css/index.css",
	}

	return c.Render(http.StatusOK, "admin", htmlvariable)

}

var wg sync.WaitGroup
var (
	m = &sync.Map{}
)
var id string

// StatRouter ダウンロードの進捗を JSON で返す
func StatRouter(c echo.Context) error {
	ck, err := c.Cookie("download-progress")
	if err != nil {
		log.Println(err)
		return err
	}
	progress := 0
	v, ok := m.Load(ck.Value)
	if ok {
		if vi, ok := v.(int); ok {
			progress = vi
		}
	}
	return c.JSON(http.StatusOK, &struct {
		Progress int `json:"progress"`
	}{
		Progress: progress,
	})
}

// AdminAnalyzeRouter  GET "/admin/:func" を処理
func AdminAnalyzeRouter(c echo.Context) error {
	db := Repository()
	defer db.Close()
	hospid, username := GetHospAndUser(c)
	if hospid != 1 {
		// invalid access, go to logout
		return c.Redirect(http.StatusFound, "/logout")
	}

	htmlvariable := adminHTMLtemplate{
		Title:        "adminコンソール",
		HospitalName: username,
		Redirect:     "",
		Message:      "",
		CSS:          "/css/index.css",
	}
	sel, _ := strconv.Atoi(c.Param("func"))
	switch sel {
	case 1:
		htmlvariable.Message = "now calculating"
		id = uuid.New().String()
		htmlvariable.Redirect = `<meta http-equiv="refresh" content="0;URL='/admin/2'" />`
		c.SetCookie(&http.Cookie{
			Name:  "download-progress",
			Value: id,
			Path:  "/",
		})
		// cookieを保存させる
	case 2:
		// https://mattn.kaoriya.net/software/lang/go/20170622160723.htm を参考にしました

		htmlvariable.Message = "now calculating"
		htmlvariable.Redirect = `
		<script>
		window.addEventListener('load', function() {
		  var prog=function progress() {
			  fetch("/stat", {
				'credentials': "same-origin"
			  }).then(function(response) {
				return response.json();
			  }).then(function(json) {
				document.querySelector('#progress').textContent = "progress: "+json.progress+"%";
			  if (json.progress >= 100) {
				  location.href = "/admin/3";
				  clearInterval();
			  }
			})
		  };
		  setInterval(prog,1000);
		}, false);
		</script>
		` // 起動時から 1000msec間隔で/statし、progressに表示、>=100なら自動起動を中止
		// 非同期で走らせる
		m.Store(id, 0)
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = analyze()
			m.Store(id, 100) // progress 100%
		}()
	case 3:
		wg.Wait()
		m.Store(id, nil)
		htmlvariable.Message = "finished"
		htmlvariable.ImageTag = "<img src=/img/test.png><br>"

		//delete cookie
		c.SetCookie(&http.Cookie{
			Name:    "download-progress",
			Value:   id,
			Path:    "/",
			Expires: time.Now().Add(-1 * time.Hour),
		})
		/*
			case 4:
				htmlvariable.Message = "now testing cookie and realtime rendering... sleep 100sec"
				id = uuid.New().String()
				htmlvariable.Redirect = `<meta http-equiv="refresh" content="0;URL='/admin/4'" />`
				c.SetCookie(&http.Cookie{
					Name:  "download-progress",
					Value: id,
					Path:  "/",
				})
			case 5:
				// https://mattn.kaoriya.net/software/lang/go/20170622160723.htm を参考にしました
				htmlvariable.Redirect = `
				<script>
				window.addEventListener('load', function() {
					var prog=function progress() {
						fetch("/stat", {
						'credentials': "same-origin"
						}).then(function(response) {
						return response.json();
						}).then(function(json) {
						document.querySelector('#progress').textContent = "progress: "+json.progress+"%";
						if (json.progress >= 100) {
							location.href = "/admin/5";
							clearInterval();
						}
					})
					};
					setInterval(prog,1000);
				}, false);
				</script>
				`
				// 非同期で走らせる
				m.Store(id, 0)
				wg.Add(1)
				go func() {
					defer wg.Done()
					for i := 0; i < 10; i++ {
						time.Sleep(2 * time.Second)
						m.Store(id, i*10)
					}
					m.Store(id, 100)
				}()

			case 6:
				wg.Wait()
				m.Store(id, nil)
				htmlvariable.Message = "finished"
		//*/
	default:
		htmlvariable.Message = "エラー"
	}
	return c.Render(http.StatusOK, "admin", htmlvariable)

}

func analyze() string {
	db := Repository()
	defer db.Close()
	ctx := context.Background()
	boil.DebugMode = false

	patients, err := models.Patients().All(ctx, db)
	if err != nil {
		panic(err)
	}

	output := ""

	for _, p := range patients {
		// fmt.Printf("%s: ", p.Initial.String)
		events, err := models.Events(qm.Where("hospital_id=? AND serialid=?", p.HospitalID, p.Serialid)).All(ctx, db)
		if err == nil {
			dropoutflag := false
			macceflag := false
			deadflag := false
			for _, e := range events {
				if e.Dropout.Valid && e.Dropout.Bool {
					dropoutflag = true
					p.Dropdate = e.Date
					p.Update(ctx, db, boil.Infer()) // save change
					break
				}
				if e.Alive.Valid && e.Alive.Bool == false {
					deadflag = true
					p.Deaddate = e.Date
					p.Update(ctx, db, boil.Infer()) // save change
					break
				}
				if e.Macce.Valid && e.Macce.Bool {
					macceflag = true
					p.Dropdate = e.Date
					p.Update(ctx, db, boil.Infer()) // save change
					break
				}

			}
			if dropoutflag == false && deadflag == false && macceflag == false {
				p.Finishdate = null.TimeFrom(time.Now())
				p.Update(ctx, db, boil.Infer()) // save change
			}

		}
	}
	m.Store(id, 50) // progress 50%
	// 強引にRを動かし解析
	var rscript string
	if rscript = os.Getenv("RSCRIPT"); rscript == "" {
		rscript = "/usr/local/bin/Rscript"
	}
	out, err := exec.Command(rscript, "analysis.R").CombinedOutput()
	if err != nil {
		output += err.Error()
	} else {
		output += string(out)
	}
	return output
}

// AdminRouterPost  POST "/admin" を処理
func AdminRouterPost(c echo.Context) error {
	db := Repository()
	defer db.Close()

	hospid, username := GetHospAndUser(c)
	if hospid != 1 {
		// invalid access, go to logout
		return c.Redirect(http.StatusFound, "/logout")
	}

	htmlvariable := adminHTMLtemplate{
		Title:        "adminコンソール",
		HospitalName: username,
		ImageTag:     "",
		Message:      "",
		Redirect:     "",
		CSS:          "/css/index.css",
	}

	return c.Render(http.StatusOK, "admin", htmlvariable)
}
