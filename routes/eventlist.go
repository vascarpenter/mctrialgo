package routes

import (
	"context"
	"mctrialgo/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type eventlistHTMLtemplate struct {
	Title        string
	HospitalName string
	Person       models.Patient
	Events       models.EventSlice
	CSS          string
}

// EventlistRouter  GET "/eventlist/:hosp/:ser" を処理
func EventlistRouter(c echo.Context) error {

	ctx := context.Background()
	db := Repository()
	defer db.Close()

	hosp := c.Param("hosp")
	ser := c.Param("ser")

	hospid := c.Get("UserID").(uint)
	hospidfromparam, err := strconv.Atoi(hosp)
	if err != nil || hospid == 0 || hospidfromparam != int(hospid) {
		// invalid access, go to logout
		return c.Redirect(http.StatusFound, "/logout")
	}

	// SQL: SELECT * FROM patients WHERE hospital_id=? AND serialid=?
	patient, err := models.Patients(qm.Where("hospital_id=? AND serialid=?", hosp, ser)).One(ctx, db)
	if err != nil {
		panic(err)
	}

	events, err := models.Events(qm.Where("hospital_id=? AND serialid=?", hosp, ser)).All(context.Background(), db)

	if err != nil {
		panic(err)
	}
	//fmt.Printf("%+v\n", events[0])
	username := c.Get("UserName").(string)

	htmlvariable := eventlistHTMLtemplate{
		Title:        "イベント一覧",
		HospitalName: username,
		Person:       *patient,
		Events:       events,
		CSS:          "/css/index.css",
	}
	return c.Render(http.StatusOK, "eventlist", htmlvariable)
}
