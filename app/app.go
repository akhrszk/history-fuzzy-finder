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
	view := NewView(lines)
	app := App{
		tview.NewApplication().SetRoot(view, true),
		cmd,
	}

	view.input.SetChangedFunc(func(text string) {
		view.ChangeQuery(text)
	})
	view.input.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyDown:
			view.history.CursorDown()
			return nil
		case tcell.KeyUp:
			view.history.CursorUp()
			return nil
		case tcell.KeyEnter:
			cmd := view.history.Selected()
			if cmd == nil {
				return nil
			}
			app.Cmd <- *cmd
			return nil
		default:
			return event
		}
	})

	return &app
}
