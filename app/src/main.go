package main

import (
	"database/sql"
	"github.com/IkezoeMakoto/dbdog/app/src/database"
	"github.com/IkezoeMakoto/dbdog/app/src/services"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	services.PrintLogo()

	// todo: dotenvから読み込む
	d := database.NewDsn("127.0.0.1", "3306", "dbdog", "test", "testdb")
	db, err := sql.Open("mysql", d.ToString())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	ts := services.GetDescribe(db)
	for _, t := range ts {
		t.Print()
	}
}
