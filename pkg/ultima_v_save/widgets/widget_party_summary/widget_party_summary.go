package widget_party_summary

import (
	"UltimaVSaveGameEditor/pkg/ultima_v_save/game_state"
	"UltimaVSaveGameEditor/pkg/ultima_v_save/widgets"
	"UltimaVSaveGameEditor/pkg/ultima_v_save/widgets/widget_help_and_status_bar"
	"github.com/rivo/tview"
)

type PartySummaryWidget struct {
	Table *tview.Table

	header []*tview.TableCell

	SaveGame *game_state.GameState

	helpAndStatusBar *widget_help_and_status_bar.HelpAndStatusBarWidget
}

func (p *PartySummaryWidget) Init(saveGame *game_state.GameState, helpAndStatusBar *widget_help_and_status_bar.HelpAndStatusBarWidget, selectionChangedFunc func(nPlayer int, nJunk int)) {
	p.helpAndStatusBar = helpAndStatusBar
	p.SaveGame = saveGame

	p.Table = tview.NewTable()
	p.Table.SetSelectable(true, false)
	p.Table.SetBorder(widgets.BDebugBorders)

	p.header = make([]*tview.TableCell, 0)

	p.header = append(p.header, widgets.CreateHeaderCell("Name").SetMaxWidth(9))
	p.header = append(p.header, widgets.CreateHeaderCell("Class"))
	p.header = append(p.header, widgets.CreateHeaderCell("Level"))

	for i, cell := range p.header {
		p.Table.SetCell(0, i, cell)
	}

	p.Table.SetSelectionChangedFunc(selectionChangedFunc)

	p.populateCharacters()
}

func (p *PartySummaryWidget) SetHelp() {
	p.helpAndStatusBar.Clear()
	p.helpAndStatusBar.Prefix = "Party List"
	p.helpAndStatusBar.AppendUpDownNav()
	p.helpAndStatusBar.AppendQuit()

}

func (p *PartySummaryWidget) populateCharacters() {
	for i, character := range p.SaveGame.Characters {
		row := i + 1
		p.Table.SetCell(row, 0, widgets.CreateDataCellStr(character.GetNameAsString()))
		p.Table.SetCell(row, 1, widgets.CreateDataCellStr(game_state.CharacterStatuses.GetById(character.Status).FriendlyName))
		p.Table.SetCell(row, 2, widgets.CreateDataCellByte(character.Level))
	}
}

func (p *PartySummaryWidget) SubComponentHasFocus() bool {
	return p.GetFocus() != nil
}

func (p *PartySummaryWidget) GetFocus() *tview.Primitive {
	if p.Table.HasFocus() {
		var prim = (tview.Primitive)(p.Table)
		return &prim
	}
	return nil
}
