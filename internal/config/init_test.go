package config

import (
	"reflect"
	"testing"

	constant "github.com/wendao2000/couply/internal/constants"
)

func TestInitConfig(t *testing.T) {
	tests := []struct {
		name    string
		want    *Config
		wantErr bool
	}{
		{
			name:    constant.TestSuccess,
			want:    &Config{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InitConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("InitConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
