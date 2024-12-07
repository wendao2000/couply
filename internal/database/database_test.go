package database

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/wendao2000/couply/internal/config"
	constant "github.com/wendao2000/couply/internal/constants"
)

func TestNewDatabase(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	tests := []struct {
		name    string
		cfg     *config.Config
		wantErr bool
	}{
		{
			name: constant.TestSuccess,
			cfg: &config.Config{
				DSN: filepath.Join(tempDir, "couply.db"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewDatabase(tt.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDatabase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
