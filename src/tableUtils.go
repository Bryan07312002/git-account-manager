package main

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
)

func getDefaultTable() table.Writer {
	t := table.NewWriter()

	t.SetStyle(table.StyleLight)
	t.Style().Color.Header = text.Colors{text.FgGreen}
	t.Style().Color.Row = text.Colors{text.FgWhite}
	t.Style().Color.Separator = text.Colors{text.FgYellow}
	t.Style().Format.Footer = text.FormatLower
	t.Style().Options.DrawBorder = false

	t.SetOutputMirror(os.Stdout)
	return t
}
