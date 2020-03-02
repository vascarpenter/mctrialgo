package routes

import (
	"context"
	"mctrialgo/models"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/volatiletech/null"
	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// LoginRouter  handles "/login"
func LoginRouter(c echo.Context) error {

	htmlvariable := struct {
		Title  string
		NoUser string
		Css    string
	}{
		Title:  "ログイン",
		NoUser: "",
		Css:    "/css/login.css",
	}
	return c.Render(http.StatusOK, "login", htmlvariable)
}

func LoginRouterPost(c echo.Context) error {

	db := Repository()
	defer db.Close()

	username := c.FormValue("userid")
	pass := c.FormValue("password")

	errStr := "指定されたユーザIDが存在しません"
	// SQL: select * from hospitals where userid = req.UserName limit 1
	hospitals, err := models.Hospitals(models.HospitalWhere.Userid.EQ(null.StringFrom(username)), qm.Limit(1)).All(context.Background(), db)
	if err != nil || hospitals == nil {
	} else {
		userpass, err := hospitals[0].Userpass.MarshalText()
		if err = bcrypt.CompareHashAndPassword(userpass, []byte(pass)); err != nil {
			errStr = "パスワードが間違っています"
		} else {
			// login success; create session and redirect to "/"
			session, _ := session.Get("oursession", c)
			session.Values["userid"] = hospitals[0].HospitalID
			err = session.Save(c.Request(), c.Response())
			return c.Redirect(http.StatusFound, "/")
		}

	}

	htmlvariable := struct {
		Title  string
		NoUser string
		Css    string
	}{
		Title:  "ログイン",
		NoUser: errStr,
		Css:    "/css/login.css",
	}
	return c.Render(http.StatusOK, "login", htmlvariable)
}
