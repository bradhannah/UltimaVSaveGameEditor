package widgets

import (
	"fmt"
	"github.com/rivo/tview"
)

func createHeaderCell(cellStr string) *tview.TableCell {
	cell := tview.NewTableCell(cellStr)
	cell.SetAttributes(defaultHeaderAttr)
	cell.SetTextColor(defaultHeaderColor)
	cell.SetSelectable(false)
	return cell
}

func createDataCellStr(cellStr string) *tview.TableCell {
	cell := tview.NewTableCell(cellStr)
	//cell.SetMaxWidth(2)
	return cell
}

func createDataCellByte(cellByte byte) *tview.TableCell {
	cell := createDataCellStr(fmt.Sprintf("%d", cellByte))
	cell.SetMaxWidth(3) // typically 0-99 or 0-255
	return cell
}

func createDataCellNum(integer int) *tview.TableCell {
	cell := createDataCellStr(fmt.Sprintf("%d", integer))
	cell.SetMaxWidth(4) // 0 - 9999
	return cell
}
