package tuiagent

import "errors"

var ErrorEmptyLogin = errors.New("login can't be empty")
var ErrorEmptyPassword = errors.New("password can't be empty")
var ErrorInvalidCardNumber = errors.New("the card number should consist of 16 or 18 digits")
var ErrorInvalidCVV = errors.New("cvv should consist of 3 digits")
var ErrorAlreadyRegistered = errors.New("user with this login already exists, please, log in")
var ErrorLoginNotFound = errors.New("no user with this login, please, register")