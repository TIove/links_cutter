package helpers

import "testing"

func TestGetUUID(t *testing.T) {
	var keys [1000]string

	for i := 0; i < 1000; i++ {
		newKey := GetUUID()
		for _, el := range keys {
			if el == newKey {
				t.Error("Duplicates were found")
			}
		}

		keys[i] = newKey
	}
}
