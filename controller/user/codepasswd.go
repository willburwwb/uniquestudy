package controller

import "encoding/base64"

func Encode(password string) string {
	return string([]byte(base64.StdEncoding.EncodeToString([]byte(password))))
}

func Decode(passwd string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(passwd)
}
