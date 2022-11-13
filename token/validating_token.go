package token

import (
	"github.com/moniesto/moniesto-be/core"
	"github.com/moniesto/moniesto-be/util"
)

func CreateValidatingToken() string {
	token := core.CreateID() + "-" + core.CreateID()

	encodedToken := util.Encode(token)

	return encodedToken
}

func GetValidatingToken(encoded_string string) (string, error) {
	decodedToken, err := util.Decode(encoded_string)
	if err != nil {
		return "", err
	}

	return decodedToken, nil

}
