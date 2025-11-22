package checks

import (
	"testing"

	"github.com/devpreflight/devpreflight/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestDockerfileChecker(t *testing.T) {
	t.Run("constructor creates checker", func(t *testing.T) {
		checker := NewDockerfileChecker()
		assert.NotNil(t, checker)
		assert.Equal(t, "dockerfile_lint", checker.Name())
	})
	
	t.Run("checker is enabled by default", func(t *testing.T) {
		checker := NewDockerfileChecker()
		cfg := &config.Config{
			Checks: config.ChecksConfig{
				DockerfileLint: true,
			},
		}
		assert.True(t, checker.Enabled(cfg))
	})
	
	t.Run("checker is disabled when config is false", func(t *testing.T) {
		checker := NewDockerfileChecker()
		cfg := &config.Config{
			Checks: config.ChecksConfig{
				DockerfileLint: false,
			},
		}
		assert.False(t, checker.Enabled(cfg))
	})
}