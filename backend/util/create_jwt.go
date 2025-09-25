package util

import (
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Playload struct {
	Sub         string `json:"sub"` //user id
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateJwt(data Playload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	byteArrHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	byteArrData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	enc := base64.URLEncoding
	enc = enc.WithPadding(base64.NoPadding)

	base64Header := enc.EncodeToString(byteArrHeader)
	base64ByteArrData := enc.EncodeToString(byteArrData)

}
