package routes

import (
	"context"
	"mctrialgo/models"
	"net/http"
	"os/exec"
	"strconv"
	"time"

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
		Message:      "",
		ImageTag:     "",
		CSS:          "/css/index.css",
	}

	return c.Render(http.StatusOK, "admin", htmlvariable)

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
		Message:      "",
		CSS:          "/css/index.css",
	}
	sel, _ := strconv.Atoi(c.Param("func"))
	switch sel {
	case 1:
		htmlvariable.Message = analyze()
		htmlvariable.ImageTag = "<img src=/img/test.png><br>"
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
	// 強引にRを動かし解析
	out, err := exec.Command("/usr/local/bin/Rscript", "analysis.R").CombinedOutput()
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
		CSS:          "/css/index.css",
	}

	return c.Render(http.StatusOK, "admin", htmlvariable)
}
