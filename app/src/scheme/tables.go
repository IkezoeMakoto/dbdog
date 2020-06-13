package scheme

type Table struct{
	Name string
}

func NewTable(n string) *Table {
	return &Table{n}
}