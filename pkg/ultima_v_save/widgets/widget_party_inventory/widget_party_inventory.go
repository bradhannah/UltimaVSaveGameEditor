package widget_party_inventory

import (
	"UltimaVSaveGameEditor/pkg/ultima_v_save"
	"UltimaVSaveGameEditor/pkg/ultima_v_save/widgets/widget_help_and_status_bar"
	"github.com/rivo/tview"
)

type PartyInventoryListWidget struct {
	InventoryTable *tview.Table

	header []*tview.TableCell

	SaveGame *ultima_v_save.SaveGame

	helpAndStatusBar *widget_help_and_status_bar.HelpAndStatusBarWidget
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

	helpAndStatusBar *widget_help_and_status_bar.HelpAndStatusBarWidget
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

	helpAndStatusBar *widget_help_and_status_bar.HelpAndStatusBarWidget
}

func (p *PartyInventoryWidget) Init(saveGame *ultima_v_save.SaveGame, helpAndStatusBar *widget_help_and_status_bar.HelpAndStatusBarWidget) {
	p.SaveGame = saveGame
	p.helpAndStatusBar = helpAndStatusBar
}

func (p *PartyInventoryWidget) SetHelp() {
	p.helpAndStatusBar.Clear()
	p.helpAndStatusBar.Prefix = "Party Inventory"
	p.helpAndStatusBar.AppendUpDownNav()
	p.helpAndStatusBar.AppendQuit()
}
