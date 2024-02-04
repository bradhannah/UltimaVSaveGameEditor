package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const defaultHeaderColor = tcell.ColorYellow
const defaultHeaderAttr = tcell.AttrBold

type PartySummaryWidget struct {
	Table *tview.Table

	header []*tview.TableCell
}

func createHeaderCell(cellStr string) *tview.TableCell {
	cell := tview.NewTableCell(cellStr)
	cell.SetMaxWidth(10)
	cell.SetAttributes(defaultHeaderAttr)
	cell.SetTextColor(defaultHeaderColor)
	return cell
}

func (p *PartySummaryWidget) Init() {
	p.Table = tview.NewTable()

	p.header = make([]*tview.TableCell, 0)

	p.header = append(p.header, createHeaderCell("Name"))
	p.header = append(p.header, createHeaderCell("Class"))
	p.header = append(p.header, createHeaderCell("Level"))

	for i, cell := range p.header {
		p.Table.SetCell(0, i, cell)
	}
}
