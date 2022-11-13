package util

import (
	b64 "encoding/base64"
)

func Encode(str string) string {
	return singleEncode(singleEncode(str))
}

func Decode(str string) (string, error) {
	decoded1, err := singleDecode(str)
	if err != nil {
		return "", err
	}

	decoded2, err := singleDecode(decoded1)
	if err != nil {
		return "", err
	}

	return decoded2, nil
}

func singleEncode(str string) string {
	sEnc := b64.StdEncoding.EncodeToString([]byte(str))

	return sEnc
}

func singleDecode(str string) (string, error) {
	sDec, err := b64.StdEncoding.DecodeString(str)

	return string(sDec), err
}
