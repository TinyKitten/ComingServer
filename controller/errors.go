package controller

import "errors"

var (
	// ErrUnauthorized 認証失敗
	ErrUnauthorized = errors.New("unauthorized")
)
