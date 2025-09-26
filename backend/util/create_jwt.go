package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         string `json:"sub"` //user id
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateJwt(secrete string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	byteArrHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	headerB64 := base64UrlEncode(byteArrHeader)

	byteArrData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	payloadB64 := base64UrlEncode(byteArrData)

	// enc := base64.URLEncoding
	// enc = enc.WithPadding(base64.NoPadding)

	// base64Header := enc.EncodeToString(byteArrHeader)
	// base64ByteArrData := enc.EncodeToString(byteArrData)

	message := headerB64 + "." + payloadB64

	byteArrSecrete := []byte(secrete)
	byteArrMessage := []byte(message)

	h := hmac.New(sha256.New, byteArrSecrete)
	h.Write(byteArrMessage)

	signature := h.Sum(nil)
	signatureB64 := base64UrlEncode(signature)

	jwt := headerB64 + "." + payloadB64 + "." + signatureB64

	return jwt, nil

}

func base64UrlEncode(data []byte) string {

	// enc := base64.URLEncoding
	// enc = enc.WithPadding(base64.NoPadding)

	// base64Str := enc.EncodeToString(data)

	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)

}
