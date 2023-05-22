package service

import "errors"

// errors
var (
	ErrPasswordMustBeMinLen      = errors.New("your password must be at least 8 characters")
	ErrPasswordsDoNotMatch       = errors.New("passwords do not match")
	ErrAlreadyRegistered         = errors.New("you have already registered")
	ErrLoginInformationInCorrect = errors.New("your login information is incorrect")
)

const PasswordMinimumLength = 8
