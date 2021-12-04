package table

import (
	"bytes"
)

// Table is a table that can be rendered to markdown. It supports a column based data API and pads columns to a constant width.
type Table struct {
	// NumRows is the number of rows the table has. Each column must have exactly this number of rows.
	NumRows int

	columns []column
}

type column struct {
	header string
	rows   []string

	width int
}

func (t *Table) AddColumn(header string, rows []string) {
	if len(rows) != t.NumRows {
		panic("wrong number of rows!")
	}

	width := len(header)
	for _, row := range rows {
		if rowWidth := len(row); rowWidth > width {
			width = rowWidth
		}
	}

	t.columns = append(t.columns, column{
		header: header,
		rows:   rows,
		width:  width,
	})
}

func (t *Table) ToBuffer(b *bytes.Buffer) {
	t.writeHeader(b)
	t.writeBody(b)
}

func (t *Table) writeHeader(b *bytes.Buffer) {
	for _, col := range t.columns {
		writeCell(b, ' ', col.header, col.width)
	}
	writeRowFinish(b)

	for _, col := range t.columns {
		writeCell(b, '-', "", col.width)
	}
	writeRowFinish(b)
}

func (t *Table) writeBody(b *bytes.Buffer) {
	for rowNum := 0; rowNum < t.NumRows; rowNum++ {
		for _, col := range t.columns {
			row := col.rows[rowNum]

			writeCell(b, ' ', row, col.width)
		}

		writeRowFinish(b)
	}
}

func writeCell(b *bytes.Buffer, fillChar byte, contents string, width int) {
	b.WriteString("|")
	b.WriteByte(fillChar)
	b.WriteString(contents)
	b.Write(bytes.Repeat([]byte{fillChar}, width-len(contents)))
	b.WriteByte(fillChar)
}

func writeRowFinish(b *bytes.Buffer) {
	b.WriteString("|\n")
}
