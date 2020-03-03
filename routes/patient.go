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

// from stack overflow; diffDate
func diffDate(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
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
	years, _, _, _, _, _ := diffDate(birth, curTime)

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
