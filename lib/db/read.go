package db

import (
	"encoding/json"
	"fmt"

	driver "github.com/arangodb/go-driver"
)

// ReadCard returns a Card from given Collection with given key
func ReadCard(col driver.Collection, key string) (interface{}, error) {
	var doc = make(map[string]interface{})
	meta, err := col.ReadDocument(nil, key, &doc)
	if err != nil {
		return nil, fmt.Errorf("failed to read card back: %s\nmeta: %s", err.Error(), meta)
	}
	return doc, nil
}

// ReadCardJSON is ReadCard but returns the Card marshalled to JSON
func ReadCardJSON(col driver.Collection, key string) (string, error) {
	doc, err := ReadCard(col, key)
	if err != nil {
		return "", err
	}

	jsonCard, err := json.Marshal(doc)
	if err != nil {
		return "", fmt.Errorf("failed to marshal card: %s", err.Error())
	}

	return string(jsonCard), nil
}
