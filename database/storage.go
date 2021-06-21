package database

var urls = make(map[string]string)

func GetValue(key string) (string, error) {
	return urls[key], nil
}

func SetValue(key string, value string) {
	urls[key] = value
}
