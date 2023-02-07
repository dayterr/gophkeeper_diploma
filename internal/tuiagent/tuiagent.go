package tuiagent

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const r = 114
const l = 108
const q = 113
const b = 98

var form = tview.NewForm()
var pages = tview.NewPages()
var flex = tview.NewFlex()
var text = tview.NewTextView().
	SetTextColor(tcell.ColorGreen).
	SetText(
		`Hello, user! Please, register or log in

(r) to register
(l) to login
(q) to quit`)
var flexAuth = tview.NewFlex()
var cardsList = tview.NewList()
var cardFlex = tview.NewFlex()
var passwordsList = tview.NewList()
var passwordFlex = tview.NewFlex()
var textsList = tview.NewList()
var textFlex = tview.NewFlex()
var filesList = tview.NewList()
var fileFlex = tview.NewFlex()
var textBack = tview.NewTextView().
	SetTextColor(tcell.ColorGreen).
	SetText(
		`(b) to get back to the menu`)

func (t TUIClient) Run() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("starting tui app")

	flex.SetDirection(tview.FlexRow).AddItem(text, 0, 1, true)
	flexAuth.SetDirection(tview.FlexRow).AddItem(text, 0, 1, true)
	cardFlex.SetDirection(tview.FlexRow).AddItem(cardsList, 0, 1, false)
	cardFlex.SetDirection(tview.FlexRow).AddItem(form, 0, 1, false)
	cardFlex.SetDirection(tview.FlexRow).AddItem(textBack, 0, 1, false)
	passwordFlex.SetDirection(tview.FlexRow).AddItem(passwordsList, 0, 1, false)
	passwordFlex.SetDirection(tview.FlexRow).AddItem(form, 0, 1, false)
	passwordFlex.SetDirection(tview.FlexRow).AddItem(textBack, 0, 1, false)
	textFlex.SetDirection(tview.FlexRow).AddItem(textsList, 0, 1, false)
	textFlex.SetDirection(tview.FlexRow).AddItem(form, 0, 1, false)
	textFlex.SetDirection(tview.FlexRow).AddItem(textBack, 0, 1, false)
	fileFlex.SetDirection(tview.FlexRow).AddItem(filesList, 0, 1, false)
	fileFlex.SetDirection(tview.FlexRow).AddItem(form, 0, 1, false)
	fileFlex.SetDirection(tview.FlexRow).AddItem(textBack, 0, 1, false)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case r:
			form.Clear(true)
			t.registerUserForm("")
			pages.SwitchToPage("Register page")
		case l:

			form.Clear(true)
			t.loginUserForm("")
			pages.SwitchToPage("Login page")
		default:
			t.TUIApp.Stop()
		}
		return event
	})

	cardFlex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case b:
			t.formMainPageLogged("user")
			pages.SwitchToPage("Authorized page")
		}
		return event
	})

	passwordFlex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case b:
			t.formMainPageLogged("user")
			pages.SwitchToPage("Authorized page")
		}
		return event
	})

	textFlex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case b:
			t.formMainPageLogged("user")
			pages.SwitchToPage("Authorized page")
		}
		return event
	})

	fileFlex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case b:
			t.formMainPageLogged("user")
			pages.SwitchToPage("Authorized page")
		}
		return event
	})

	pages.AddPage("Main page", flex, true, true)
	pages.AddPage("Card page", form, true, false)
	pages.AddPage("Register page", form, true, false)
	pages.AddPage("Login page", form, true, false)
	pages.AddPage("Authorized page", form, true, false)
	pages.AddPage("Card list", cardFlex, true, false)
	pages.AddPage("Password list", passwordFlex, true, false)
	pages.AddPage("Text list", textFlex, true, false)
	pages.AddPage("File list", fileFlex, true, false)

	err := t.TUIApp.SetRoot(pages, true).EnableMouse(true).Run()
	if err != nil {
		log.Fatal().Err(err).Msg("error starting tui")
	}
}

