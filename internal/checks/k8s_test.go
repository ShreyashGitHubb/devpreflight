package checks

import (
	"testing"

	"github.com/devpreflight/devpreflight/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestK8sChecker(t *testing.T) {
	t.Run("constructor creates checker", func(t *testing.T) {
		checker := NewK8sChecker()
		assert.NotNil(t, checker)
		assert.Equal(t, "k8s_validate", checker.Name())
	})
	
	t.Run("checker is enabled by default", func(t *testing.T) {
		checker := NewK8sChecker()
		cfg := &config.Config{
			Checks: config.ChecksConfig{
				K8sValidate: true,
			},
		}
		assert.True(t, checker.Enabled(cfg))
	})
	
	t.Run("checker is disabled when config is false", func(t *testing.T) {
		checker := NewK8sChecker()
		cfg := &config.Config{
			Checks: config.ChecksConfig{
				K8sValidate: false,
			},
		}
		assert.False(t, checker.Enabled(cfg))
	})
}