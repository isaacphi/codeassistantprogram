package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetText("Welcome to the cap code assistant program!").
		SetTextAlign(tview.AlignCenter).
		SetDoneFunc(func(key tcell.Key) {
			app.Stop()
		})

	if err := app.SetRoot(textView, true).Run(); err != nil {
		panic(err)
	}
}
