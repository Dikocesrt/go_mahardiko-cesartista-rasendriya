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
	case constants.ErrGetAllDatabase:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}