package tasks

import (
	"context"
	"testing"
)

func TestDiscoveryTaskHandler(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DiscoveryTaskHandler(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("DiscoveryTaskHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
