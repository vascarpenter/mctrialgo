package routes

import (
	"context"
	"mctrialgo/models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type patienteditHTMLtemplate struct {
	Title        string
	HospitalName string
	Person       models.Patient
	CSS          string
}

// PatientEditRouter  GET "/patientedit/:hosp/:ser" を処理
func PatientEditRouter(c echo.Context) error {
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

	patient, err := models.Patients(qm.Where("hospital_id=? AND serialid=?", hosp, ser)).One(ctx, db)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%+v\n", patient)
	hospname := c.Get("UserName").(string)
	editing, _ := patient.Initial.MarshalText()
	editText := hospname + " " + string(editing) + "さんの基本情報の編集"

	htmlvariable := patienteditHTMLtemplate{
		Title:        editText,
		HospitalName: hospname,
		Person:       *patient,
		CSS:          "/css/register.css",
	}
	return c.Render(http.StatusOK, "patientedit", htmlvariable)

}

// PatientEditRouterPost  POST "/patientedit/:hosp/:ser" を処理
func PatientEditRouterPost(c echo.Context) error {
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

	patient, err := models.Patients(qm.Where("hospital_id=? AND serialid=?", hosp, ser)).One(ctx, db)
	if err != nil {
		panic(err)
	}
	initial := c.FormValue("initial")

	loc, _ := time.LoadLocation("Asia/Tokyo")
	birth, err := time.ParseInLocation("2006-01-02", string(c.FormValue("birth")), loc)
	if err != nil {
		panic(err)
	}
	allowdate, err := time.ParseInLocation("2006-01-02", string(c.FormValue("allowdate")), loc)
	if err != nil {
		panic(err)
	}
	startdate, err := time.ParseInLocation("2006-01-02", string(c.FormValue("startdate")), loc)
	if err != nil {
		panic(err)
	}
	trialgroup, err := strconv.Atoi(c.FormValue("trialgroup"))
	if err != nil {
		panic(err)
	}
	curTime := time.Now()
	years, _, _, _, _, _ := DiffDate(birth, curTime)

	patient.PatientID = null.StringFrom(c.FormValue("patientid"))
	patient.Initial = null.StringFrom(initial)
	patient.Birthdate = null.TimeFrom(birth)
	patient.Age = null.IntFrom(years)
	patient.Trialgroup = trialgroup
	patient.Allowdate = null.TimeFrom(allowdate)
	patient.Startdate = null.TimeFrom(startdate)
	patient.Female = null.BoolFrom(c.FormValue("sex") != "0")
	//fmt.Printf("%+v\n", patient)
	patient.Update(ctx, db, boil.Infer())
	return c.Redirect(http.StatusFound, "/")
}
