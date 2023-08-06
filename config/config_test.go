package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldBeLoadConfig(t *testing.T) {
	config, err := LoadConfig("../")

	assert.NoError(t, err)
	assert.NotEmpty(t, config.Driver)
	assert.NotEmpty(t, config.Source)
	assert.NotEmpty(t, config.Port)
}
