package main

import (
	"database/sql"
	"github.com/IkezoeMakoto/dbdog/app/src/database"
	"github.com/IkezoeMakoto/dbdog/app/src/services"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	services.PrintLogo()

	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env")
	}

	d := database.NewDsn(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("mysql", d.ToString())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	ts := services.GetDescribe(db)

	// todo: env から出力モード見て出力方法変える
	// txt, html, csv, puml
	for _, t := range ts {
		t.Print()
	}
}
