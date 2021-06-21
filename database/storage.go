package database

import "errors"

var urls = make(map[string]string)

func GetValue(key string) (string, error) {
	val, ok := urls[key]
	if !ok {
		return val, errors.New("this key doesn't exist")
	}

	return val, nil
}

func SetValue(key string, value string) {
	urls[key] = value
}
