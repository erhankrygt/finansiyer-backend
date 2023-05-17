package service

import "errors"

// errors
var (
	ErrPasswordMustBeMinLen = errors.New("your password must be at least " + string(PasswordMinimumLength) + " characters")
	ErrPasswordsDoNotMatch  = errors.New("passwords do not match")
)

const PasswordMinimumLength = 8
