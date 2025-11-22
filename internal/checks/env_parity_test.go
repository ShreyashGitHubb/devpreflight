package checks

import (
	"testing"

	"github.com/devpreflight/devpreflight/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestEnvParityChecker(t *testing.T) {
	t.Run("constructor creates checker", func(t *testing.T) {
		checker := NewEnvParityChecker()
		assert.NotNil(t, checker)
		assert.Equal(t, "env_parity", checker.Name())
	})
	
	t.Run("checker is enabled by default", func(t *testing.T) {
		checker := NewEnvParityChecker()
		cfg := &config.Config{
			Checks: config.ChecksConfig{
				EnvParity: true,
			},
		}
		assert.True(t, checker.Enabled(cfg))
	})
	
	t.Run("checker is disabled when config is false", func(t *testing.T) {
		checker := NewEnvParityChecker()
		cfg := &config.Config{
			Checks: config.ChecksConfig{
				EnvParity: false,
			},
		}
		assert.False(t, checker.Enabled(cfg))
	})
}