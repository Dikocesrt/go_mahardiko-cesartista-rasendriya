package constants

import "errors"

var ErrInsertDatabase error = errors.New("invalid add data in database")
var ErrEmptyInput error = errors.New("name or price cannot be empty")
var ErrUserNotFound error = errors.New("user not found")
var ErrGetDatabase error = errors.New("failed get data from database")
var ErrHitApi error = errors.New("failed hit api")