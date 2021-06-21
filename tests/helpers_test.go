package tests

import (
	"links_cutter/helpers"
	"testing"
)

func TestGenerateStrings(t *testing.T) {
	var nums [1000]string

	for i := 0; i < 1000; i++ {
		newUUID := helpers.GetUUID()
		for _, el := range nums {
			if el == newUUID {
				t.Fatal("Generated ID is repeated")
			}
		}
		nums[i] = newUUID
	}
}
