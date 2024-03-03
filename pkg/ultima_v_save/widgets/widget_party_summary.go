package widgets

import (
	"UltimaVSaveGameEditor/pkg/ultima_v_save"
	"github.com/rivo/tview"
)

type PartySummaryWidget struct {
	Table *tview.Table

	header []*tview.TableCell

	SaveGame *ultima_v_save.SaveGame
}

func (p *PartySummaryWidget) Init(saveGame *ultima_v_save.SaveGame, selectionChangedFunc func(nPlayer int, nJunk int)) {
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

func (p *PartySummaryWidget) populateCharacters() {
	for i, character := range p.SaveGame.Characters {
		row := i + 1
		p.Table.SetCell(row, 0, createDataCellStr(character.GetNameAsString()))
		p.Table.SetCell(row, 1, createDataCellStr(ultima_v_save.CharacterStatuses.GetById(character.Status).FriendlyName))
		p.Table.SetCell(row, 2, createDataCellByte(character.Level))
	}
}
