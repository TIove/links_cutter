package tests

import (
	"links_cutter/database"
	"testing"
)

func TestCreateAndGetFormDB(t *testing.T) {
	long := "hello"
	short := "qwRWms2_fW"
	database.SetValue(short, long)

	newLong, err := database.GetValue(short)

	if err != nil || newLong != long {
		t.Fatalf("Want %v, but got %v", long, newLong)
	}
}

func TestGetByNonExistentKeyFromDb(t *testing.T) {
	newLong, err := database.GetValue("query")

	if err == nil || newLong != "" {
		t.Fatalf("Returned element by non-existent key")
	}
}
