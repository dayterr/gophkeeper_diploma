package tuiagent

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/dayterr/gophkeeper_diploma/internal/storage"
	"github.com/dayterr/gophkeeper_diploma/internal/validators"
)

func (t TUIClient) addTextForm(msg string) *tview.Form {
	var txt storage.Text

	form.AddInputField("Text *", "", 20, nil, func(text string) {
		txt.Data = text
	}).SetFieldTextColor(tcell.ColorGreen).SetLabelColor(tcell.ColorGreen)

	form.AddInputField("Metadata", "", 20, nil, func(md string) {
		txt.Metadata = md
	}).SetFieldTextColor(tcell.ColorGreen).SetLabelColor(tcell.ColorGreen)

	if msg != "" {
		form.AddTextView("",
			msg,
			40, 5, false, false)
	}

	form.AddButton("Save text", func() {
		err := validators.ValidateText(txt)
		switch err {
		case validators.ErrorTextFieldEmpty:
			form.Clear(true)
			t.addCardForm(validators.ErrorTextFieldEmpty.Error())
		default:
			err = t.SendText(txt)
			if err != nil {
				form.Clear(true)
				t.addTextForm(err.Error())
			} else {
				form.Clear(true)
				form = t.formMainPageLogged("")
			}
		}

	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)

	return form
}

func (t TUIClient) listTextsForm() {
	textsList.Clear()
	texts, err := t.ListTexts()
	if err != nil {
		textsList.AddItem("some error occurred, please try again", "", 1, nil)
	}

	for index, text := range texts {
		ti := fmt.Sprintf("id is %d", text.ID)
		textsList.AddItem(text.Data+" "+ti+" "+text.Metadata, "", rune(49+index), nil)
	}
	t.TUIApp.SetFocus(textsList)

}

func (t TUIClient) textActionsForm() *tview.Form {
	form.AddButton("Add a text", func() {
		form.Clear(true)
		form = t.addTextForm("")
	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)

	form.AddButton("List/delete texts", func() {
		form.Clear(true)
		t.textDeleteForm()
		t.listTextsForm()
		pages.SwitchToPage("Text list")
	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)

	return form
}

func (t TUIClient) textDeleteForm() *tview.Form {
	var textID string

	form.AddInputField("Text id *", "", 20, nil, func(id string) {
		textID = id
	}).SetFieldTextColor(tcell.ColorGreen).SetLabelColor(tcell.ColorGreen)

	form.AddButton("Delete a text", func() {
		t.DeleteText(textID)
		form = t.formMainPageLogged("user")
		pages.SwitchToPage("Authorized page")
	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)

	return form
}
