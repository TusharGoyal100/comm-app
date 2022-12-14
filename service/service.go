package service

import (
	"communication-app/connector"
	"encoding/json"
)

func CreateFile() (string, error) {
	allMessages, err := connector.GetAll()
	if err != nil {
		return "", err
	}

	content, err := json.Marshal(allMessages)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
