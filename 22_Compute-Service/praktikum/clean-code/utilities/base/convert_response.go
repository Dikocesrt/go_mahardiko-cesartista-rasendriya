package base

import (
	"clean/constants"
	"net/http"
)

func ConvertResponseCode(err error) int {
	switch err {
	case constants.ErrInsertDatabase:
		return http.StatusInternalServerError
	case constants.ErrEmptyInput:
		return http.StatusBadRequest
	case constants.ErrUserNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}