package routes

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/volatiletech/sqlboiler/boil"
)

// Repository DBを開いてORMと合体させる
func Repository() *sql.DB {
	var host, mysqluser, mysqlpass, mysqldb string
	if host = os.Getenv("MYSQL_HOST"); host == "" {
		host = "127.0.0.1"
	}
	if mysqluser = os.Getenv("MYSQL_USER"); mysqluser == "" {
		mysqluser = "oge"
	}
	if mysqlpass = os.Getenv("MYSQL_PASSWORD"); mysqlpass == "" {
		mysqlpass = "hogehogeA00"
	}
	if mysqldb = os.Getenv("MYSQL_DATABASE"); mysqldb == "" {
		mysqldb = "studydb"
	}
	sqlstring := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true&loc=Asia%%2FTokyo", mysqluser, mysqlpass, host, mysqldb)
	db, err := sql.Open("mysql", sqlstring)

	if err != nil {
		panic(err)
	}

	boil.SetDB(db)
	boil.DebugMode = true
	return db
}

// DiffDate  from stack overflow; diffDate
func DiffDate(a, b time.Time) (year, month, day, hour, min, sec int) {
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
