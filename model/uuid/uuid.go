package uuid

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"log"
	puuid "tracker/proto/uuid"

	guuid "github.com/satori/go.uuid"
	logger "github.com/sirupsen/logrus"
)

//TODO: Swtich UUID packages
type UUID struct{}

func BytesToString(uid []uint8) (string, error) {
	if uid == nil {
		return "", errors.New("nil")
	}
	uu, err := guuid.FromBytes(uid)
	if err != nil {
		logger.Warning("Failed to provide good UUID to convert to string", err.Error())
	}
	return uu.String(), nil
}

func ParseUUID(uid string) (val puuid.UUID, err error) {
	ret, err := guuid.FromString(uid)
	if err != nil {
		return puuid.UUID{}, err
	}
	var uu puuid.UUID
	uu.Value = ret.Bytes()
	return uu, nil
}

func ParseUUIDFromBytes(uid []uint8) (val puuid.UUID, err error) {
	ret, err := guuid.FromBytes(uid)

	if err != nil {
		log.Panic("Failed to convert UUID from Resonse into bytes: \n Bytes: ", uid)
		return puuid.UUID{}, err
	}
	var uu puuid.UUID
	uu.Value = ret.Bytes()
	return uu, nil
}

// newUUID generates a random UUID according to RFC 4122
func NewUUID() ([]byte, error) {
	ret := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, ret)
	if n != len(ret) || err != nil {
		return nil, err
	}
	// variant bits; see section 4.1.1
	ret[8] = ret[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	ret[6] = ret[6]&^0xf0 | 0x40
	return ret, nil
	//return fmt.Sprintf("%x-%x-%x-%x-%x", ret[0:4], ret[4:6], ret[6:8], ret[8:10], ret[10:]), nil
}

func NewUUIDString() (string, error) {
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
