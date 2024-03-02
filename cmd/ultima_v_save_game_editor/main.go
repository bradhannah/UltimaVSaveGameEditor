package main

import (
	"UltimaVSaveGameEditor/pkg/ultima_v_save"
	"UltimaVSaveGameEditor/pkg/ultima_v_save/widgets"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type UltimaVSaveGameEditorApp struct {
	app           *tview.Application
	leftSidePages *tview.Pages // dynamic pages
	mainFlex      *tview.Flex  // main screen

	rightSideGrid *tview.Grid

	partySummaryWidget           *widgets.PartySummaryWidget
	playerCharacterDetailsWidget *widgets.PartyCharacterDetails
}

var app = UltimaVSaveGameEditorApp{}

var SaveGame = ultima_v_save.SaveGame{}

func CreateInputHandlerTabToNext(next tview.Primitive) func(event *tcell.EventKey) *tcell.EventKey {
	return func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == '\t' {
			app.app.SetFocus(next)
			return nil
		}
		return event
	}
}

func globalInputHandler(eventKey *tcell.EventKey) *tcell.EventKey {
	keyRune := eventKey.Rune()
	if keyRune == 'q' || keyRune == 'Q' {
		app.app.Stop()
		return nil
	}

	return eventKey
}

func _initApp() {
	SaveGame, err := ultima_v_save.GetCharactersFromSave("/Users/bradhannah/Google Drive/My Drive/games/u5/Games/Ultima_5/Gold/SAVED.GAM")
	if err != nil {
		return
	}

	app.leftSidePages = tview.NewPages()
	app.leftSidePages.AddPage("MainFlex", tview.NewBox().SetBorder(true), true, true)
	app.leftSidePages.SetTitle("Edit")
	app.leftSidePages.SetBorder(true)
	//opof := tview.Primitive(*app.partySummaryWidget.Table))
	//app.app.SetFocus(app.mainFlex)

	app.partySummaryWidget = &widgets.PartySummaryWidget{}
	app.partySummaryWidget.Init(SaveGame, func(nPlayer int, nJunk int) {
		app.playerCharacterDetailsWidget.SetPlayer(nPlayer - 1)
	})
	app.playerCharacterDetailsWidget = &widgets.PartyCharacterDetails{}
	app.playerCharacterDetailsWidget.Init(SaveGame)

	app.rightSideGrid = tview.NewGrid()
	app.rightSideGrid.SetTitle("Just da facts...")
	app.rightSideGrid.SetRows(7+2, 0)
	app.rightSideGrid.SetColumns(0)
	app.rightSideGrid.AddItem(app.partySummaryWidget.Table, 0, 0, 1, 1, 0, 0, false)
	app.rightSideGrid.AddItem(app.playerCharacterDetailsWidget.Form, 1, 0, 1, 1, 0, 0, false)

	app.mainFlex = tview.NewFlex()
	app.mainFlex.AddItem(app.leftSidePages, 0, 1, false)
	app.mainFlex.AddItem(app.rightSideGrid, 0, 1, false)

	app.partySummaryWidget.Table.SetFixed(7, 3)

	app.app = tview.NewApplication().SetRoot(app.mainFlex, true)

	app.app.EnableMouse(true)

	app.app.SetInputCapture(globalInputHandler)

	app.leftSidePages.SetInputCapture(CreateInputHandlerTabToNext(app.partySummaryWidget.Table))
}

func main() {
	_initApp()

	err := app.app.Run()
	if err != nil {
		panic("This is bad man...")
	}
}
