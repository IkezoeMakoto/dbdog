package scheme

type Tables struct {
	List []*Table
}

func (ts *Tables) Push(t *Table) {
	ts.List = append(ts.List, t)
}

func (ts Tables) Print() {
	for _, t := range ts.List {
		t.Print()
	}
}
