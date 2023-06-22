package app

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type App struct {
	*tview.Application
	Cmd chan string
}

func NewApp(lines []string) *App {
	cmd := make(chan string)
	view := NewView(lines, cmd)
	app := App{
		tview.NewApplication().SetRoot(view, true),
		cmd,
	}

	// ECSキーで抜ける
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyESC:
			app.Stop()
			return nil
		default:
			return event
		}
	})

	return &app
}
