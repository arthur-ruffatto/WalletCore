package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John", "h0f7P@example.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John", client.Name)
	assert.Equal(t, "h0f7P@example.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John", "h0f7P@example.com")
	err := client.Update("John Updated", "h0f7P@example2.com")
	assert.Nil(t, err)
	assert.Equal(t, "John Updated", client.Name)
	assert.Equal(t, "h0f7P@example2.com", client.Email)
}

func TestUpdateClienteWhenArgsAreInvalid(t *testing.T) {
	client, _ := NewClient("John", "h0f7P@example.com")
	err := client.Update("", "")
	assert.NotNil(t, err)
	assert.Equal(t, "name is required", err.Error())
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("John", "h0f7P@example.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Account))
}
