package tuiagent

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/dayterr/gophkeeper_diploma/internal/storage"
	"github.com/dayterr/gophkeeper_diploma/internal/validators"
)

func (t TUIClient) addCardForm(msg string) *tview.Form {
	var c storage.Card

	form.AddInputField("Card Number *", "", 20, nil, func(cn string) {
		c.CardNumber = cn
	}).SetFieldTextColor(tcell.ColorGreen).SetLabelColor(tcell.ColorGreen)

	form.AddInputField("Expiration Date *", "", 20, nil, func(ed string) {
		c.ExpDate = ed
	}).SetFieldTextColor(tcell.ColorGreen).SetLabelColor(tcell.ColorGreen)

	form.AddInputField("Cardholder *", "", 20, nil, func(ch string) {
		c.Cardholder = ch
	}).SetFieldTextColor(tcell.ColorGreen).SetLabelColor(tcell.ColorGreen)

	form.AddInputField("CVV *", "", 20, nil, func(cvv string) {
		c.CVV = cvv
	}).SetFieldTextColor(tcell.ColorGreen).SetLabelColor(tcell.ColorGreen)

	form.AddInputField("Metadata", "", 20, nil, func(md string) {
		c.Metadata = md
	}).SetFieldTextColor(tcell.ColorGreen).SetLabelColor(tcell.ColorGreen)

	if msg != "" {
		form.AddTextView("",
			msg,
			40, 5, false, false)
	}

	form.AddButton("Save card", func() {
		err := validators.ValidateCard(c)
		switch err {
		case validators.ErrorInvalidCardNumber:
			form.Clear(true)
			t.addCardForm(validators.ErrorInvalidCardNumber.Error())
		case validators.ErrorInvalidCVV:
			form.Clear(true)
			t.addCardForm(validators.ErrorInvalidCVV.Error())
		default:
			err = t.SendCard(c)
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

func (t TUIClient) listCardsForm() {
	cardsList.Clear()
	cards, err := t.ListCards()
	if err != nil {
		textsList.AddItem("some error occurred, please try again", "", 1, nil)
	}

	for index, card := range cards {
		ci := fmt.Sprintf("id is %d", card.ID)
		cardsList.AddItem(card.CardNumber+" "+card.ExpDate+" "+ci+" "+card.Metadata, "", rune(49+index), nil)
	}
	t.TUIApp.SetFocus(cardsList)

}

func (t TUIClient) cardActionsForm(msg string) *tview.Form {
	form.AddButton("Add a card", func() {
		form.Clear(true)
		form = t.addCardForm("")
	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)

	form.AddButton("List/delete cards", func() {
		form.Clear(true)
		t.cardDeleteForm()
		t.listCardsForm()
		pages.SwitchToPage("Card list")
	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)

	if msg != "" {
		form.AddTextView("",
			msg,
			40, 5, false, false)
	}

	return form
}

func (t TUIClient) cardDeleteForm() *tview.Form {
	var cardID string

	form.AddInputField("Card id *", "", 20, nil, func(id string) {
		cardID = id
	}).SetFieldTextColor(tcell.ColorGreen).SetLabelColor(tcell.ColorGreen)

	form.AddButton("Delete a card", func() {
		t.DeleteCard(cardID)
		form = t.formMainPageLogged("user")
		pages.SwitchToPage("Authorized page")
	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)

	return form
}
