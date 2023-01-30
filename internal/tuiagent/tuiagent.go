package tuiagent

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

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

func (t TUIClient) Run() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("starting tui app")

	flex.SetDirection(tview.FlexRow).AddItem(text, 0, 1, true)
	flexAuth.SetDirection(tview.FlexRow).AddItem(text, 0, 1, true)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 114:
			form.Clear(true)
			t.registerUserForm("")
			pages.SwitchToPage("Register page")
			/*fileLogger.Info().Msg(u.Login)
			err := t.RegisterUser(u)
			if err != nil {
				pages.SwitchToPage("Main page")
			} else {
				textAuth := fmt.Sprintf(`Hello, %s! Which data wuld you like to work with?`, u.Login)
				_ = textAuth
			}*/

		case 108:
			form.Clear(true)
			t.loginUserForm("")
			pages.SwitchToPage("Login page")
		case 97:
			form.Clear(true)
			addCardForm("")
			pages.SwitchToPage("Card page")
		default:
			t.TUIApp.Stop()
		}
		return event
	})

	/*form.AddTextArea("", "Hello, user! Please, register or log in", 5,
		5, 5, func(text string) {}("Hello, user! Please, register or log in"))*/

	pages.AddPage("Main page", flex, true, true)
	pages.AddPage("Card page", form, true, false)
	pages.AddPage("Register page", form, true, false)
	pages.AddPage("Login page", form, true, false)
	pages.AddPage("Authorized page", form, true, false)

	err := t.TUIApp.SetRoot(pages, true).EnableMouse(true).Run()
	if err != nil {
		log.Fatal().Err(err).Msg("error starting tui")
	}
}

