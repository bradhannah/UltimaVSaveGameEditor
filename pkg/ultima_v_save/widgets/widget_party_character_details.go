package widgets

import (
	"UltimaVSaveGameEditor/pkg/ultima_v_save"
	"github.com/rivo/tview"
)

type PartyCharacterDetails struct {
	Form *tview.Form

	ultima_v_save.PlayerCharacter
}

func (p *PartyCharacterDetails) Init() {
	p.Form = tview.NewForm()

	p.Form.SetBorder(true)

	nNameLen := len(ultima_v_save.PlayerCharacter{}.Name)
	p.Form.AddInputField(
		"Name",
		"Ooof",
		nNameLen,
		nil,
		nil,
	)
}
