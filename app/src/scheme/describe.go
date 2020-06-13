package scheme

import (
	"database/sql"
	"fmt"
)

type Desc struct {
	Field   sql.NullString
	Type    sql.NullString
	Null    sql.NullString
	Key     sql.NullString
	Default sql.NullString
	Extra   sql.NullString
}

func (s Desc) ToString() string {
	return "Field:" + s.Field.String +
		" Type:" + s.Type.String +
		" Null:" + s.Null.String +
		" Key:" + s.Key.String +
		" Default:" + s.Default.String +
		" Extra:" + s.Extra.String
}

func (s Desc) Print() {
	fmt.Println(s.ToString())
}
