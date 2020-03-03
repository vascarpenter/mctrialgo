package routes

import (
	"context"
	"mctrialgo/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/boil"
	"golang.org/x/crypto/bcrypt"
)

type resetpassHTMLtemplate struct {
	Title        string
	HospitalName string
	ErrPassRest  string
	CSS          string
}

// ResetPassRouter  GET "/resetpass" を処理
func ResetPassRouter(c echo.Context) error {

	hospname := c.Get("UserName").(string)
	htmlvariable := resetpassHTMLtemplate{
		Title:        "パスワードの変更",
		HospitalName: hospname,
		ErrPassRest:  "",
		CSS:          "/css/resetpass.css",
	}
	return c.Render(http.StatusOK, "resetpass", htmlvariable)
}

// ResetPassRouterPost  POST "/resetpass" を処理
func ResetPassRouterPost(c echo.Context) error {

	hospid := c.Get("UserID").(uint)
	hospname := c.Get("UserName").(string)
	htmlvariable := resetpassHTMLtemplate{
		Title:        "パスワードの変更",
		HospitalName: hospname,
		ErrPassRest:  "",
		CSS:          "/css/resetpass.css",
	}

	db := Repository()
	defer db.Close()
	ctx := context.Background()

	oldpass := c.FormValue("oldpass")
	newpass := c.FormValue("newpass")

	hospital, err := models.Hospitals(models.HospitalWhere.HospitalID.EQ(hospid)).One(ctx, db)
	if err != nil || hospital == nil {
		htmlvariable.ErrPassRest = "ユーザが存在しません"
		return c.Render(http.StatusOK, "resetpass", htmlvariable)
	}
	userpass := hospital.Userpass
	if err = bcrypt.CompareHashAndPassword([]byte(userpass), []byte(oldpass)); err != nil {
		htmlvariable.ErrPassRest = "パスワードが間違っています"
		return c.Render(http.StatusOK, "resetpass", htmlvariable)
	}

	newpasscrypt, _ := bcrypt.GenerateFromPassword([]byte(newpass), 10)
	hospital.Userpass = string(newpasscrypt)
	hospital.Update(ctx, db, boil.Infer())

	htmlvariable.ErrPassRest = "パスワードが変更されました"
	return c.Render(http.StatusOK, "resetpass", htmlvariable)
}
