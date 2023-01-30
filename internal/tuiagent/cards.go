package tuiagent

import (
	"github.com/dayterr/gophkeeper_diploma/internal/storage"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const MsgCardSaved = "The card data has been successfully saved"

func addCardForm(msg string) *tview.Form {
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
		err := validateCard(c)
		switch err {
		case ErrorInvalidCardNumber:
			form.Clear(true)
			addCardForm(ErrorInvalidCardNumber.Error())
		case ErrorInvalidCVV:
			form.Clear(true)
			addCardForm(ErrorInvalidCVV.Error())
		default:
			form.AddTextView("",
				MsgCardSaved,
				40, 5, false, false)
			pages.SwitchToPage("Main page")
		}

	}).SetButtonBackgroundColor(tcell.ColorDarkBlue).SetButtonTextColor(tcell.ColorGreen)


	return form
}

func listCardsForm() *tview.Form {

	return form
}
