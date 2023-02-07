package tuiagent

import (
	"fmt"
	"github.com/dayterr/gophkeeper_diploma/internal/storage"
	"github.com/dayterr/gophkeeper_diploma/internal/validators"
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
		err := validators.ValidateUser(u)
		switch err {
		case validators.ErrorLoginTooShort:
			form.Clear(true)
			t.registerUserForm(validators.ErrorLoginTooShort.Error())
		case validators.ErrorPasswordTooShort:
			form.Clear(true)
			t.registerUserForm(validators.ErrorPasswordTooShort.Error())
		default:
			err := t.RegisterUser(u)
			if err == validators.ErrorAlreadyRegistered {
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
		err := validators.ValidateUser(u)
		switch err {
		case validators.ErrorLoginTooShort:
			form.Clear(true)
			t.loginUserForm(validators.ErrorLoginTooShort.Error())
		case validators.ErrorPasswordTooShort:
			form.Clear(true)
			t.loginUserForm(validators.ErrorPasswordTooShort.Error())
		default:
			err := t.LogUser(u)
			if err == validators.ErrorLoginNotFound {
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
	form.Clear(true)
	msg := fmt.Sprintf("Hello, %s! Which data would you like to work with?", login)
	form.AddTextView("", msg,
		50, 5, false, false)

	form.AddButton("Cards", func() {
		form.Clear(true)
		t.cardActionsForm("")
	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)

	form.AddButton("Passwords", func() {
		form.Clear(true)
		t.passwordActionsForm()
	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)

	form.AddButton("Files", func() {
		form.Clear(true)
		t.fileActionsForm()
	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)

	form.AddButton("Texts", func() {
		form.Clear(true)
		t.textActionsForm()
	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)


	return form
}

