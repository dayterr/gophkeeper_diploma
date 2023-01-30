package tuiagent

import (
	"fmt"
	"github.com/dayterr/gophkeeper_diploma/internal/storage"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const MsgTryAgain = "Please, try again"

func (t TUIClient) registerUserForm(msg string) *tview.Form {
	var u storage.User

	form.AddInputField("Login *", "", 20, nil, func(l string) {
		u.Login = l
	}).SetFieldTextColor(tcell.ColorGreen).SetLabelColor(tcell.ColorGreen)

	form.AddInputField("Password *", "", 20, nil, func(p string) {
		u.Password = p
	}).SetFieldTextColor(tcell.ColorGreen).SetLabelColor(tcell.ColorGreen)

	if msg != "" {
		form.AddTextView("",
			msg,
			40, 5, false, false)
	}

	form.AddButton("Save", func() {
		err := validateUser(u)
		switch err {
		case ErrorEmptyLogin:
			form.Clear(true)
			t.registerUserForm(ErrorEmptyLogin.Error())
		case ErrorEmptyPassword:
			form.Clear(true)
			t.registerUserForm(ErrorEmptyPassword.Error())
		default:
			err := t.RegisterUser(u)
			if err == ErrorAlreadyRegistered {
				form.Clear(true)
				t.loginUserForm(err.Error())
			} else {
				if err != nil {
					form.Clear(true)
					t.registerUserForm(MsgTryAgain)
				} else {
					form.Clear(true)
					form = t.formMainPageLogged(u.Login)
				}
			}
		}

	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)

	return form

}

func (t TUIClient) loginUserForm(msg string) *tview.Form {
	var u storage.User

	form.AddInputField("Login *", "", 20, nil, func(l string) {
		u.Login = l
	}).SetFieldTextColor(tcell.ColorGreen).SetLabelColor(tcell.ColorGreen)

	form.AddInputField("Password *", "", 20, nil, func(p string) {
		u.Password = p
	}).SetFieldTextColor(tcell.ColorGreen).SetLabelColor(tcell.ColorGreen)

	if msg != "" {
		form.AddTextView("",
			msg,
			45, 5, false, false)
	}

	form.AddButton("Save", func() {
		err := validateUser(u)
		switch err {
		case ErrorEmptyLogin:
			form.Clear(true)
			t.loginUserForm(ErrorEmptyLogin.Error())
		case ErrorEmptyPassword:
			form.Clear(true)
			t.loginUserForm(ErrorEmptyPassword.Error())
		default:
			err := t.LogUser(u)
			if err == ErrorLoginNotFound {

				form.Clear(true)
				t.registerUserForm(err.Error())
			} else {
				if err != nil {
					form.Clear(true)
					t.loginUserForm(MsgTryAgain)
				} else {
					form.Clear(true)
					form = t.formMainPageLogged(u.Login)
				}
			}
		}

	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)

	return form
}

func (t TUIClient) formMainPageLogged(login string) *tview.Form {
	msg := fmt.Sprintf("Hello, %s! Which data would you like to work with?", login)
	form.AddTextView("", msg,
		50, 5, false, false)

	form.AddButton("Cards", func() {

	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)

	form.AddButton("Passwords", func() {
		//TODO
	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)

	form.AddButton("Binary data", func() {
		//TODO
	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)

	form.AddButton("Text data", func() {
		//TODO
	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)


	return form
}
