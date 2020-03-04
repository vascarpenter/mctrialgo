package routes

import (
	"context"
	"mctrialgo/models"
	"net/http"
	"time"

	"github.com/volatiletech/null"

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
	ctx := context.Background()

	userid := c.Get("UserID") // useid is not nil, because if nil, middleware detects it
	if userid == 0 || userid == nil {
		return c.Redirect(http.StatusFound, "/login")
	}
	useridint := userid.(uint)
	username := c.Get("UserName").(string)

	// SQL: SELECT * FROM patients WHERE hospital_id = ?
	patients, err := models.Patients(models.PatientWhere.HospitalID.EQ(useridint)).All(ctx, db)
	if err != nil {
		panic(err)
	}
	count64, _ := models.Patients(models.PatientWhere.HospitalID.EQ(useridint)).Count(ctx, db)
	count := int(count64)
	// fmt.Printf("%+v\n", patients[0])

	// 差分を計算
	for i := 0; i < count; i++ {
		if patients[i].Startdate.Valid {
			diff := time.Now().Sub(patients[i].Startdate.Time)
			patients[i].Diffdays = null.IntFrom(int(diff.Hours() / 24))
		}
	}

	htmlvariable := indexHTMLtemplate{
		Title:        "患者一覧",
		HospitalName: username,
		Patients:     patients,
		CSS:          "/css/index.css",
	}
	return c.Render(http.StatusOK, "index", htmlvariable)
}
