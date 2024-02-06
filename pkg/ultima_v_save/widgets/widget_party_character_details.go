package widgets

import (
	"UltimaVSaveGameEditor/pkg/ultima_v_save"
	"github.com/rivo/tview"
)

type PartyCharacterDetails struct {
	Form *tview.Form

	SaveGame *ultima_v_save.SaveGame

	nameInputField *tview.InputField
	classDropDown  *tview.DropDown
}

func (p *PartyCharacterDetails) Init(saveGame *ultima_v_save.SaveGame) {
	p.SaveGame = saveGame
	p.Form = tview.NewForm()
	p.Form.SetTitleAlign(tview.AlignTop)
	p.Form.SetBorder(true)

	p.nameInputField = createInputField("Name", "", ultima_v_save.NMaxPlayerNameSize)
	p.nameInputField.SetAcceptanceFunc(createAcceptanceFunc(true, ultima_v_save.NMaxPlayerNameSize))
	p.Form.AddFormItem(p.nameInputField)

	p.classDropDown = createDropDown("Class", ultima_v_save.NMaxPlayerNameSize)
	updateClassDropDown(false, p.classDropDown)
	p.Form.AddFormItem(p.classDropDown)

}

func updateClassDropDown(bIsAvatar bool, d *tview.DropDown) {
	clearAllOptions(d)

	d.AddOption(ultima_v_save.CharacterClassMap[ultima_v_save.Avatar], nil)
	if !bIsAvatar {
		d.AddOption(ultima_v_save.CharacterClassMap[ultima_v_save.Fighter], nil)
		d.AddOption(ultima_v_save.CharacterClassMap[ultima_v_save.Bard], nil)
		d.AddOption(ultima_v_save.CharacterClassMap[ultima_v_save.Wizard], nil)
	}
	d.SetCurrentOption(0)
}

func clearAllOptions(d *tview.DropDown) {
	for i := d.GetOptionCount() - 1; i >= 0; i-- {
		d.RemoveOption(i)
	}
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
	p.nameInputField.SetText(player.GetNameAsString())
	setCurrentDropDownOptionsByClass(player.Class, p.classDropDown)
}
