package routes

import (
	"context"
	"fmt"
	"mctrialgo/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type indexHTMLtemplate struct {
	Title        string
	HospitalName string
	Patients     models.PatientSlice
	CSS          string
}

// IndexRouter  GET "/" を処理
func IndexRouter(c echo.Context) error {

	db := Repository()
	defer db.Close()

	userid := c.Get("UserID") // useid is not nil, because if nil, middleware detects it
	if userid == 0 || userid == nil {
		return c.Redirect(http.StatusFound, "/login")
	}
	useridint := userid.(uint)
	username := c.Get("UserName").(string)

	// SQL: SELECT * FROM patients WHERE hospital_id = ?
	patients, err := models.Patients(models.PatientWhere.HospitalID.EQ(useridint)).All(context.Background(), db)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", patients[0])

	htmlvariable := indexHTMLtemplate{
		Title:        "患者一覧",
		HospitalName: username,
		Patients:     patients,
		CSS:          "/css/index.css",
	}
	return c.Render(http.StatusOK, "index", htmlvariable)
}
