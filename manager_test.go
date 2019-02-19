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
	_, err := getManagerClient()
	if err != nil {
		t.Errorf("Failed to create manager client: %s", err.Error())
	}
}

func TestCtor2(t *testing.T) {
	keyfilePath := "./testdata/local/user"
	pwd := "user"
	_, err := protocol.NewManagerClient(ganache, managerAddr, keyfilePath, pwd)
	if err == nil {
		t.Errorf("Was expecting error")
	}
}

func TestAddValidator(t *testing.T) {
	m, err := getManagerClient()

	err = m.AddValidator(context.Background(), randAddress())
	if err != nil {
		t.Errorf("Failed to add validator: %s", err.Error())
	}
}

func TestRemoveValidator(t *testing.T) {
	m, err := getManagerClient()

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
	streamId := big.NewInt(int64(getRandInt()))
	RTMP := randStringRunes(10)

	m, err := getManagerClient()
	if err != nil {
		t.Errorf("Failed to get manager client: %s", err.Error())
	}

	u, err := getUserClient()
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

func TestAllowRefund(t *testing.T) { // & UserClient.ClaimRefund
	streamId := big.NewInt(int64(getRandInt()))

	err := createNewStream(streamId)
	if err != nil {
		t.Errorf("Failed to create a new stream: %s", err.Error())
	}

	m, err := getManagerClient()
	if err != nil {
		t.Errorf("Failed to get manager client: %s", err.Error())
	}

	err = m.AllowRefund(context.Background(), streamId)
	if err != nil {
		t.Errorf("Failed to allow refund: %s", err.Error())
	}

	u, err := getUserClient()
	if err != nil {
		t.Errorf("Failed to get user client: %s", err.Error())
	}

	err = u.ClaimRefund(context.Background(), streamId)
	if err != nil {
		t.Errorf("Failed to get refund: %s", err.Error())
	}
}

func TestAddInputChunk(t *testing.T) {
	streamId := big.NewInt(int64(getRandInt()))

	err := createNewStream(streamId)
	if err != nil {
		t.Errorf("Failed to create a new stream: %s", err.Error())
	}

	m, err := getManagerClient()
	if err != nil {
		t.Errorf("Failed to get manager client: %s", err.Error())
	}

	err = m.AddInputChunk(context.Background(), streamId, streamId)
	if err != nil {
		t.Errorf("Failed to add input chunk: %s", err.Error())
	}
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
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func getManagerClient() (*protocol.ManagerClient, error) {
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

func getUserClient() (*protocol.UserClient, error) {
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

func createNewStream(streamID *big.Int) error {
	bitrates := []*big.Int{big.NewInt(1)}
	bitrates = append(bitrates, big.NewInt(1))
	RTMP := randStringRunes(10)

	m, err := getManagerClient()
	if err != nil {
		return err
	}

	u, err := getUserClient()
	if err != nil {
		return err
	}

	err = u.RequestStream(context.Background(), streamID, RTMP, bitrates)
	if err != nil {
		return err
	}

	chunks := []*big.Int{}

	err = m.ApproveStreamCreation(context.Background(), streamID, chunks)
	if err != nil {
		return err
	}

	funds := new(big.Int).Exp(big.NewInt(10), big.NewInt(19), nil) // send funds

	err = u.CreateStream(context.Background(), streamID, funds)
	if err != nil {
		return err
	}

	return nil
}

func getRandInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}
