package routes

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/boil"
)

// Repository DBを開いてORMと合体させる
func Repository() *sql.DB {
	db, err := sql.Open("mysql", "oge:hogehogeA00@tcp(127.0.0.1:3306)/studydb?parseTime=true&loc=Asia%2FTokyo")

	if err != nil {
		panic(err)
	}

	boil.SetDB(db)
	boil.DebugMode = true
	return db
}
