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
)

type patientHTMLtemplate struct {
	Title        string
	HospitalName string
	AllowDate    string
	StartDate    string
	CSS          string
}

// PatientRouter  GET "/patient" を処理
func PatientRouter(c echo.Context) error {

	curDate := time.Now().Format("2006-01-02")
	hospname := c.Get("UserName").(string)
	htmlvariable := patientHTMLtemplate{
		Title:        "新規患者の登録",
		HospitalName: hospname,
		AllowDate:    curDate,
		StartDate:    curDate,
		CSS:          "/css/register.css",
	}
	return c.Render(http.StatusOK, "patient", htmlvariable)

}

// PatientRouterPost  POST "/patient" を処理
func PatientRouterPost(c echo.Context) error {
	ctx := context.Background()
	db := Repository()
	defer db.Close()

	var patient models.Patient

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

	hospid := c.Get("UserID").(uint)
	count, _ := models.Patients(models.PatientWhere.HospitalID.EQ(hospid)).Count(ctx, db)
	count++
	patient.PatientID = null.StringFrom(c.FormValue("patientid"))
	patient.HospitalID = hospid
	patient.Serialid = uint(count)
	patient.Initial = null.StringFrom(initial)
	patient.Birthdate = null.TimeFrom(birth)
	patient.Age = null.IntFrom(years)
	patient.Trialgroup = trialgroup
	patient.Allowdate = null.TimeFrom(allowdate)
	patient.Startdate = null.TimeFrom(startdate)
	patient.Female = null.BoolFrom(c.FormValue("sex") != "0")
	//	fmt.Printf("%+v\n", patient)

	patient.Insert(ctx, db, boil.Infer())
	return c.Redirect(http.StatusFound, "/")

}
