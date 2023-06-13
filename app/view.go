package app

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type View struct {
	*tview.Flex
	input   *QueryInput
	history *HistoryView
}

func NewView(lines []string, cmd chan string) *View {
	input := NewQueryInput()
	history := NewHistoryView(lines)

	input.SetChangedFunc(func(query string) {
		history.Filter(query)
	})
	input.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyDown:
			history.CursorDown()
			return nil
		case tcell.KeyUp:
			history.CursorUp()
			return nil
		case tcell.KeyEnter:
			s := history.Selected()
			if s == nil {
				return nil
			}
			cmd <- *s
			return nil
		default:
			return event
		}
	})

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(input, 3, 1, true).
		AddItem(history, 0, 1, false)

	return &View{
		layout,
		input,
		history,
	}
}

func (v *View) ChangeQuery(query string) {
	v.history.Filter(query)
}
