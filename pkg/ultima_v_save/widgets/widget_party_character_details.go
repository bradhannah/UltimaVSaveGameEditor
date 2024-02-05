package widgets

import (
	"UltimaVSaveGameEditor/pkg/ultima_v_save"
	"github.com/rivo/tview"
)

type PartyCharacterDetails struct {
	Form *tview.Form

	SaveGame *ultima_v_save.SaveGame
	//ultima_v_save.PlayerCharacter

	nameTextView *tview.TextView
}

func (p *PartyCharacterDetails) Init(saveGame *ultima_v_save.SaveGame) {
	p.SaveGame = saveGame
	p.Form = tview.NewForm()

	p.Form.SetBorder(true)

	p.nameTextView = tview.NewTextView().
		SetText("oof").
		SetLabelWidth(7).
		SetLabel("Name")

	p.Form.AddFormItem(p.nameTextView)
}

func (p *PartyCharacterDetails) SetPlayer(nPlayer int) {
	player := p.SaveGame.Characters[nPlayer]

	p.nameTextView.SetText(player.GetNameAsString())
}
