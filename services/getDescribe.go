package services

import (
	"database/sql"
	"fmt"
	"github.com/IkezoeMakoto/dbdog/app/src/scheme"
)

func GetDescribe(db *sql.DB) scheme.Tables {
	//  -- テーブル一覧取得 --
	rows, err := db.Query("show tables")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}
	var tables scheme.Tables
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			fmt.Println(err)
		}
		t := scheme.NewTable(name)
		t.GetColumns(db)
		tables.Push(t)
	}
	return tables
}
