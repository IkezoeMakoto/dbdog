package scheme

import (
	"database/sql"
	"fmt"
)

type Table struct {
	Name    string
	Columns []*Desc
}

func NewTable(n string) *Table {
	return &Table{Name: n, Columns: nil}
}

// テーブル詳細取得
func (t *Table) GetColumns(db *sql.DB) {
	rows, err := db.Query("DESCRIBE " + t.Name)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		d := Desc{}
		err := rows.Scan(&d.Field, &d.Type, &d.Null, &d.Key, &d.Default, &d.Extra)
		if err != nil {
			fmt.Println(err)
		}
		t.Columns = append(t.Columns, &d)
	}
}

func (t *Table) Print() {
	fmt.Println("# " + t.Name)
	for _, c := range t.Columns {
		c.Print()
	}
	fmt.Println("\n")
}

func (t *Table) DescKeys() []string {
	return []string{
		"Field",
		"Type",
		"Null",
		"Key",
		"Default",
		"Extra",
	}
}