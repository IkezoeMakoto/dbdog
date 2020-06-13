package scheme

import (
	"database/sql"
)

type Desc struct {
	Field sql.NullString
	Type sql.NullString
	Null sql.NullString
	Key sql.NullString
	Default sql.NullString
	Extra sql.NullString
}

// func NewDesc(f, t, n, k, d, e) *Desc {
// 	return &Desc{f, t, n, k, d, e}
// }

func (s Desc) ToString() string {
	return "Field:" + s.Field.String +
		" Type:" + s.Type.String +
		" Null:" + s.Null.String +
		" Key:" + s.Key.String +
		" Default:" + s.Default.String +
		" Extra:" + s.Extra.String
}
