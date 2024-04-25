package main

import (
	"UltimaVSaveGameEditor/pkg/ultima_v_save/game_state"
	"UltimaVSaveGameEditor/pkg/ultima_v_save/widgets/widget_help_and_status_bar"
	"UltimaVSaveGameEditor/pkg/ultima_v_save/widgets/widget_party_character_details"
	"UltimaVSaveGameEditor/pkg/ultima_v_save/widgets/widget_party_summary"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type UltimaVSaveGameEditorApp struct {
	app           *tview.Application
	leftSidePages *tview.Pages // dynamic pages
	mainFlex      *tview.Flex  // main screen
	topGrid       *tview.Grid

	rightSideGrid *tview.Grid

	partySummaryWidget           *widget_party_summary.PartySummaryWidget
	playerCharacterDetailsWidget *widget_party_character_details.PartyCharacterDetailsWidget

	helpAndStatusBar *widget_help_and_status_bar.HelpAndStatusBarWidget

	OriginalSaveGameState *game_state.GameState
}

var app = UltimaVSaveGameEditorApp{}

//var SaveGame = ultima_v_save.SaveGame{}

func CreateInputHandlerTabToNext(next tview.Primitive) func(event *tcell.EventKey) *tcell.EventKey {
	return func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == '\t' {
			app.app.SetFocus(next)
			return nil
		}
		return event
	}
}

func updateHelpMessage() {
	if app.partySummaryWidget.SubComponentHasFocus() {
		app.partySummaryWidget.SetHelp()
	} else if app.playerCharacterDetailsWidget.SubComponentHasFocus() {
		app.playerCharacterDetailsWidget.SetHelp()
	} else {
		app.helpAndStatusBar.Clear()
	}
}

func globalInputHandler(eventKey *tcell.EventKey) *tcell.EventKey {
	keyRune := eventKey.Rune()

	if eventKey.Key() == tcell.KeyTab {
		app.app.SetFocus(app.playerCharacterDetailsWidget.Form)
	} else {
		switch keyRune {
		case 'q', 'Q':
			app.app.Stop()
			return nil
		}
	}

	return eventKey
}

func _initApp() {
	var originalGameState game_state.GameState
	app.OriginalSaveGameState = &originalGameState

	err := originalGameState.LoadSaveGame("/Users/bradhannah/Google Drive/My Drive/games/u5/Games/Ultima_5/Gold/SAVED.GAM")
	if err != nil {
		return
	}

	app.helpAndStatusBar = &widget_help_and_status_bar.HelpAndStatusBarWidget{}
	app.helpAndStatusBar.Init()

	app.leftSidePages = tview.NewPages()
	app.leftSidePages.AddPage("MainFlex", tview.NewBox().SetBorder(true), true, true)
	app.leftSidePages.SetTitle("Edit")
	app.leftSidePages.SetBorder(true)

	app.partySummaryWidget = &widget_party_summary.PartySummaryWidget{}
	app.partySummaryWidget.Init(&originalGameState,
		app.helpAndStatusBar,
		func(nPlayer int, _ int) {
			app.playerCharacterDetailsWidget.SetPlayer(nPlayer - 1)
		})
	app.playerCharacterDetailsWidget = &widget_party_character_details.PartyCharacterDetailsWidget{}
	app.playerCharacterDetailsWidget.Init(&originalGameState, app.helpAndStatusBar)

	app.rightSideGrid = tview.NewGrid()
	app.rightSideGrid.SetTitle("Just da facts...")
	app.rightSideGrid.SetRows(7+2, 0)
	app.rightSideGrid.SetColumns(0)
	app.rightSideGrid.AddItem(app.partySummaryWidget.Table, 0, 0, 1, 1, 0, 0, false)
	app.rightSideGrid.AddItem(app.playerCharacterDetailsWidget.Form, 1, 0, 1, 1, 0, 0, false)

	app.mainFlex = tview.NewFlex()
	app.mainFlex.AddItem(app.leftSidePages, 0, 1, false)
	app.mainFlex.AddItem(app.rightSideGrid, 0, 1, false)

	app.topGrid = tview.NewGrid()
	app.topGrid.SetRows(0, 1)
	app.topGrid.SetColumns(0)
	app.topGrid.AddItem(app.mainFlex, 0, 0, 1, 1, 1, 1, true)
	app.topGrid.AddItem(app.helpAndStatusBar.Bar, 1, 0, 1, 1, 1, 1, false)

	app.partySummaryWidget.Table.SetFixed(7, 3)

	app.app = tview.NewApplication().SetRoot(app.topGrid, true)

	app.app.EnableMouse(true)

	app.app.SetInputCapture(globalInputHandler)
	app.app.SetBeforeDrawFunc(func(screen tcell.Screen) bool {
		// just before drawing, let's make sure we are showing correct help
		updateHelpMessage()
		return false
	})
	app.app.SetFocus(app.partySummaryWidget.Table)
	app.leftSidePages.SetInputCapture(CreateInputHandlerTabToNext(app.partySummaryWidget.Table))

}

func main() {
	_initApp()

	err := app.app.Run()
	if err != nil {
		panic("This is bad man...")
	}
}
