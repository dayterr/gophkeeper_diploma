package tuiagent

import "github.com/dayterr/gophkeeper_diploma/internal/storage"

func validateUser(user storage.User) error {
	if user.Login == "" {
		return ErrorEmptyLogin
	}

	if user.Password == "" {
		return ErrorEmptyPassword
	}

	return nil
}

func validateCard(card storage.Card) error {
	if len(card.CardNumber) != 16 || len(card.CardNumber) != 18 {
		return ErrorInvalidCardNumber
	}

	if len(card.CVV) != 3 {
		return ErrorInvalidCVV
	}

	return nil
}
