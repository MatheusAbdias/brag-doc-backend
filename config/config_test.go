package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldBeLoadConfig(t *testing.T) {
	config, err := LoadConfig("../")

	assert.NoError(t, err)
	assert.NotEmpty(t, config.PostgresDriver)
	assert.NotEmpty(t, config.PostgresSource)
	assert.NotEmpty(t, config.Port)
}
