package clientError

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var splittedWith string = "*_*"

func CreateError(http_code int, error_code string) error {
	error_str := strconv.Itoa(http_code) + splittedWith + error_code + splittedWith + errorMessages[error_code]

	return errors.New(error_str)
}

func ParseError(err error) (int, gin.H) {
	error_message := err.Error()
	messages := strings.Split(error_message, splittedWith)

	code, _ := strconv.Atoi(messages[0])

	return code, gin.H{"error_code": messages[1], "error": messages[2]}
}

func GetError(error_code string) gin.H {
	return gin.H{
		"error_code": error_code,
		"error":      errorMessages[error_code],
	}
}
