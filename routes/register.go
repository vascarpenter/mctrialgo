package routes

import (
	"context"
	"mctrialgo/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

type registerHTMLtemplate struct {
	Title       string
	EmailExists string
	CSS         string
}

// RegisterRouter  GET "/register" を処理
func RegisterRouter(c echo.Context) error {

	htmlvariable := registerHTMLtemplate{
		Title:       "病院登録",
		EmailExists: "",
		CSS:         "/css/register.css",
	}
	return c.Render(http.StatusOK, "register", htmlvariable)
}

// RegisterRouterPost  POST "/register" を処理
func RegisterRouterPost(c echo.Context) error {

	db := Repository()
	defer db.Close()

	htmlvariable := registerHTMLtemplate{
		Title:       "病院登録",
		EmailExists: "",
		CSS:         "/css/register.css",
	}

	hospName := c.FormValue("hospname")
	userid := c.FormValue("userid")
	userpass := c.FormValue("userpass")
	mailaddress := c.FormValue("mailaddress")
	createdAt := time.Now()

	// 'SELECT * FROM hospitals WHERE userid = ? LIMIT 1'
	ctx := context.Background()
	hospitals, err := models.Hospitals(models.HospitalWhere.Userid.EQ(userid), qm.Limit(1)).All(ctx, db)

	if err == nil && len(hospitals) > 0 {
		htmlvariable.EmailExists = "同じユーザーIDがすでに存在します"
		return c.Render(http.StatusOK, "register", htmlvariable)
	}

	userpasscrypt, err := bcrypt.GenerateFromPassword([]byte(userpass), 10)

	var hospital models.Hospital
	count, err := models.Hospitals().Count(ctx, db)
	hospital.HospitalID = uint(count) + 1
	hospital.Name = hospName
	hospital.Userid = userid
	hospital.Userpass = string(userpasscrypt)
	hospital.Mailaddress = null.StringFrom(mailaddress)
	hospital.CreatedAt = createdAt
	//	fmt.Printf("%+v\n", hospital)
	hospital.Insert(ctx, db, boil.Infer())

	htmlvariable.EmailExists = "追加しました"
	return c.Render(http.StatusOK, "register", htmlvariable)

	//	return c.Redirect(http.StatusFound, "/login")
}
