package routes

import (
	"context"
	"mctrialgo/models"
	"net/http"
	"strconv"
	"time"

	"github.com/volatiletech/null"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type eventHTMLtemplate struct {
	Title        string
	HospitalName string
	Hospid       string
	Serid        string
	CSS          string
	TodayDate    string
}

// EventRouter  GET "/event/:hosp/:ser" を処理
func EventRouter(c echo.Context) error {
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

	editing, _ := patient.Initial.MarshalText()
	editText := string(editing) + "さんのイベントの追加"
	hospname := c.Get("UserName").(string)
	curDate := time.Now().Format("2006-01-02")

	htmlvariable := eventHTMLtemplate{
		Title:        editText,
		HospitalName: hospname,
		Hospid:       hosp,
		Serid:        ser,
		CSS:          "/css/event.css",
		TodayDate:    curDate,
	}
	return c.Render(http.StatusOK, "event", htmlvariable)
}

// EventRouterPost  POST "/event/:hosp/:ser" を処理
func EventRouterPost(c echo.Context) error {
	db := Repository()
	defer db.Close()
	ctx := context.Background()

	hosp := c.Param("hosp")
	ser := c.Param("ser")

	hospid := c.Get("UserID").(uint)
	hospidfromparam, err := strconv.Atoi(hosp)
	if err != nil || hospid == 0 || hospidfromparam != int(hospid) {
		// invalid access, go to logout
		return c.Redirect(http.StatusFound, "/logout")
	}

	loc, _ := time.LoadLocation("Asia/Tokyo")
	eventdate, err := time.ParseInLocation("2006-01-02", string(c.FormValue("date")), loc)
	if err != nil {
		panic(err)
	}
	eventtext := c.FormValue("eventtext")
	alive := c.FormValue("alive") != "0"
	bodyheight, _ := strconv.Atoi(c.FormValue("bh"))
	bodyweight, _ := strconv.Atoi(c.FormValue("bw"))
	sbp, _ := strconv.Atoi(c.FormValue("sbp"))
	dbp, _ := strconv.Atoi(c.FormValue("dbp"))
	hr, _ := strconv.Atoi(c.FormValue("hr"))

	var event models.Event
	count, err := models.Events(qm.Where("hospital_id=? AND serialid=?", hosp, ser)).Count(ctx, db)

	serInt, _ := strconv.Atoi(ser)
	event.HospitalID = hospid
	event.Serialid = uint(serInt)
	event.Eventid = uint(count + 1)
	event.Alive = null.BoolFrom(alive)
	event.Date = null.TimeFrom(eventdate)
	event.BH = null.IntFrom(bodyheight)
	event.BW = null.IntFrom(bodyweight)
	event.SBP = null.IntFrom(sbp)
	event.DBP = null.IntFrom(dbp)
	event.HR = null.IntFrom(hr)
	event.Event = null.StringFrom(eventtext)

	//	fmt.Printf("%+v\n", event)

	event.Insert(ctx, db, boil.Infer())

	return c.Redirect(http.StatusFound, "/eventlist/"+hosp+"/"+ser)
}
