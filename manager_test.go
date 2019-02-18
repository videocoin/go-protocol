package protocol_test

import (
	"context"
	"math/big"
	"math/rand"
	"testing"
	"time"

	protocol "github.com/VideoCoin/go-protocol"
)

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
	m, err := getManager()

	addr := randAddress()

	err = m.AddValidator(context.Background(), addr)
	if err != nil {
		t.Errorf("Failed to add validator: %s", err.Error())
	}

	err = m.RemoveValidator(context.Background(), addr)
	if err != nil {
		t.Errorf("Failed to remove validator: %s", err.Error())
	}

}

func TestApproveStreamCreation(t *testing.T) {
	bitrates := []*big.Int{big.NewInt(1)}
	bitrates = append(bitrates, big.NewInt(1))
	streamId := big.NewInt(int64(rand.Int()))
	RTMP := randStringRunes(10)

	m, err := getManager()
	if err != nil {
		t.Errorf("Failed to get manager client: %s", err.Error())
	}

	u, err := getUser()
	if err != nil {
		t.Errorf("Failed to get user client: %s", err.Error())
	}

	err = u.RequestStream(context.Background(), streamId, RTMP, bitrates)
	if err != nil {
		t.Errorf("Failed to request a stream: %s", err.Error())
	}

	chunks := []*big.Int{}

	err = m.ApproveStreamCreation(context.Background(), streamId, chunks)
	if err != nil {
		t.Errorf("Failed to approve a stream: %s", err.Error())
	}
}

func TestAllowRefund(t *testing.T) {
	// getNewStream()
}

// utils
var manager *protocol.ManagerClient = nil
var user *protocol.UserClient = nil

const ganache = "http://localhost:8545"
const managerAddr = "0xB20eBC07c7804E4c3B273aFc8Cf76Ae1F7a71618"

func randAddress() string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("1234567890abcdefABCDEF")

	b := make([]rune, 40)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return "0x" + string(b)
}

func randStringRunes(n int) string {

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func getManager() (*protocol.ManagerClient, error) {
	keyfilePath := "./testdata/local/manager"
	pwd := "local"

	if manager == nil {
		m, err := protocol.NewManagerClient(ganache, managerAddr, keyfilePath, pwd)
		manager = m

		return m, err
	} else {
		return manager, nil
	}
}

func getUser() (*protocol.UserClient, error) {
	keyfilePath := "./testdata/local/manager"
	pwd := "local"

	if user == nil {
		u, err := protocol.NewUserClient(ganache, managerAddr, keyfilePath, pwd)
		user = u

		return u, err
	} else {
		return user, nil
	}
}

func getNewStream() error {
	bitrates := []*big.Int{big.NewInt(1)}
	bitrates = append(bitrates, big.NewInt(1))
	streamId := big.NewInt(int64(rand.Int()))
	RTMP := randStringRunes(10)

	m, err := getManager()
	if err != nil {
		return err
	}

	u, err := getUser()
	if err != nil {
		return err
	}

	err = u.RequestStream(context.Background(), streamId, RTMP, bitrates)

	chunks := []*big.Int{}

	err = m.ApproveStreamCreation(context.Background(), streamId, chunks)
	if err != nil {
		return err
	}

	return nil
}
