package model

type BaseResponseError struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}