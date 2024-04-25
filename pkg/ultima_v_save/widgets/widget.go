package widgets

import (
	"UltimaVSaveGameEditor/pkg/ultima_v_save/game_state"
	"fmt"
	"github.com/rivo/tview"
	"strconv"
	"unicode"
)

type SubComponentContainer interface {
	SubComponentHasFocus() bool
	GetFocus() *tview.Primitive
}

func CreateHeaderCell(cellStr string) *tview.TableCell {
	cell := tview.NewTableCell(cellStr)
	cell.SetAttributes(defaultHeaderAttr)
	cell.SetTextColor(defaultHeaderColor)
	cell.SetSelectable(false)
	return cell
}

func CreateDataCellStr(cellStr string) *tview.TableCell {
	cell := tview.NewTableCell(cellStr)
	return cell
}

func CreateDataCellByte(cellByte byte) *tview.TableCell {
	cell := CreateDataCellStr(fmt.Sprintf("%d", cellByte))
	cell.SetMaxWidth(3) // typically 0-99 or 0-255
	return cell
}

func createDataCellNum(integer int) *tview.TableCell {
	cell := CreateDataCellStr(fmt.Sprintf("%d", integer))
	cell.SetMaxWidth(4) // 0 - 9999
	return cell
}

func CreateInputField(label string, field string, nMaxLength int) *tview.InputField {
	inputField := tview.NewInputField()

	inputField.SetLabelWidth(nMaxLength)
	inputField.SetLabel(label)
	inputField.SetText(field)
	inputField.SetFieldWidth(nMaxLength + 1)
	inputField.SetTitleAlign(tview.AlignTop)
	return inputField
}

func CreateAcceptanceFunc(bOnlyAlpha bool, bOnlyNumber bool, nMaxSize int) func(string, rune) bool {
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

func CreateNumericAcceptanceFunc(nMinSize uint16, nMaxSize uint16) func(string, rune) bool {
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

func CreateDropDown(label string, nWidth int) *tview.DropDown {
	dropDown := tview.NewDropDown()
	dropDown.SetFieldWidth(nWidth)
	dropDown.SetLabel(label)
	return dropDown
}

func SetDropDownOptionsByClass(characterClass game_state.CharacterClass, dropDown *tview.DropDown) {
	nIndex := game_state.CharacterClasses.GetIndex(characterClass)
	dropDown.SetCurrentOption(nIndex)
}

func SetDropDownByStatus(status game_state.CharacterStatus, dropDown *tview.DropDown) {
	nIndex := game_state.CharacterStatuses.GetIndex(status)
	dropDown.SetCurrentOption(nIndex)
}

func ClearAllOptionsInDropDown(d *tview.DropDown) {
	for i := d.GetOptionCount() - 1; i >= 0; i-- {
		d.RemoveOption(i)
	}
}
