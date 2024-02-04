package main

import (
	"UltimaVSaveGameEditor/pkg/ultima_v_save"
	"UltimaVSaveGameEditor/pkg/ultima_v_save/widgets"
	"github.com/rivo/tview"
)

type UltimaVSaveGameEditorApp struct {
	app           *tview.Application
	leftSidePages *tview.Pages // dynamic pages
	mainFlex      *tview.Flex  // main screen

	rightSideFlex *tview.Flex

	partySummaryWidget *widgets.PartySummaryWidget
}

var ultimaVSaveGameEditorApp = UltimaVSaveGameEditorApp{}

func _initApp() {
	ultimaVSaveGameEditorApp.leftSidePages = tview.NewPages()
	ultimaVSaveGameEditorApp.leftSidePages.AddPage("MainFlex", tview.NewBox().SetBorder(true), true, true)
	ultimaVSaveGameEditorApp.leftSidePages.SetTitle("Edit")
	ultimaVSaveGameEditorApp.leftSidePages.SetBorder(true)

	ultimaVSaveGameEditorApp.rightSideFlex = tview.NewFlex()
	ultimaVSaveGameEditorApp.rightSideFlex.SetBorder(true)
	ultimaVSaveGameEditorApp.rightSideFlex.SetTitle("Just Da Facts")

	ultimaVSaveGameEditorApp.partySummaryWidget = &widgets.PartySummaryWidget{}
	ultimaVSaveGameEditorApp.partySummaryWidget.Init()

	ultimaVSaveGameEditorApp.rightSideFlex.AddItem(ultimaVSaveGameEditorApp.partySummaryWidget.Table, 0, 1, false)

	ultimaVSaveGameEditorApp.mainFlex = tview.NewFlex()
	ultimaVSaveGameEditorApp.mainFlex.AddItem(ultimaVSaveGameEditorApp.leftSidePages, 0, 1, true)
	ultimaVSaveGameEditorApp.mainFlex.AddItem(ultimaVSaveGameEditorApp.rightSideFlex, 0, 1, false)

	ultimaVSaveGameEditorApp.app = tview.NewApplication().SetRoot(ultimaVSaveGameEditorApp.mainFlex, true)
}

func main() {
	_initApp()

	_, err := ultima_v_save.GetCharactersFromSave("/Users/bradhannah/Google Drive/My Drive/games/u5/Games/Ultima_5/Gold/SAVED.GAM")
	if err != nil {
		return
	}

	err = ultimaVSaveGameEditorApp.app.Run()
	if err != nil {
		panic("This is bad man...")
	}
}
