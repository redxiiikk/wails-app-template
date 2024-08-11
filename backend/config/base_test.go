package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestNewApplicationConfig(t *testing.T) {
	tempDir := t.TempDir()
	t.Setenv("HOME", tempDir)

	config, err := NewApplicationConfig("test")()

	dir, _ := os.UserConfigDir()
	assert.NoError(t, err, "NewApplicationConfig() failed: ", err)
	assert.NotNil(t, config, "NewApplicationConfig() is nil")
	assert.Equal(t, "test", config.AppName, "AppName is not correct: ", config.AppName)
	assert.Equal(t, DevEnv, config.Env, "Env is not correct: ", config.AppName)
	assert.Equal(t, filepath.Join(dir, "test-dev"), config.DataDir, "DataDir is not correct: ", config.AppName)
}
