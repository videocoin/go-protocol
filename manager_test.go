package protocol_test

import (
	"context"
	"math/big"
	"math/rand"
	"testing"
	"time"

	protocol "github.com/VideoCoin/go-protocol"
	"github.com/stretchr/testify/assert"
)

func TestCtor(t *testing.T) {
	_, err := getManagerClient()
	assert.Nil(t, err, "Failed to create manager client")
}

func TestCtor2(t *testing.T) {
	keyfilePath := "./testdata/local/user"
	pwd := "user"

	_, err := protocol.NewManagerClient(ganache, managerAddr, keyfilePath, pwd)
	assert.NotNil(t, err, "Was expecting error")
}

func TestAddValidator(t *testing.T) {
	m, err := getManagerClient()

	err = m.AddValidator(context.Background(), randAddress())
	assert.Nil(t, err, "Failed to add validator")
}

func TestAddValidator2(t *testing.T) {
	m, err := getManagerClient()

	validator := randAddress()

	err = m.AddValidator(context.Background(), validator)
	assert.Nil(t, err, "Failed to add validator")

	err = m.AddValidator(context.Background(), validator)
	assert.NotNil(t, err, "Was expecting error")
}

func TestRemoveValidator(t *testing.T) {
	m, err := getManagerClient()

	addr := randAddress()

	err = m.AddValidator(context.Background(), addr)
	assert.Nil(t, err, "Failed to add validator")

	err = m.RemoveValidator(context.Background(), addr)
	assert.Nil(t, err, "Failed to remove validator")
}

func TestRemoveValidator2(t *testing.T) {
	m, err := getManagerClient()

	addr := randAddress()

	err = m.RemoveValidator(context.Background(), addr)
	assert.NotNil(t, err, "Was expecting error")
}

func TestApproveStreamCreation(t *testing.T) {
	bitrates := []*big.Int{big.NewInt(1)}
	bitrates = append(bitrates, big.NewInt(1))
	streamId := big.NewInt(int64(getRandInt()))
	RTMP := randStringRunes(10)

	m, err := getManagerClient()
	assert.Nil(t, err, "Failed to get manager client")

	u, err := getUserClient()
	assert.Nil(t, err, "Failed to get user client")

	err = u.RequestStream(context.Background(), streamId, RTMP, bitrates)
	assert.Nil(t, err, "Failed to request a stream")

	chunks := []*big.Int{}
	err = m.ApproveStreamCreation(context.Background(), streamId, chunks)
	assert.Nil(t, err, "Failed to approve a stream")
}

func TestApproveStreamCreation2(t *testing.T) {
	streamId := big.NewInt(int64(getRandInt()))
	chunks := []*big.Int{}

	m, err := getManagerClient()
	assert.Nil(t, err, "Failed to get manager client")

	err = m.ApproveStreamCreation(context.Background(), streamId, chunks)
	assert.NotNil(t, err, "Was expecting error")
}

func TestAllowRefund(t *testing.T) { // & UserClient.ClaimRefund
	streamId := big.NewInt(int64(getRandInt()))

	err := createNewStream(streamId)
	assert.Nil(t, err, "Failed to create a new stream")

	m, err := getManagerClient()
	assert.Nil(t, err, "Failed to get manager client")

	err = m.AllowRefund(context.Background(), streamId)
	assert.Nil(t, err, "Failed to allow refund")

	isAllowed, err := m.RefundAllowed(context.Background(), streamId)
	assert.Nil(t, err, "Failed to allow refund")
	assert.True(t, isAllowed, "Failed to allow refund")

	u, err := getUserClient()
	assert.Nil(t, err, "Failed to get user client")

	err = u.ClaimRefund(context.Background(), streamId)
	assert.Nil(t, err, "Failed to get refund")
}

func TestAllowRefund2(t *testing.T) {
	streamId := big.NewInt(int64(getRandInt()))

	m, err := getManagerClient()
	assert.Nil(t, err, "Failed to get manager client")

	err = m.AllowRefund(context.Background(), streamId)
	assert.NotNil(t, err, "Was expecting error")
}

func TestAllowRefund3(t *testing.T) {
	streamId := big.NewInt(int64(getRandInt()))

	err := createNewStream(streamId)
	assert.Nil(t, err, "Failed to create a new stream")

	m, err := getManagerClient()
	assert.Nil(t, err, "Failed to get manager client")

	err = m.AllowRefund(context.Background(), streamId)
	assert.Nil(t, err, "Failed to allow refund")

	err = m.AllowRefund(context.Background(), streamId)
	assert.NotNil(t, err, "Was expecting error")
}

func TestAddInputChunk(t *testing.T) {
	streamId := big.NewInt(int64(getRandInt()))

	err := createNewStream(streamId)
	assert.Nil(t, err, "Failed to create a new stream")

	m, err := getManagerClient()
	assert.Nil(t, err, "Failed to get manager client")

	err = m.AddInputChunk(context.Background(), streamId, streamId)
	assert.Nil(t, err, "Failed to add input chunk")
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
