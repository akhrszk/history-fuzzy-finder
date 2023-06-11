package app

import (
	"fmt"
	"strings"

	"github.com/rivo/tview"
)

type HistoryView struct {
	*tview.TextView
	src    []string
	lines  []string
	cursor int
}

func NewHistoryView(src []string) *HistoryView {
	textView := tview.NewTextView()
	textView.SetBorder(true)
	textView.SetDynamicColors(true)

	lines := make([]string, 0)
	copy(lines, src)
	return &HistoryView{
		textView,
		src,
		lines,
		0,
	}
}

func (h *HistoryView) Filter(q string) {
	lines := make([]string, len(h.src))
	copy(lines, h.src)
	for _, token := range strings.Split(q, " ") {
		lines = reduce(lines, []string{}, func(acc string, curr []string) []string {
			if !strings.Contains(acc, token) {
				return curr
			}
			return append(curr, acc)
		})
	}
	h.lines = lines
	h.cursor = 0
	h.Sync()
}

func (h *HistoryView) CursorDown() {
	if len(h.lines) == 0 {
		return
	}
	if h.cursor+1 < len(h.lines) {
		h.cursor += 1
	} else {
		h.cursor = 0
	}
	h.Sync()
}

func (h *HistoryView) CursorUp() {
	if len(h.lines) == 0 {
		return
	}
	if h.cursor == 0 {
		h.cursor = len(h.lines) - 1
	} else {
		h.cursor -= 1
	}
	h.Sync()
}

func (h *HistoryView) Sync() {
	w := h.BatchWriter()
	w.Clear()
	for i, line := range h.lines {
		if i == int(h.cursor) {
			fmt.Fprintln(w, "[yellow]"+line+"[white]")
		} else {
			fmt.Fprintln(w, line)
		}
	}
	w.Close()
}

func (h *HistoryView) Selected() *string {
	if len(h.lines) == 0 {
		return nil
	}
	return &h.lines[h.cursor]
}
