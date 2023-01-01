package token

import (
	"github.com/moniesto/moniesto-be/core"
)

func CreateValidatingToken() string {
	token := core.CreateID() + "-" + core.CreateID()

	return token
}

func EncodeValidatingToken(token string) string {
	encodedToken := core.Encode(token)

	return encodedToken
}

func GetValidatingToken(encoded_string string) (string, error) {
	decodedToken, err := core.Decode(encoded_string)
	if err != nil {
		return "", err
	}

	return decodedToken, nil

}
