package handlers

import "errors"

var ErrIncorrectUserRegistrationInfo = errors.New("incorrect user registration info")
var ErrNoToken = errors.New("no token provided")
