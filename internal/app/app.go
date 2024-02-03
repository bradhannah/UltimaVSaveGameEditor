package main

import "github.com/rivo/tview"

type UltimaVSaveGameEditorApp struct {
	app       *tview.Application
	mainPages *tview.Pages
}

var ultimaVSaveGameEditorApp = UltimaVSaveGameEditorApp{}

func _initApp() {
	ultimaVSaveGameEditorApp.mainPages = tview.NewPages()
	mainBox := tview.NewBox()
	mainBox.SetBorder(true)
	ultimaVSaveGameEditorApp.mainPages.AddPage("MainBox", mainBox, true, true)
	ultimaVSaveGameEditorApp.app = tview.NewApplication().SetRoot(ultimaVSaveGameEditorApp.mainPages, true)
}

func main() {
	_initApp()
	err := ultimaVSaveGameEditorApp.app.Run()
	if err != nil {
		panic("This is bad man...")
	}
}
