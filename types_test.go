package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	acc, err := NewAccount("clinton", "clinton", "clinton")
	assert.Nil(t, err)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", acc)
}
