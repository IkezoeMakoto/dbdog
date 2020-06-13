package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Dsn struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

// todo: ファイル切り出す
func (d *Dsn) ToString() string {
	return 	d.User + ":" + d.Pass + "@tcp(" + d.Host + ":" + d.Port + ")/" + d.Name + "?charset=utf8&parseTime=True&loc=Local"
}

func main() {
	fmt.Println("hello World!")

	// todo: dotenvから読み込む
	d := Dsn{
		"127.0.0.1",
		"3306",
		"dbdog",
		"test",
		"testdb",
	}
	db, err := sql.Open("mysql", d.ToString())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//  -- テーブル一覧取得 --
	rows, err := db.Query("show tables")
	if err != nil {
		panic(err.Error())
	}
	var names []string
	for rows.Next() {
		name := ""
		err := rows.Scan(&name)
		if err != nil {
			fmt.Println(err)
		}
		names = append(names, name)
		fmt.Println("--- " + name + " ---")
		// テーブル詳細取得
		var dbField,dbType,dbNull,dbKey,dbDefault,dbExtra sql.NullString
		rows, err := db.Query("DESCRIBE " + name)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&dbField, &dbType, &dbNull, &dbKey, &dbDefault, &dbExtra)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(dbField, dbType, dbNull, dbKey, dbDefault, dbExtra)
		}
	}
}
