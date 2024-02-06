package widgets

import (
	"UltimaVSaveGameEditor/pkg/ultima_v_save"
	"fmt"
	"github.com/rivo/tview"
	"unicode"
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

func createInputField(label string, field string, nMaxLength int) *tview.InputField {
	inputField := tview.NewInputField()

	inputField.SetLabelWidth(nMaxLength)
	inputField.SetLabel(label)
	inputField.SetText(field)
	inputField.SetFieldWidth(nMaxLength + 1)
	inputField.SetTitleAlign(tview.AlignTop)
	return inputField
}

func createAcceptanceFunc(bOnlyAlpha bool, nMaxSize int) func(string, rune) bool {
	return func(textToCheck string, lastChar rune) bool {
		if len(textToCheck) > nMaxSize {
			return false
		}
		if bOnlyAlpha && !unicode.IsLetter(lastChar) {
			return false
		}
		return true
	}
}

func createDropDown(label string, nWidth int) *tview.DropDown {
	dropDown := tview.NewDropDown()
	dropDown.SetFieldWidth(nWidth)
	dropDown.SetLabel(label)
	return dropDown
}

func setCurrentDropDownOptionsByClass(characterClass ultima_v_save.CharacterClass, dropDown *tview.DropDown) {
	nIndex := ultima_v_save.FindIndexFromSliceT[ultima_v_save.CharacterClass](ultima_v_save.CharacterClassOrderedOptions, characterClass)
	dropDown.SetCurrentOption(nIndex)
}
