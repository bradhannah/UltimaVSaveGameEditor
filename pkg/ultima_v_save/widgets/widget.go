package widgets

import (
	"UltimaVSaveGameEditor/pkg/ultima_v_save"
	"fmt"
	"github.com/rivo/tview"
	"strconv"
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

func createAcceptanceFunc(bOnlyAlpha bool, bOnlyNumber bool, nMaxSize int) func(string, rune) bool {
	return func(textToCheck string, lastChar rune) bool {
		nTextLength := len(textToCheck)
		bIsCharacterNumber := unicode.IsNumber(lastChar)
		if nTextLength > nMaxSize {
			return false
		}
		//if nTextLength == 0 && bIsCharacterNumber {
		//	// wish I could force it to be a value if it was empty
		//	return true
		//}
		if bOnlyAlpha && !unicode.IsLetter(lastChar) {
			return false
		}
		if bOnlyNumber && !bIsCharacterNumber {
			return false
		}
		return true
	}
}

func createNumericAcceptanceFunc(nMinSize uint16, nMaxSize uint16) func(string, rune) bool {
	return func(textToCheck string, lastChar rune) bool {
		nTextLength := len(textToCheck)
		if nTextLength == 0 {
			// wish I could set this to default
			return true
		}
		bIsCharacterNumber := unicode.IsNumber(lastChar)
		if !bIsCharacterNumber {
			return false
		}
		nValue, err := strconv.ParseUint(textToCheck, 10, 16)
		nValue16 := uint16(nValue)
		if err != nil {
			return false
		}
		if nTextLength == 1 && lastChar == '0' || nValue16 == 0 {
			return false
		}
		if nValue16 < nMinSize || nValue16 > nMaxSize {
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

func setDropDownOptionsByClass(characterClass ultima_v_save.CharacterClass, dropDown *tview.DropDown) {
	nIndex := ultima_v_save.CharacterClasses.GetIndex(characterClass)
	dropDown.SetCurrentOption(nIndex)
}

func setDropDownByStatus(status ultima_v_save.CharacterStatus, dropDown *tview.DropDown) {
	nIndex := ultima_v_save.CharacterStatuses.GetIndex(status)
	dropDown.SetCurrentOption(nIndex)
}

func clearAllOptionsInDropDown(d *tview.DropDown) {
	for i := d.GetOptionCount() - 1; i >= 0; i-- {
		d.RemoveOption(i)
	}
}
