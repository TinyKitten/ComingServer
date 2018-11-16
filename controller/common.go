package controller

import (
	"strconv"
	"time"

	hashids "github.com/speps/go-hashids"
)

func generateToken() (string, error) {
	now := time.Now()
	hd := hashids.NewData()
	hd.Salt = strconv.FormatInt(now.Unix(), 10)
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}
	id, err := h.Encode([]int{1, 2, 3})
	if err != nil {
		return "", err
	}
	return id, err
}
