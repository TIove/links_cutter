package database

var urls = make(map[string] string)

func GetValue(key string) string {
	return urls[key]
}

func SetValue(key string, value string) {
	urls[key] = value
}
