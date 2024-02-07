package widgets

import (
	"UltimaVSaveGameEditor/pkg/ultima_v_save"
	"fmt"
	"github.com/rivo/tview"
)

type PartyCharacterDetails struct {
	Form *tview.Form

	SaveGame *ultima_v_save.SaveGame

	nameInputField  *tview.InputField
	genderDropDown  *tview.DropDown
	statusDropDown  *tview.DropDown
	classDropDown   *tview.DropDown
	levelInputField *tview.InputField
	expInputField   *tview.InputField
}

func (p *PartyCharacterDetails) Init(saveGame *ultima_v_save.SaveGame) {
	p.SaveGame = saveGame
	p.Form = tview.NewForm()
	p.Form.SetTitleAlign(tview.AlignTop)
	p.Form.SetBorder(true)

	// Name
	p.nameInputField = createInputField("Name", "", ultima_v_save.NMaxPlayerNameSize)
	p.nameInputField.SetAcceptanceFunc(createAcceptanceFunc(true, false, ultima_v_save.NMaxPlayerNameSize))
	p.Form.AddFormItem(p.nameInputField)

	// Status
	p.statusDropDown = createDropDown("Status", 8)
	p.statusDropDown.AddOption(ultima_v_save.CharacterStatusMap[ultima_v_save.Good], nil)
	p.statusDropDown.AddOption(ultima_v_save.CharacterStatusMap[ultima_v_save.Poisoned], nil)
	p.statusDropDown.AddOption(ultima_v_save.CharacterStatusMap[ultima_v_save.Sleep], nil)
	p.statusDropDown.AddOption(ultima_v_save.CharacterStatusMap[ultima_v_save.Charmed], nil)
	p.statusDropDown.AddOption(ultima_v_save.CharacterStatusMap[ultima_v_save.Dead], nil)
	p.Form.AddFormItem(p.statusDropDown)

	// Gender
	p.genderDropDown = createDropDown("Gender", 6)
	p.genderDropDown.AddOption(ultima_v_save.CharacterGenderMap[ultima_v_save.Male], nil)
	p.genderDropDown.AddOption(ultima_v_save.CharacterGenderMap[ultima_v_save.Female], nil)
	p.Form.AddFormItem(p.genderDropDown)

	// Class
	p.classDropDown = createDropDown("Class", ultima_v_save.NMaxPlayerNameSize)
	updateClassDropDown(false, p.classDropDown)
	p.Form.AddFormItem(p.classDropDown)

	// Level
	p.levelInputField = createInputField("Level", "1", 1)
	p.levelInputField.SetAcceptanceFunc(createNumericAcceptanceFunc(1, 9))
	p.levelInputField.SetPlaceholder("1")
	p.Form.AddFormItem(p.levelInputField)

	// XP
	p.expInputField = createInputField("Exp", "0", 4)
	p.expInputField.SetAcceptanceFunc(createNumericAcceptanceFunc(0, 9999))
	p.Form.AddFormItem(p.expInputField)

}

func updateClassDropDown(bIsAvatar bool, d *tview.DropDown) {
	clearAllOptionsInDropDown(d)

	d.AddOption(ultima_v_save.CharacterClassMap[ultima_v_save.Avatar], nil)
	if !bIsAvatar {
		d.AddOption(ultima_v_save.CharacterClassMap[ultima_v_save.Fighter], nil)
		d.AddOption(ultima_v_save.CharacterClassMap[ultima_v_save.Bard], nil)
		d.AddOption(ultima_v_save.CharacterClassMap[ultima_v_save.Wizard], nil)
	}
	d.SetCurrentOption(0)
}

func (p *PartyCharacterDetails) SetPlayer(nPlayer int) {
	if nPlayer < 0 || nPlayer >= ultima_v_save.NPlayers {
		// just in case
		return
	}
	player := p.SaveGame.Characters[nPlayer]

	updateClassDropDown(player.Class == ultima_v_save.Avatar, p.classDropDown)
	p.setPlayerFormValues(&player)
}

func (p *PartyCharacterDetails) setPlayerFormValues(player *ultima_v_save.PlayerCharacter) {
	// Name
	p.nameInputField.SetText(player.GetNameAsString())
	// Gender
	nGenderIndex := func() int {
		if player.Gender == ultima_v_save.Male {
			return 0
		}
		return 1
	}()
	p.genderDropDown.SetCurrentOption(nGenderIndex)
	// Class
	setCurrentDropDownOptionsByClass(player.Class, p.classDropDown)
	// Level
	p.levelInputField.SetText(fmt.Sprintf("%0d", player.Level))
	// Exp
	p.expInputField.SetText(fmt.Sprintf("%d", player.Exp))
}
