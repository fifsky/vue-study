package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	assert.NotEmpty(t, App.DB["default"], "Database config unmarshal error")
	assert.NotEmpty(t, App.Log, "Nsq config unmarshal error")
}
