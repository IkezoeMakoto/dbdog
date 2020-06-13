package main

import (
	"database/sql"
	"fmt"
	"github.com/IkezoeMakoto/dbdog/app/src/database"
	"github.com/IkezoeMakoto/dbdog/app/src/scheme"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("hello World!")

	// todo: dotenvから読み込む
	d := database.NewDsn("127.0.0.1", "3306", "dbdog", "test", "testdb")
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
	var tables []*scheme.Table
	for rows.Next() {
		name := ""
		err := rows.Scan(&name)
		if err != nil {
			fmt.Println(err)
		}
		tables = append(tables, scheme.NewTable(name))
		fmt.Println("--- " + name + " ---")
		// テーブル詳細取得
		var desc scheme.Desc
		rows, err := db.Query("DESCRIBE " + name)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&desc.Field, &desc.Type, &desc.Null, &desc.Key, &desc.Default, &desc.Extra)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(desc.ToString())
		}
	}
}
