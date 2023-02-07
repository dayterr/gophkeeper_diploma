package validators

import "errors"

var ErrorLoginTooShort = errors.New("login can't be shorter than 4 symbols")
var ErrorPasswordTooShort = errors.New("password can't be shorter than 8 symbols")
var ErrorInvalidCardNumber = errors.New("the card number should consist of 16 or 18 digits")
var ErrorInvalidCVV = errors.New("cvv should consist of 3 digits")
var ErrorEmptyExpDate = errors.New("exp date can't be empty")
var ErrorAlreadyRegistered = errors.New("user with this login already exists, please, log in")
var ErrorLoginNotFound = errors.New("no user with this login, please, register")
var ErrorTextFieldEmpty = errors.New("text field can't be empty")
var ErrorFileNameFieldEmpty = errors.New("filename field can't be empty")
var ErrorFileFieldEmpty = errors.New("file is empty")
var ErrorCodeBadRequestRecieved = errors.New("please, fill all the required fields")
var ErrorCodeInternalErrorRecieved = errors.New("server side error, please, try again")
var ErrorDeleteFailed = errors.New("please, make sure that the data exists and belongs to you")