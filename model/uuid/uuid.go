package uuid

import (
	"crypto/rand"
	"fmt"
	"io"

	guuid "github.com/google/uuid"
)

type UUID struct{}

func ParseUUID(uid string) (val guuid.UUID, err error) {
	ret, err := guuid.Parse(uid)
	if err != nil {
		return guuid.UUID{}, err
	}
	return ret, nil
}

// newUUID generates a random UUID according to RFC 4122
func NewUUID() (string, error) {
	ret := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, ret)
	if n != len(ret) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	ret[8] = ret[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	ret[6] = ret[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", ret[0:4], ret[4:6], ret[6:8], ret[8:10], ret[10:]), nil
}

func (u *UUID) String() {

}
