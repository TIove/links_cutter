package tests

import (
	"links_cutter/commands"
	"testing"
)

func TestCreateAndGet(t *testing.T) {
	long := "hello"
	short := commands.CreateURL(long)

	newLong, err := commands.GetURL(short)

	if err != nil || newLong != long {
		t.Fatalf("Want %v, but got %v", long, newLong)
	}
}

func TestCreateDifferentUUIDs(t *testing.T) {
	long1 := "hello1"
	long2 := "hello2"
	long3 := "hello3"
	long4 := "hello4"

	short1 := commands.CreateURL(long1)
	short2 := commands.CreateURL(long2)
	short3 := commands.CreateURL(long3)
	short4 := commands.CreateURL(long4)

	if short1 == short2 ||
		short1 == short3 ||
		short1 == short4 ||
		short2 == short3 ||
		short2 == short4 ||
		short3 == short4 {

		t.Fatalf("There are generated equal strings")
	}
}

func TestGetByNonExistentKey(t *testing.T) {
	newLong, err := commands.GetURL("query")

	if err == nil || newLong != "" {
		t.Fatalf("Returned element by non-existent key")
	}
}
