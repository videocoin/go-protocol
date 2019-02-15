package protocol_test

import (
	"context"
	"math/rand"
	"testing"
	"time"

	protocol "github.com/VideoCoin/go-protocol"
)

var manager *protocol.ManagerClient = nil

func TestCtor(t *testing.T) {
	_, err := getManager()
	if err != nil {
		t.Errorf("Failed to create manager client: %s", err.Error())
	}
}

func TestAddValidator(t *testing.T) {
	m, err := getManager()

	err = m.AddValidator(context.Background(), randAddress())
	if err != nil {
		t.Errorf("Failed to add validator: %s", err.Error())
	}
}

func TestRemoveValidator(t *testing.T) {

}

func TestApproveStreamCreation(t *testing.T) {

}

func randAddress() string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("1234567890abcdefABCDEF")

	b := make([]rune, 40)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return "0x" + string(b)
}

func getManager() (*protocol.ManagerClient, error) {
	ganache := "http://localhost:8545"
	managerAddr := "0xB20eBC07c7804E4c3B273aFc8Cf76Ae1F7a71618"
	keyfilePath := "./testdata/local"
	pwd := "local"

	if manager == nil {
		m, err := protocol.NewManagerClient(ganache, managerAddr, keyfilePath, pwd)
		manager = m

		return m, err
	} else {
		return manager, nil
	}
}
