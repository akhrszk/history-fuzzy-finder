package app

import "github.com/rivo/tview"

type View struct {
	*tview.Flex
	input   *QueryInput
	history *HistoryView
}

func NewView(lines []string) *View {
	input := NewQueryInput()
	history := NewHistoryView(lines)

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
