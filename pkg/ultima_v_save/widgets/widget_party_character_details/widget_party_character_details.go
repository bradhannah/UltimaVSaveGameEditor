package widget_party_character_details

import (
	//"UltimaVSaveGameEditor/pkg/ultima_v_save"
	. "UltimaVSaveGameEditor/pkg/ultima_v_save/game_state"
	"UltimaVSaveGameEditor/pkg/ultima_v_save/widgets"
	"UltimaVSaveGameEditor/pkg/ultima_v_save/widgets/widget_help_and_status_bar"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type PartyCharacterDetailsWidget struct {
	Form *tview.Form

	OriginalGameState GameState
	GameState         *GameState

	nameInputField  *tview.InputField
	genderDropDown  *tview.DropDown
	statusDropDown  *tview.DropDown
	classDropDown   *tview.DropDown
	levelInputField *tview.InputField
	expInputField   *tview.InputField

	helpAndStatusBar *widget_help_and_status_bar.HelpAndStatusBarWidget
}

var partyCharacterDetailsWidget *PartyCharacterDetailsWidget

func (p *PartyCharacterDetailsWidget) Init(saveGame *GameState, helpAndStatusBar *widget_help_and_status_bar.HelpAndStatusBarWidget) {
	partyCharacterDetailsWidget = p
	p.helpAndStatusBar = helpAndStatusBar

	p.GameState = saveGame
	// make a copy of the data
	p.OriginalGameState = *saveGame

	p.Form = tview.NewForm()
	p.Form.SetTitleAlign(tview.AlignTop)
	p.Form.SetBorder(true)

	// Name
	p.nameInputField = widgets.CreateInputField("Name", "", NMaxPlayerNameSize)
	p.nameInputField.SetAcceptanceFunc(widgets.CreateAcceptanceFunc(true, false, NMaxPlayerNameSize))
	p.Form.AddFormItem(p.nameInputField)

	// Status
	p.statusDropDown = widgets.CreateDropDown("Status", 8)

	for _, val := range CharacterStatuses {
		p.statusDropDown.AddOption(val.FriendlyName, nil)
	}
	p.Form.AddFormItem(p.statusDropDown)

	// Gender
	p.genderDropDown = widgets.CreateDropDown("Gender", 6)
	p.genderDropDown.AddOption(CharacterGenders.GetById(Male).FriendlyName, nil)
	p.genderDropDown.AddOption(CharacterGenders.GetById(Female).FriendlyName, nil)
	p.Form.AddFormItem(p.genderDropDown)

	// Class
	p.classDropDown = widgets.CreateDropDown("Class", NMaxPlayerNameSize)
	updateClassDropDown(false, p.classDropDown)
	p.Form.AddFormItem(p.classDropDown)

	// Level
	p.levelInputField = widgets.CreateInputField("Level", "1", 1)
	p.levelInputField.SetAcceptanceFunc(widgets.CreateNumericAcceptanceFunc(1, 9))
	p.levelInputField.SetPlaceholder("1")
	p.Form.AddFormItem(p.levelInputField)

	// XP
	p.expInputField = widgets.CreateInputField("Exp", "0", 4)
	p.expInputField.SetAcceptanceFunc(widgets.CreateNumericAcceptanceFunc(0, 9999))
	p.Form.AddFormItem(p.expInputField)

	p.Form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			partyCharacterDetailsWidget.SaveChanges()
			return nil
		}
		return event
	},
	)
}

func (p *PartyCharacterDetailsWidget) SetHelp() {
	p.helpAndStatusBar.Clear()
	p.helpAndStatusBar.Prefix = "[black:white]Character Select[-]"
	p.helpAndStatusBar.AppendUpDownNav()
	p.helpAndStatusBar.AppendQuit()
}

func updateClassDropDown(bIsAvatar bool, d *tview.DropDown) {
	widgets.ClearAllOptionsInDropDown(d)

	d.AddOption(CharacterClasses.GetById(Avatar).FriendlyName, nil)
	if !bIsAvatar {
		d.AddOption(CharacterClasses.GetById(Fighter).FriendlyName, nil)
		d.AddOption(CharacterClasses.GetById(Bard).FriendlyName, nil)
		d.AddOption(CharacterClasses.GetById(Wizard).FriendlyName, nil)
	}
	d.SetCurrentOption(0)
}

func (p *PartyCharacterDetailsWidget) SetPlayer(nPlayer int) {
	if nPlayer < 0 || nPlayer >= NPlayers {
		// just in case
		return
	}
	player := p.GameState.Characters[nPlayer]

	updateClassDropDown(player.Class == Avatar, p.classDropDown)
	p.setPlayerFormValues(&player)
}

var currentPlayer PlayerCharacter

func (p *PartyCharacterDetailsWidget) setPlayerFormValues(player *PlayerCharacter) {
	currentPlayer = *player
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
	widgets.SetDropDownOptionsByClass(player.Class, p.classDropDown)
	// Level
	p.levelInputField.SetText(fmt.Sprintf("%0d", player.Level))
	// Exp
	p.expInputField.SetText(fmt.Sprintf("%d", player.Exp))
	// Status
	widgets.SetDropDownByStatus(player.Status, p.statusDropDown)
}

func (p *PartyCharacterDetailsWidget) SubComponentHasFocus() bool {
	return p.GetFocus() != nil
}

func (p *PartyCharacterDetailsWidget) GetFocus() *tview.Primitive {
	if p.Form.HasFocus() {
		var prim = (tview.Primitive)(p.Form)
		return &prim
	}

	for i := 0; i < p.Form.GetFormItemCount(); i++ {
		item := p.Form.GetFormItem(i)
		if item.HasFocus() {
			prim := item.(tview.Primitive)
			return &prim
		}
	}

	return nil
}

func (p *PartyCharacterDetailsWidget) SaveChanges() error {

	for nPlayer, player := range p.GameState.Characters {
		err := p.GameState.SaveCharactersOnSave("/Users/bradhannah/Google Drive/My Drive/games/u5/Games/Ultima_5/Gold/SAVED.GAM", uint(nPlayer), player)
		if err != nil {
			return err
		}
	}
	return nil
}
