package routes

import (
	"context"
	"mctrialgo/models"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo/v4"
)

type loginHTMLtemplate struct {
	Title  string
	NoUser string
	CSS    string
}

// LoginRouter  GET "/login" を処理
func LoginRouter(c echo.Context) error {

	htmlvariable := loginHTMLtemplate{
		Title:  "多施設臨床トライアルシステム ログイン",
		NoUser: "",
		CSS:    "/css/login.css",
	}
	return c.Render(http.StatusOK, "login", htmlvariable)
}

// LoginRouterPost  POST "/login" を処理
func LoginRouterPost(c echo.Context) error {

	db := Repository()
	defer db.Close()

	userID := c.FormValue("userid")
	pass := c.FormValue("password")

	errStr := "指定されたユーザIDが存在しません"
	// SQL: select * from hospitals where userid = req.userid limit 1
	hospital, err := models.Hospitals(models.HospitalWhere.Userid.EQ(userID)).One(context.Background(), db)
	if err != nil || hospital == nil {
	} else {
		userpass := hospital.Userpass
		if err = bcrypt.CompareHashAndPassword([]byte(userpass), []byte(pass)); err != nil {
			errStr = "パスワードが間違っています"
		} else {
			// login success; create session and redirect to "/"
			session, _ := session.Get("oursession", c)
			session.Values["userid"] = hospital.HospitalID
			err = session.Save(c.Request(), c.Response())

			if hospital.HospitalID == 1 {
				return c.Redirect(http.StatusFound, "/admin")
			}

			return c.Redirect(http.StatusFound, "/")
		}

	}

	htmlvariable := loginHTMLtemplate{
		Title:  "多施設臨床トライアルシステム ログイン",
		NoUser: errStr,
		CSS:    "/css/login.css",
	}
	return c.Render(http.StatusOK, "login", htmlvariable)
}
