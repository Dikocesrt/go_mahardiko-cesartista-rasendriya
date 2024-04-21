package constants

import "errors"

var ErrInsertDatabase error = errors.New("invalid add data in database")
var ErrEmptyInput error = errors.New("name, email and password cannot be empty")
var ErrUserNotFound error = errors.New("user not found")