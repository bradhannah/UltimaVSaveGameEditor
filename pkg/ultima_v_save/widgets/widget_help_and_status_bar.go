package widgets

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strings"
)

const highlightKeyColor = tcell.ColorYellow
const functionDescColor = tcell.ColorLightCyan
const keySeparatorColor = tcell.ColorGhostWhite
const prefixColors = "black:white"
const resetColors = "white:black"

type HelpInfo interface {
	SetHelp()
}

type KeyMaps struct {
	Keys         []*tcell.EventKey
	FunctionDesc string
}

type HelpAndStatusBar struct {
	Bar     *tview.TextView
	KeyMaps []KeyMaps
	Prefix  string
}

func (h *HelpAndStatusBar) Init() {
	h.Bar = tview.NewTextView()
	h.Bar.SetSize(1, 0)
	h.Bar.SetDynamicColors(true)

	h.Clear()
}

func (h *HelpAndStatusBar) Clear() {
	h.KeyMaps = make([]KeyMaps, 0)
	h.Prefix = ""
	h.Bar.SetText("")
}

func (h *HelpAndStatusBar) AppendKeyMap(keyMap KeyMaps) *HelpAndStatusBar {
	h.KeyMaps = append(h.KeyMaps, keyMap)
	return h
}

func (h *HelpAndStatusBar) GetHelpAndStatusStr() string {
	statusStr := ""
	if h.Prefix != "" {
		statusStr = fmt.Sprintf("[%s]%s[%s]  ", prefixColors, h.Prefix, resetColors)
	}
	for i, val := range h.KeyMaps {
		if i > 0 {
			statusStr += "  "
		}
		statusStr += fmt.Sprintf("%s [%s]%s", combineKeys(val.Keys), functionDescColor, val.FunctionDesc)
	}
	return statusStr
}

func combineKeys(keys []*tcell.EventKey) string {
	cKeys := ""
	for i, key := range keys {
		if i > 0 {
			cKeys += fmt.Sprintf("[%s]/", keySeparatorColor)
		}
		cKeys += fmt.Sprintf("[%s]%s", highlightKeyColor, getAdjustedKey(key))
	}
	return cKeys
}

func getAdjustedKey(key *tcell.EventKey) string {
	switch key.Key() {
	case tcell.KeyUp:
		return string(rune('\u2191'))
	case tcell.KeyDown:
		return string(rune('\u2193'))
	default:
		if strings.HasPrefix(key.Name(), "Rune") {
			return string(key.Rune())
		}
		return key.Name()
	}
}

func (h *HelpAndStatusBar) AppendUpDownNav() {
	h.AppendKeyMap(KeyMaps{
		Keys: []*tcell.EventKey{
			tcell.NewEventKey(tcell.KeyUp, ' ', tcell.ModNone),
			tcell.NewEventKey(tcell.KeyDown, ' ', tcell.ModNone),
		},
		FunctionDesc: "Navigate",
	})
}

func (h *HelpAndStatusBar) AppendQuit() {
	h.AppendKeyMap(KeyMaps{
		Keys: []*tcell.EventKey{
			tcell.NewEventKey(tcell.KeyRune, 'q', tcell.ModNone),
		},
		FunctionDesc: "Quit",
	})
	h.Bar.SetText(h.GetHelpAndStatusStr())
}
