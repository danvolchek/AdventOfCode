package table_test

import (
	"github.com/danvolchek/AdventOfCode/cmd/readme/table"
	"testing"
)

type testColumn struct {
	header string
	rows   []string
}

type testCase struct {
	name    string
	numRows int
	columns []testColumn

	want string
}

func TestTable(t *testing.T) {
	for _, testCase := range []testCase{
		{
			name:    "no padding",
			numRows: 2,
			columns: []testColumn{
				{
					header: "foo",
					rows: []string{
						"bar",
						"baz",
					},
				},
				{
					header: "faa",
					rows: []string{
						"bir",
						"biz",
					},
				},
			},

			want: "| foo | faa |\n|-----|-----|\n| bar | bir |\n| baz | biz |\n",
		},
		{
			name:    "header padding",
			numRows: 2,
			columns: []testColumn{
				{
					header: "foofoofoo",
					rows: []string{
						"bar",
						"baz",
					},
				},
				{
					header: "faa",
					rows: []string{
						"bir",
						"biz",
					},
				},
			},

			want: "| foofoofoo | faa |\n|-----------|-----|\n| bar       | bir |\n| baz       | biz |\n",
		},
		{
			name:    "row padding",
			numRows: 2,
			columns: []testColumn{
				{
					header: "foo",
					rows: []string{
						"bar",
						"baz",
					},
				},
				{
					header: "faa",
					rows: []string{
						"birbirbir",
						"biz",
					},
				},
			},

			want: "| foo | faa       |\n|-----|-----------|\n| bar | birbirbir |\n| baz | biz       |\n",
		},
		{
			name:    "empty cells",
			numRows: 2,
			columns: []testColumn{
				{
					header: "",
					rows: []string{
						"bar",
						"baz",
					},
				},
				{
					header: "faa",
					rows: []string{
						"",
						"biz",
					},
				},
			},

			want: "|     | faa |\n|-----|-----|\n| bar |     |\n| baz | biz |\n",
		},
	} {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			tbl := &table.Table{NumRows: testCase.numRows}

			for _, testColumn := range testCase.columns {
				tbl.AddColumn(testColumn.header, testColumn.rows)
			}

			got := tbl.ToMarkdown()

			if got != testCase.want {
				t.Errorf("got\n%s\nwant\n%s\n", got, testCase.want)
			}
		})
	}
}
