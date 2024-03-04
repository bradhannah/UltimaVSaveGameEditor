package widgets

import (
	"UltimaVSaveGameEditor/pkg/ultima_v_save"
	"github.com/rivo/tview"
)

type PartySummaryWidget struct {
	Table *tview.Table

	header []*tview.TableCell

	SaveGame *ultima_v_save.SaveGame

	helpAndStatusBar *HelpAndStatusBar
}

func (p *PartySummaryWidget) Init(saveGame *ultima_v_save.SaveGame, helpAndStatusBar *HelpAndStatusBar, selectionChangedFunc func(nPlayer int, nJunk int)) {
	p.helpAndStatusBar = helpAndStatusBar
	p.SaveGame = saveGame

	p.Table = tview.NewTable()
	p.Table.SetSelectable(true, false)
	p.Table.SetBorder(bDebugBorders)

	p.header = make([]*tview.TableCell, 0)

	p.header = append(p.header, createHeaderCell("Name").SetMaxWidth(9))
	p.header = append(p.header, createHeaderCell("Class"))
	p.header = append(p.header, createHeaderCell("Level"))

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
		p.Table.SetCell(row, 0, createDataCellStr(character.GetNameAsString()))
		p.Table.SetCell(row, 1, createDataCellStr(ultima_v_save.CharacterStatuses.GetById(character.Status).FriendlyName))
		p.Table.SetCell(row, 2, createDataCellByte(character.Level))
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
