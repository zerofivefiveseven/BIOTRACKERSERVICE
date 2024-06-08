package usecases

import "errors"

var ErrUserAlreadyExists = errors.New("user with this email already exists")
var ErrUserNotFound = errors.New("user not found")
