package model

type BaseResponseSuccess struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}