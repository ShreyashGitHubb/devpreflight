package checks

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEnvParityChecker(t *testing.T) {
	tests := []struct {
		name     string
		envFiles map[string][]string
		want     bool
		wantErr  bool
	}{
		{
			name: "matching environments",
			envFiles: map[string][]string{
				"dev": {"APP_KEY", "DB_HOST"},
				"prod": {"APP_KEY", "DB_HOST"},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "missing env var in prod",
			envFiles: map[string][]string{
				"dev": {"APP_KEY", "DB_HOST", "DEBUG"},
				"prod": {"APP_KEY", "DB_HOST"},
			},
			want:    false,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checker := NewEnvParityChecker()
			got, err := checker.Check(tt.envFiles)
			
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}