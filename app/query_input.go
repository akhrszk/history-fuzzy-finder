package app

import "github.com/rivo/tview"

type QueryInput struct {
	*tview.InputField
}

func NewQueryInput() *QueryInput {
	input := tview.NewInputField()
	input.SetLabel("filter:").SetBorder(true)
	return &QueryInput{
		input,
	}
}
