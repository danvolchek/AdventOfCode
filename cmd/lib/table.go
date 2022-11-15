package lib

import (
	"bytes"
)

// Table is a table that can be rendered to markdown. It supports a column based data API and pads columns to a constant width.
type Table struct {
	// NumRows is the number of rows the table has. Each column must have exactly this number of rows.
	NumRows int

	columns []column
}

// column is a column in the table.
type column struct {
	// header is the first row in the column.
	header string

	// rows is the rest of the rows in the column.
	rows []string

	// width is the largest length of the header/any row in the column.
	width int
}

// AddColumn adds a column to the table to the right of existing columns.
// len(rows) must be equal to NumRows or AddColumn panics.
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

// ToBuffer writes the table to a buffer with a trailing newline.
func (t *Table) ToBuffer(b *bytes.Buffer) {
	t.writeHeader(b)
	t.writeBody(b)
}

// writeHeader writes the first row of the table and header separator to the buffer.
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

// writeBody writes the rest of the table rows to the buffer.
func (t *Table) writeBody(b *bytes.Buffer) {
	for rowNum := 0; rowNum < t.NumRows; rowNum++ {
		for _, col := range t.columns {
			row := col.rows[rowNum]

			writeCell(b, ' ', row, col.width)
		}

		writeRowFinish(b)
	}
}

// writeCell writes a single cell to the buffer. fillChar is used to pad contents so that the total
// length written is equal to width; contents is centered between the padding.
func writeCell(b *bytes.Buffer, fillChar byte, contents string, width int) {
	b.WriteString("|")
	b.WriteByte(fillChar)
	b.WriteString(contents)
	b.Write(bytes.Repeat([]byte{fillChar}, width-len(contents)))
	b.WriteByte(fillChar)
}

// writeRowFinish writes the end of a row to the buffer.
func writeRowFinish(b *bytes.Buffer) {
	b.WriteString("|\n")
}
