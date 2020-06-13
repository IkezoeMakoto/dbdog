package services

import (
	"database/sql"
	"fmt"
	"github.com/IkezoeMakoto/dbdog/app/src/scheme"
)

func GetDescribe(db *sql.DB) []*scheme.Table {
	//  -- テーブル一覧取得 --
	rows, err := db.Query("show tables")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}
	var tables []*scheme.Table
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			fmt.Println(err)
		}
		t := scheme.NewTable(name)
		t.GetColumns(db)
		tables = append(tables, t)
	}
	return tables
}
