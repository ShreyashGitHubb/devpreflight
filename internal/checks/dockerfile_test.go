package checks

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDockerfileChecker(t *testing.T) {
	tests := []struct {
		name       string
		dockerfile string
		want       bool
		wantErr    bool
	}{
		{
			name: "valid dockerfile",
			dockerfile: `FROM golang:1.24-alpine
WORKDIR /app
COPY . .
RUN go build -o main .
CMD ["./main"]`,
			want:    true,
			wantErr: false,
		},
		{
			name: "missing FROM instruction",
			dockerfile: `WORKDIR /app
COPY . .
RUN go build -o main .
CMD ["./main"]`,
			want:    false,
			wantErr: true,
		},
		{
			name: "using latest tag",
			dockerfile: `FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o main .
CMD ["./main"]`,
			want:    false,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checker := NewDockerfileChecker()
			got, err := checker.Check(tt.dockerfile)
			
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}