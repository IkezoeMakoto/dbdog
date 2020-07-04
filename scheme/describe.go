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

type DescDetail struct {
	Key   string
	Value string
}
type DescDetails struct {
	List []DescDetail
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

func (s Desc) ToDescDetails() DescDetails {
	return DescDetails{
		List: []DescDetail{
			{Key: "Field", Value: s.Field.String},
			{Key: "Type", Value: s.Type.String},
			{Key: "Null", Value: s.Null.String},
			{Key: "Key", Value: s.Key.String},
			{Key: "Default", Value: s.Default.String},
			{Key: "Extra", Value: s.Extra.String},
		},
	}
}
