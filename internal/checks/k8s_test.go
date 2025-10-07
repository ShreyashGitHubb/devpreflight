package checks

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestK8sChecker(t *testing.T) {
	tests := []struct {
		name     string
		manifest string
		want     bool
		wantErr  bool
	}{
		{
			name: "valid deployment",
			manifest: `apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp
spec:
  replicas: 3
  template:
    spec:
      containers:
      - name: myapp
        image: myapp:1.0.0
        resources:
          limits:
            cpu: "1"
            memory: "512Mi"
          requests:
            cpu: "0.5"
            memory: "256Mi"`,
			want:    true,
			wantErr: false,
		},
		{
			name: "missing resource limits",
			manifest: `apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp
spec:
  replicas: 3
  template:
    spec:
      containers:
      - name: myapp
        image: myapp:1.0.0`,
			want:    false,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checker := NewK8sChecker()
			got, err := checker.Check(tt.manifest)
			
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}