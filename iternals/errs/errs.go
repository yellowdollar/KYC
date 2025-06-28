package errs

import "errors"

var (
	ErrUserIDNotFound    = errors.New("user ID not found")
	ErrUserAlreadyExists = errors.New("user with this phone number already exists")
)
