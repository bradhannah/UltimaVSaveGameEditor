package widgets

import (
	"UltimaVSaveGameEditor/pkg/ultima_v_save"
	"github.com/rivo/tview"
)

type PartyInventoryListWidget struct {
	InventoryTable *tview.Table

	header []*tview.TableCell

	SaveGame *ultima_v_save.SaveGame

	helpAndStatusBar *HelpAndStatusBarWidget
}

func (p *PartyInventoryListWidget) SetHelp() {
	p.helpAndStatusBar.Clear()
	p.helpAndStatusBar.Prefix = "Party Inventory"
	p.helpAndStatusBar.AppendUpDownNav()
	p.helpAndStatusBar.AppendQuit()
}

type InventoryItemDetailsWidget struct {
	DetailsForm *tview.Form

	SaveGame *ultima_v_save.SaveGame

	helpAndStatusBar *HelpAndStatusBarWidget
}

func (i *InventoryItemDetailsWidget) SetHelp() {
	i.helpAndStatusBar.Clear()
	i.helpAndStatusBar.Prefix = "Inventory Detail"
	i.helpAndStatusBar.AppendUpDownNav()
	i.helpAndStatusBar.AppendQuit()
}

type PartyInventoryWidget struct {
	outerGrid *tview.Grid

	PartyInventoryListWidget   *PartyInventoryListWidget
	InventoryItemDetailsWidget *InventoryItemDetailsWidget

	SaveGame *ultima_v_save.SaveGame

	helpAndStatusBar *HelpAndStatusBarWidget
}

func (p *PartyInventoryWidget) Init(saveGame *ultima_v_save.SaveGame, helpAndStatusBar *HelpAndStatusBarWidget) {
	p.SaveGame = saveGame
	p.helpAndStatusBar = helpAndStatusBar
}

func (p *PartyInventoryWidget) SetHelp() {
	p.helpAndStatusBar.Clear()
	p.helpAndStatusBar.Prefix = "Party Inventory"
	p.helpAndStatusBar.AppendUpDownNav()
	p.helpAndStatusBar.AppendQuit()
}
