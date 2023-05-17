package service

import "errors"

// errors
var (
	ErrPasswordMustBeMinLen   = errors.New("your password must be at least 8 characters")
	ErrPasswordsDoNotMatch    = errors.New("passwords do not match")
	ErrAlreadyUsedPhoneNumber = errors.New("already used with this number")
)

const PasswordMinimumLength = 8
