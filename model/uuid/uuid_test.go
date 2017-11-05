package uuid

import (
	"fmt"
	"testing"

	guuid "github.com/satori/go.uuid"
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
	val := puid.Value
	if err != nil {
		t.Errorf("Failed to parse UUID")
	}
	res, _ := guuid.FromBytes(val)

	if res.String() != uid {
		t.Errorf("Parsed UUID is not equal to original UUID. Failed \n From  %v : To  %v", res.String(), puid)
	}

}
