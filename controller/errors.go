package controller

import "errors"

var (
	// ErrInternalServerError 内部サーバエラー
	ErrInternalServerError = errors.New("internal server error")
	// ErrUnauthorized 認証失敗
	ErrUnauthorized = errors.New("unauthorized")
)
