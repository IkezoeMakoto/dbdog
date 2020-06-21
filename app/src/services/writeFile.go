package services

import (
	"github.com/IkezoeMakoto/dbdog/app/src/scheme"
	"os"
)

type WriteFile struct {
	ts scheme.Tables
	fn string
	md  []string
	html []string
}

func NewWriteFile(ts scheme.Tables, fn string) *WriteFile {
	return &WriteFile{
		ts: ts,
		fn: fn,
		md: []string{},
		html: []string{},
	}
}

func (wf *WriteFile) OutputMarkDown() error {
	// h1 dbdog
	wf.md = append(wf.md, "# DBdog")
	for _, t := range wf.ts.List {
		// h2 table
		wf.md = append(wf.md, "## " + t.Name)

		// table-header keys
		header := "|"
		partition := "|"
		for _, h := range t.DescKeys() {
			header = header + h + "|"
			partition = partition + "---|"
		}
		wf.md = append(wf.md, header)
		wf.md = append(wf.md, partition)

		// table-body details
		for _, c := range t.Columns {
			line := "|"
			for _, d := range c.ToDescDetails().List {
				line = line + d.Value + "|"
			}
			wf.md = append(wf.md, line)
		}
		wf.md = append(wf.md, "")
	}
	err := writer(wf.fn, wf.md)
	if err != nil {
		return err
	}
	return nil
}

func writer(filename string, texts []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, text := range texts {
		_, err := file.WriteString(text + "\n")
		// fmt.Fprint()の場合
		// _, err := fmt.Fprint(file, line)
		if err != nil {
			return err
		}
	}
	return nil
}
