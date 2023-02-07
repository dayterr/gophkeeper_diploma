package validators

import "github.com/dayterr/gophkeeper_diploma/internal/storage"



func ValidateUser(user storage.User) error {
	if len(user.Login) < 4 {
		return ErrorLoginTooShort
	}

	if len(user.Password) < 8 {
		return ErrorPasswordTooShort
	}

	return nil
}

func ValidateCard(card storage.Card) error {
	if len(card.CardNumber) != 16 && len(card.CardNumber) != 18 {

		return ErrorInvalidCardNumber
	}

	if len(card.CVV) != 3 {
		return ErrorInvalidCVV
	}

	if card.ExpDate == "" {
		return ErrorEmptyExpDate
	}

	return nil
}

func ValidatePassword(password storage.Password) error {
	if len(password.Login) < 4 {
		return ErrorLoginTooShort
	}

	if len(password.Password) < 8 {
		return ErrorPasswordTooShort
	}

	return nil
}

func ValidateText(text storage.Text) error {
	if text.Data == "" {
		return ErrorTextFieldEmpty
	}

	return nil
}

func ValidateBinary(binary storage.Binary) error {
	if binary.Filename == "" {
		return ErrorFileNameFieldEmpty
	}

	if len(binary.Data) == 0 {
		return ErrorFileFieldEmpty
	}

	return nil
}