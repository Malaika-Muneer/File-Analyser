package errors

import (
	"errors"
)

var (
	ErrMissusernameorpassword = errors.New("username and password required")
	ErrUserExist              = errors.New("User_exist")
	ErrDbError                = errors.New("database error")
	ErrHashpasswordfailed     = errors.New("Password_Hashing_failed")
	ErrUserInsertion          = errors.New("error inserting user")
	ErrMissingUserCredential  = errors.New("missing User Credentials")
	ErrInvalidCredental       = errors.New("invalid Credentials")
	ErrPasswordFetchingFailed = errors.New("fetching Password Failed")
	ErrInvalidToken           = errors.New("invalid token")
)
