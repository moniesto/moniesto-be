package message

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	db "github.com/moniesto/moniesto-be/db/sqlc"
)

func GetMessage(language db.UserLanguage, key Message, inputData ...any) (string, error) {
	// STEP: get correct path of the file
	dir, err := filepath.Abs(filepath.Dir("./"))
	if err != nil {
		return "", err
	}
	filePath := filepath.Join(dir, "util/message/messages/", string(language)+".json")

	// STEP: read JSON file
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// STEP: unmarshal JSON content
	var data map[string]string
	err = json.Unmarshal(fileContent, &data)
	if err != nil {
		return "", err
	}

	// STEP: get value based on key
	value, exists := data[string(key)]
	if !exists {
		return "", fmt.Errorf("key '%s' not found in language '%s'", key, language)
	}

	// STEP: Fill data
	value = fillData(value, inputData...)

	return value, nil
}

func fillData(value string, data ...any) string {
	dataLength := len(data)

	if dataLength == 0 {
		return value
	}

	for i, v := range data {
		replacer := fmt.Sprintf("$%d", i+1) // $1 | $2

		value = strings.ReplaceAll(value, replacer, fmt.Sprint(v))
	}

	return value
}
