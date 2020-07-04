package main

import (
	"database/sql"
	"github.com/IkezoeMakoto/dbdog/app/src/database"
	"github.com/IkezoeMakoto/dbdog/app/src/services"

	_ "github.com/go-sql-driver/mysql"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Db  Db
}

type Db struct {
	Host    string `toml:host`
	Port    string `toml:port`
	User    string `toml:user`
	Pass    string `toml:pass`
	Name    string `toml:name`
}

func main() {
	services.PrintLogo()

	var c Config
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		panic("error config.toml unload")
	}

	d := database.NewDsn(
		c.Db.Host,
		c.Db.Port,
		c.Db.User,
		c.Db.Pass,
		c.Db.Name,
	)
	db, err := sql.Open("mysql", d.ToString())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	ts := services.GetDescribe(db)

	ts.Print()
	wf := services.NewWriteFile(ts, "dbdog.txt")
	wf.OutputMarkDown()
}
