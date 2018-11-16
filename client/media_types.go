// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "ComingServer": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/TinyKitten/ComingServer/design
// --out=$(GOPATH)/src/github.com/TinyKitten/ComingServer
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// 認証成功 (default view)
//
// Identifier: application/vnd.auth.succes+json; view=default
type AuthSucces struct {
	// コード
	ScreenName string `form:"screen_name" json:"screen_name" yaml:"screen_name" xml:"screen_name"`
	// トークン
	Token string `form:"token" json:"token" yaml:"token" xml:"token"`
}

// Validate validates the AuthSucces media type instance.
func (mt *AuthSucces) Validate() (err error) {
	if mt.ScreenName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "screen_name"))
	}
	if mt.Token == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "token"))
	}
	return
}

// DecodeAuthSucces decodes the AuthSucces instance encoded in resp body.
func (c *Client) DecodeAuthSucces(resp *http.Response) (*AuthSucces, error) {
	var decoded AuthSucces
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// ピア (default view)
//
// Identifier: application/vnd.peer+json; view=default
type Peer struct {
	// コード
	Code string `form:"code" json:"code" yaml:"code" xml:"code"`
	// 作成日
	CreatedAt int64 `form:"created_at" json:"created_at" yaml:"created_at" xml:"created_at"`
	// 更新日
	UpdatedAt int64 `form:"updated_at" json:"updated_at" yaml:"updated_at" xml:"updated_at"`
}

// Validate validates the Peer media type instance.
func (mt *Peer) Validate() (err error) {
	if mt.Code == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "code"))
	}

	return
}

// DecodePeer decodes the Peer instance encoded in resp body.
func (c *Client) DecodePeer(resp *http.Response) (*Peer, error) {
	var decoded Peer
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// ピア位置情報 (default view)
//
// Identifier: application/vnd.peer.location+json; view=default
type PeerLocation struct {
	// 作成日
	CreatedAt int64 `form:"created_at" json:"created_at" yaml:"created_at" xml:"created_at"`
	// 緯度
	Latitude float64 `form:"latitude" json:"latitude" yaml:"latitude" xml:"latitude"`
	// 経度
	Longitude float64 `form:"longitude" json:"longitude" yaml:"longitude" xml:"longitude"`
	// 更新日
	UpdatedAt int64 `form:"updated_at" json:"updated_at" yaml:"updated_at" xml:"updated_at"`
}

// Validate validates the PeerLocation media type instance.
func (mt *PeerLocation) Validate() (err error) {

	return
}

// DecodePeerLocation decodes the PeerLocation instance encoded in resp body.
func (c *Client) DecodePeerLocation(resp *http.Response) (*PeerLocation, error) {
	var decoded PeerLocation
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// PeerLocationCollection is the media type for an array of PeerLocation (default view)
//
// Identifier: application/vnd.peer.location+json; type=collection; view=default
type PeerLocationCollection []*PeerLocation

// Validate validates the PeerLocationCollection media type instance.
func (mt PeerLocationCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodePeerLocationCollection decodes the PeerLocationCollection instance encoded in resp body.
func (c *Client) DecodePeerLocationCollection(resp *http.Response) (PeerLocationCollection, error) {
	var decoded PeerLocationCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// PeerCollection is the media type for an array of Peer (default view)
//
// Identifier: application/vnd.peer+json; type=collection; view=default
type PeerCollection []*Peer

// Validate validates the PeerCollection media type instance.
func (mt PeerCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodePeerCollection decodes the PeerCollection instance encoded in resp body.
func (c *Client) DecodePeerCollection(resp *http.Response) (PeerCollection, error) {
	var decoded PeerCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// ポッド (default view)
//
// Identifier: application/vnd.pod+json; view=default
type Pod struct {
	// コード
	Code string `form:"code" json:"code" yaml:"code" xml:"code"`
	// 作成日
	CreatedAt int64 `form:"created_at" json:"created_at" yaml:"created_at" xml:"created_at"`
	// 緯度
	Latitude int64 `form:"latitude" json:"latitude" yaml:"latitude" xml:"latitude"`
	// 経度
	Longitude int64 `form:"longitude" json:"longitude" yaml:"longitude" xml:"longitude"`
	// 鳴っている
	Rumbling bool `form:"rumbling" json:"rumbling" yaml:"rumbling" xml:"rumbling"`
	// 更新日
	UpdatedAt int64 `form:"updated_at" json:"updated_at" yaml:"updated_at" xml:"updated_at"`
}

// Validate validates the Pod media type instance.
func (mt *Pod) Validate() (err error) {
	if mt.Code == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "code"))
	}

	return
}

// DecodePod decodes the Pod instance encoded in resp body.
func (c *Client) DecodePod(resp *http.Response) (*Pod, error) {
	var decoded Pod
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// PodCollection is the media type for an array of Pod (default view)
//
// Identifier: application/vnd.pod+json; type=collection; view=default
type PodCollection []*Pod

// Validate validates the PodCollection media type instance.
func (mt PodCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodePodCollection decodes the PodCollection instance encoded in resp body.
func (c *Client) DecodePodCollection(resp *http.Response) (PodCollection, error) {
	var decoded PodCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// トークン (default view)
//
// Identifier: application/vnd.token+json; view=default
type Token struct {
	// 対象ID
	ID int64 `form:"id" json:"id" yaml:"id" xml:"id"`
	// トークン
	Token string `form:"token" json:"token" yaml:"token" xml:"token"`
	// 更新日
	UpdatedAt int64 `form:"updated_at" json:"updated_at" yaml:"updated_at" xml:"updated_at"`
}

// Validate validates the Token media type instance.
func (mt *Token) Validate() (err error) {

	if mt.Token == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "token"))
	}

	return
}

// DecodeToken decodes the Token instance encoded in resp body.
func (c *Client) DecodeToken(resp *http.Response) (*Token, error) {
	var decoded Token
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// ユーザー (default view)
//
// Identifier: application/vnd.user+json; view=default
type User struct {
	// 作成日
	CreatedAt int64 `form:"created_at" json:"created_at" yaml:"created_at" xml:"created_at"`
	// ID
	ID int64 `form:"id" json:"id" yaml:"id" xml:"id"`
	// スクリーンネーム
	ScreenName string `form:"screen_name" json:"screen_name" yaml:"screen_name" xml:"screen_name"`
	// 更新日
	UpdatedAt int64 `form:"updated_at" json:"updated_at" yaml:"updated_at" xml:"updated_at"`
}

// Validate validates the User media type instance.
func (mt *User) Validate() (err error) {

	if mt.ScreenName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "screen_name"))
	}

	return
}

// DecodeUser decodes the User instance encoded in resp body.
func (c *Client) DecodeUser(resp *http.Response) (*User, error) {
	var decoded User
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// UserCollection is the media type for an array of User (default view)
//
// Identifier: application/vnd.user+json; type=collection; view=default
type UserCollection []*User

// Validate validates the UserCollection media type instance.
func (mt UserCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeUserCollection decodes the UserCollection instance encoded in resp body.
func (c *Client) DecodeUserCollection(resp *http.Response) (UserCollection, error) {
	var decoded UserCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}
