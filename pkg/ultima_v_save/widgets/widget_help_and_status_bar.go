package widgets

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const highlightKeyColor = tcell.ColorYellow
const functionDescColor = tcell.ColorLightCyan

type KeyMaps struct {
	Keys         []*tcell.EventKey
	FunctionDesc string
}

type HelpAndStatusBar struct {
	Bar     *tview.TextView
	KeyMaps []KeyMaps
}

func (h *HelpAndStatusBar) Init() {
	h.Bar = tview.NewTextView()
	h.Bar.SetSize(1, 0)
	h.Bar.SetDynamicColors(true)

	h.Clear()
}

func (h *HelpAndStatusBar) Clear() {
	h.KeyMaps = make([]KeyMaps, 0)
}

func (h *HelpAndStatusBar) AppendKeyMap(keyMap KeyMaps) *HelpAndStatusBar {
	h.KeyMaps = append(h.KeyMaps, keyMap)
	return h
}

func (h *HelpAndStatusBar) GetHelpAndStatusStr() string {
	statusStr := ""
	for i, val := range h.KeyMaps {
		if i > 0 {
			statusStr += "  "
		}
		//statusStr += val.Keys[0].Name() + " " + val.FunctionDesc
		statusStr += fmt.Sprintf("[%s]%s [%s]%s", highlightKeyColor, getAdjustedKey(val.Keys[0]), functionDescColor, val.FunctionDesc)
	}
	return statusStr
}

func getAdjustedKey(key *tcell.EventKey) string {
	switch key.Key() {
	case tcell.KeyUp:
		return string(rune('\u2191'))
	default:
		return key.Name()
	}
}
