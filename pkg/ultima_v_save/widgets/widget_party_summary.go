package widgets

import (
	"UltimaVSaveGameEditor/pkg/ultima_v_save"
	"github.com/rivo/tview"
)

type PartySummaryWidget struct {
	Table *tview.Table

	header []*tview.TableCell
}

func (p *PartySummaryWidget) Init() {
	p.Table = tview.NewTable()
	p.Table.SetSelectable(true, false)
	p.Table.SetFixed(7, 0)
	p.Table.SetBorder(bDebugBorders)

	p.header = make([]*tview.TableCell, 0)

	p.header = append(p.header, createHeaderCell("Name").SetMaxWidth(9))
	p.header = append(p.header, createHeaderCell("Class"))
	p.header = append(p.header, createHeaderCell("Level"))

	for i, cell := range p.header {
		p.Table.SetCell(0, i, cell)
	}

	p.populateCharacters()
}

func (p *PartySummaryWidget) populateCharacters() {
	SaveGame, err := ultima_v_save.GetCharactersFromSave("/Users/bradhannah/Google Drive/My Drive/games/u5/Games/Ultima_5/Gold/SAVED.GAM")
	if err != nil {
		return
	}

	for i, character := range SaveGame.Characters {
		row := i + 1
		p.Table.SetCell(row, 0, createDataCellStr(character.GetNameAsString()))
		p.Table.SetCell(row, 1, createDataCellStr(ultima_v_save.CharacterClassMap[character.Class]))
		p.Table.SetCell(row, 2, createDataCellByte(character.Level))
	}
}
