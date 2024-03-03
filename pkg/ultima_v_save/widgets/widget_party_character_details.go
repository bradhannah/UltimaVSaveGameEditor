package widgets

import (
	//"UltimaVSaveGameEditor/pkg/ultima_v_save"
	. "UltimaVSaveGameEditor/pkg/ultima_v_save"
	"fmt"
	"github.com/rivo/tview"
)

type PartyCharacterDetails struct {
	Form *tview.Form

	SaveGame *SaveGame

	nameInputField  *tview.InputField
	genderDropDown  *tview.DropDown
	statusDropDown  *tview.DropDown
	classDropDown   *tview.DropDown
	levelInputField *tview.InputField
	expInputField   *tview.InputField
}

func (p *PartyCharacterDetails) Init(saveGame *SaveGame) {
	p.SaveGame = saveGame
	p.Form = tview.NewForm()
	p.Form.SetTitleAlign(tview.AlignTop)
	p.Form.SetBorder(true)

	// Name
	p.nameInputField = createInputField("Name", "", NMaxPlayerNameSize)
	p.nameInputField.SetAcceptanceFunc(createAcceptanceFunc(true, false, NMaxPlayerNameSize))
	p.Form.AddFormItem(p.nameInputField)

	// Status
	p.statusDropDown = createDropDown("Status", 8)

	for _, val := range CharacterStatuses {
		p.statusDropDown.AddOption(val.FriendlyName, nil)
	}
	p.Form.AddFormItem(p.statusDropDown)

	// Gender
	p.genderDropDown = createDropDown("Gender", 6)
	p.genderDropDown.AddOption(CharacterGenders.GetById(Male).FriendlyName, nil)
	p.genderDropDown.AddOption(CharacterGenders.GetById(Female).FriendlyName, nil)
	p.Form.AddFormItem(p.genderDropDown)

	// Class
	p.classDropDown = createDropDown("Class", NMaxPlayerNameSize)
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

	d.AddOption(CharacterClasses.GetById(Avatar).FriendlyName, nil)
	if !bIsAvatar {
		d.AddOption(CharacterClasses.GetById(Fighter).FriendlyName, nil)
		d.AddOption(CharacterClasses.GetById(Bard).FriendlyName, nil)
		d.AddOption(CharacterClasses.GetById(Wizard).FriendlyName, nil)
	}
	d.SetCurrentOption(0)
}

func (p *PartyCharacterDetails) SetPlayer(nPlayer int) {
	if nPlayer < 0 || nPlayer >= NPlayers {
		// just in case
		return
	}
	player := p.SaveGame.Characters[nPlayer]

	updateClassDropDown(player.Class == Avatar, p.classDropDown)
	p.setPlayerFormValues(&player)
}

func (p *PartyCharacterDetails) setPlayerFormValues(player *PlayerCharacter) {
	// Name
	p.nameInputField.SetText(player.GetNameAsString())
	// Gender
	nGenderIndex := func() int {
		if player.Gender == Male {
			return 0
		}
		return 1
	}()
	p.genderDropDown.SetCurrentOption(nGenderIndex)
	// Class
	setDropDownOptionsByClass(player.Class, p.classDropDown)
	// Level
	p.levelInputField.SetText(fmt.Sprintf("%0d", player.Level))
	// Exp
	p.expInputField.SetText(fmt.Sprintf("%d", player.Exp))
	// Status
	setDropDownByStatus(player.Status, p.statusDropDown)
}
