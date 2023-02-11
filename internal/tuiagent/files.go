package tuiagent

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/dayterr/gophkeeper_diploma/internal/storage"
	"github.com/dayterr/gophkeeper_diploma/internal/validators"
)

func (t TUIClient) addFileForm(msg string) *tview.Form {
	var b storage.Binary
	var filepath string

	form.AddInputField("Filename *", "", 20, nil, func(fn string) {
		b.Filename = fn
	}).SetFieldTextColor(tcell.ColorGreen).SetLabelColor(tcell.ColorGreen)

	form.AddInputField("Filepath *", "", 20, nil, func(fp string) {
		filepath = fp
	}).SetFieldTextColor(tcell.ColorGreen).SetLabelColor(tcell.ColorGreen)

	form.AddInputField("Metadata", "", 20, nil, func(md string) {
		b.Metadata = md
	}).SetFieldTextColor(tcell.ColorGreen).SetLabelColor(tcell.ColorGreen)

	if msg != "" {
		form.AddTextView("",
			msg,
			40, 5, false, false)
	}

	form.AddButton("Save file", func() {
		data, err := os.ReadFile(filepath)
		if err != nil {
			form.Clear(true)
			t.addFileForm("please, make sure that file exists")
		}
		b.Data = data

		err = validators.ValidateBinary(b)
		switch err {
		case validators.ErrorFileNameFieldEmpty:
			form.Clear(true)
			t.addFileForm(validators.ErrorFileNameFieldEmpty.Error())
		case validators.ErrorFileFieldEmpty:
			form.Clear(true)
			t.addFileForm(validators.ErrorFileFieldEmpty.Error())
		default:
			err = t.SendFile(b)
			if err != nil {
				form.Clear(true)
				t.addFileForm(err.Error())
			} else {
				form.Clear(true)
				form = t.formMainPageLogged("")
			}
		}

	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)

	return form
}

func (t TUIClient) listFilesForm() {
	filesList.Clear()
	texts, err := t.ListFiles()
	if err != nil {
		filesList.AddItem("some error occurred, please try again", "", 1, nil)
	}

	for index, binary := range texts {
		bi := fmt.Sprintf("id is %d", binary.ID)
		filesList.AddItem(binary.Filename+" "+bi+" "+binary.Metadata, "", rune(49+index), nil)
	}
	t.TUIApp.SetFocus(filesList)

}

func (t TUIClient) fileActionsForm() *tview.Form {
	form.AddButton("Add a file", func() {
		form.Clear(true)
		form = t.addFileForm("")
	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)

	form.AddButton("List/delete files", func() {
		form.Clear(true)
		t.fileDeleteForm()
		t.listFilesForm()
		pages.SwitchToPage("File list")
	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)

	return form
}

func (t TUIClient) fileDeleteForm() *tview.Form {
	var fileID string

	form.AddInputField("File id *", "", 20, nil, func(id string) {
		fileID = id
	}).SetFieldTextColor(tcell.ColorGreen).SetLabelColor(tcell.ColorGreen)

	form.AddButton("Delete a file", func() {
		t.DeleteFile(fileID)
		form = t.formMainPageLogged("user")
		pages.SwitchToPage("Authorized page")
	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)

	return form
}
