package widgets

import (
	"UltimaVSaveGameEditor/pkg/ultima_v_save"
	"github.com/rivo/tview"
)

type PartyCharacterDetails struct {
	Form *tview.Form

	SaveGame *ultima_v_save.SaveGame

	nameTextView   *tview.TextView
	nameInputField *tview.InputField
}

func (p *PartyCharacterDetails) Init(saveGame *ultima_v_save.SaveGame) {
	p.SaveGame = saveGame
	p.Form = tview.NewForm()
	p.Form.SetTitleAlign(tview.AlignTop)

	p.Form.SetBorder(true)

	p.nameTextView = tview.NewTextView().
		SetLabelWidth(ultima_v_save.NMaxPlayerNameSize).
		SetLabel("Name")

	p.nameInputField = createInputField("Name", "", ultima_v_save.NMaxPlayerNameSize)
	p.nameInputField.SetAcceptanceFunc(createAcceptanceFunc(true, ultima_v_save.NMaxPlayerNameSize))

	p.Form.AddFormItem(p.nameInputField)
	p.Form.AddFormItem(p.nameInputField)
}

func (p *PartyCharacterDetails) SetPlayer(nPlayer int) {
	if nPlayer < 0 || nPlayer >= ultima_v_save.NPlayers {
		// just in case
		return
	}
	player := p.SaveGame.Characters[nPlayer]

	name := player.GetNameAsString()
	p.nameTextView.SetText(name)
	p.nameInputField.SetText(name)
}
