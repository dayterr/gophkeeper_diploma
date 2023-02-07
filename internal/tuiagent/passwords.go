package tuiagent

import (
	"github.com/dayterr/gophkeeper_diploma/internal/storage"
	"github.com/dayterr/gophkeeper_diploma/internal/validators"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"fmt"
)

func (t TUIClient) addPasswordForm(msg string) *tview.Form {
	var p storage.Password

	form.AddInputField("Login *", "", 20, nil, func(l string) {
		p.Login = l
	}).SetFieldTextColor(tcell.ColorGreen).SetLabelColor(tcell.ColorGreen)

	form.AddInputField("Password *", "", 20, nil, func(pw string) {
		p.Password = pw
	}).SetFieldTextColor(tcell.ColorGreen).SetLabelColor(tcell.ColorGreen)

	form.AddInputField("Metadata", "", 20, nil, func(md string) {
		p.Metadata = md
	}).SetFieldTextColor(tcell.ColorGreen).SetLabelColor(tcell.ColorGreen)

	if msg != "" {
		form.AddTextView("",
			msg,
			40, 5, false, false)
	}


	form.AddButton("Save card", func() {
		err := validators.ValidatePassword(p)
		switch err {
		case validators.ErrorLoginTooShort:
			form.Clear(true)
			t.addCardForm(validators.ErrorLoginTooShort.Error())
		case validators.ErrorPasswordTooShort:
			form.Clear(true)
			t.addCardForm(validators.ErrorPasswordTooShort.Error())
		default:
			err = t.SendPassword(p)
			if err != nil {
				form.Clear(true)
				t.addCardForm(err.Error())
			} else {
				form.Clear(true)
				form = t.formMainPageLogged("")
			}
		}

	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)

	return form
}

func (t TUIClient) listPasswordsForm()  {
	passwordsList.Clear()
	passwords, err := t.ListPasswords()
	if err != nil {
		textsList.AddItem("some error occurred, please try again", "", 1, nil)
	}

	for index, password := range passwords {
		pi := fmt.Sprintf("id is %d", password.ID)
		passwordsList.AddItem(password.Login + " " + password.Password + " " + pi + " " + password.Metadata, "", rune(49+index), nil)
	}
	t.TUIApp.SetFocus(passwordsList)

}

func (t TUIClient) passwordActionsForm() *tview.Form {
	form.AddButton("Add a password", func() {
		form.Clear(true)
		form = t.addPasswordForm("")
	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)

	form.AddButton("List/delete passwords", func() {
		form.Clear(true)
		t.passwordDeleteForm()
		t.listPasswordsForm()
		pages.SwitchToPage("Password list")
	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)

	return form
}

func (t TUIClient) passwordDeleteForm() *tview.Form {
	var passwordID string

	form.AddInputField("Password id *", "", 20, nil, func(id string) {
		passwordID = id
	}).SetFieldTextColor(tcell.ColorGreen).SetLabelColor(tcell.ColorGreen)

	form.AddButton("Delete a password", func() {
		t.DeletePassword(passwordID)
		form = t.formMainPageLogged("user")
		pages.SwitchToPage("Authorized page")
	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)

	return form
}
