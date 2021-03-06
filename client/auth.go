// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "ComingServer": auth Resource Client
//
// Command:
// $ goagen
// --design=github.com/TinyKitten/ComingServer/design
// --out=$(GOPATH)/src/github.com/TinyKitten/ComingServer
// --version=v1.3.1

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// AuthAuthPayload is the auth auth action payload.
type AuthAuthPayload struct {
	// パスワード
	Password string `form:"password" json:"password" yaml:"password" xml:"password"`
	// スクリーンネーム
	ScreenName string `form:"screen_name" json:"screen_name" yaml:"screen_name" xml:"screen_name"`
}

// AuthAuthPath computes a request path to the auth action of auth.
func AuthAuthPath() string {

	return fmt.Sprintf("/v1/auth")
}

// JWTトークンを生成する
func (c *Client) AuthAuth(ctx context.Context, path string, payload *AuthAuthPayload, contentType string) (*http.Response, error) {
	req, err := c.NewAuthAuthRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAuthAuthRequest create the request corresponding to the auth action endpoint of the auth resource.
func (c *Client) NewAuthAuthRequest(ctx context.Context, path string, payload *AuthAuthPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
