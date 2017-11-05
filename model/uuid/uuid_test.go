package uuid

import (
	"fmt"
	"testing"
)

func TestNewUUID(t *testing.T) {
	uid, err := NewUUID()
	if err != nil {
		fmt.Println("Failed to Create UUID")
	}
	fmt.Println(uid)
}

func TestParseUUID(t *testing.T) {
	uid, err := NewUUID()
	puid, err := ParseUUID(uid)
	if err != nil {
		t.Errorf("Failed to parse UUID")
	}

	if puid.String() != uid {
		t.Errorf("Parsed UUID is not equal to original UUID. Failed")
	}

}
