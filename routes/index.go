package routes

import (
	"context"
	"mctrialgo/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null"
)

// IndexRouter  handles "/"
func IndexRouter(c echo.Context) error {

	db := Repository()
	defer db.Close()

	userid := c.Get("UserID") // useid is not nil, because if nil, middleware detects it
	useridint := int(userid.(uint))
	username := c.Get("UserName").(string)

	// SQL: SELECT * FROM patients WHERE hospital_id = ?
	patients, err := models.Patients(models.PatientWhere.HospitalID.EQ(null.IntFrom(useridint))).All(context.Background(), db)

	if err != nil {
		panic(err)
	}

	htmlvariable := struct {
		Title        string
		HospitalName string
		Patients     models.PatientSlice
		Css          string
	}{
		Title:        "テストページ",
		HospitalName: username,
		Patients:     patients,
		Css:          "/css/index.css",
	}
	return c.Render(http.StatusOK, "index", htmlvariable)
}