package tests

import (
	"links_cutter/app"
	"testing"
)

func TestCreateAndGetFormDB(t *testing.T) {
	long := "hello"
	short := "qwRWms2_fW"
	app.SetValue(short, long)

	newLong, err := app.GetValue(short)

	if err != nil || newLong != long {
		t.Fatalf("Want %v, but got %v", long, newLong)
	}
}

func TestGetByNonExistentKeyFromDb(t *testing.T) {
	newLong, err := app.GetValue("query")

	if err == nil || newLong != "" {
		t.Fatalf("Returned element by non-existent key")
	}
}
