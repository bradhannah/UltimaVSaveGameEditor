package main

import (
	"UltimaVSaveGameEditor/pkg/ultima_v_save/widgets"
	"github.com/rivo/tview"
)

type UltimaVSaveGameEditorApp struct {
	app           *tview.Application
	leftSidePages *tview.Pages // dynamic pages
	mainFlex      *tview.Flex  // main screen

	rightSideFlex *tview.Flex

	partySummaryWidget           *widgets.PartySummaryWidget
	playerCharacterDetailsWidget *widgets.PartyCharacterDetails
}

var ultimaVSaveGameEditorApp = UltimaVSaveGameEditorApp{}

func _initApp() {
	ultimaVSaveGameEditorApp.leftSidePages = tview.NewPages()
	ultimaVSaveGameEditorApp.leftSidePages.AddPage("MainFlex", tview.NewBox().SetBorder(true), true, true)
	ultimaVSaveGameEditorApp.leftSidePages.SetTitle("Edit")
	ultimaVSaveGameEditorApp.leftSidePages.SetBorder(true)

	ultimaVSaveGameEditorApp.rightSideFlex = tview.NewFlex().SetDirection(tview.FlexRow)
	ultimaVSaveGameEditorApp.rightSideFlex.SetBorder(true)
	ultimaVSaveGameEditorApp.rightSideFlex.SetTitle("Just Da Facts")

	ultimaVSaveGameEditorApp.partySummaryWidget = &widgets.PartySummaryWidget{}
	ultimaVSaveGameEditorApp.partySummaryWidget.Init()
	ultimaVSaveGameEditorApp.playerCharacterDetailsWidget = &widgets.PartyCharacterDetails{}
	ultimaVSaveGameEditorApp.playerCharacterDetailsWidget.Init()

	ultimaVSaveGameEditorApp.rightSideFlex.AddItem(ultimaVSaveGameEditorApp.partySummaryWidget.Table, 0, 1, false)
	ultimaVSaveGameEditorApp.rightSideFlex.AddItem(ultimaVSaveGameEditorApp.playerCharacterDetailsWidget.Form, 0, 1, false)

	ultimaVSaveGameEditorApp.mainFlex = tview.NewFlex()
	ultimaVSaveGameEditorApp.mainFlex.AddItem(ultimaVSaveGameEditorApp.leftSidePages, 0, 1, false)
	ultimaVSaveGameEditorApp.mainFlex.AddItem(ultimaVSaveGameEditorApp.rightSideFlex, 0, 1, false)

	ultimaVSaveGameEditorApp.app = tview.NewApplication().SetRoot(ultimaVSaveGameEditorApp.mainFlex, true)

	ultimaVSaveGameEditorApp.app.EnableMouse(true)
}

func main() {
	_initApp()

	err := ultimaVSaveGameEditorApp.app.Run()
	if err != nil {
		panic("This is bad man...")
	}
}
