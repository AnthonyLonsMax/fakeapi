package util_test

import (
	"testing"

	"github.com/ProImpact/fakeapi/util"
)

func TestPartialUpdate(t *testing.T) {
	tests := []struct {
		name string
		src  map[string]any
		dest map[string]any
	}{
		{
			name: "Simple partial update",
			src: map[string]any{
				"name":   []string{"jose", "hernesto"},
				"status": false,
			},
			dest: map[string]any{
				"name": []string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			util.PartialUpdate(tt.src, tt.dest)
			nameData := tt.dest["name"].([]string)
			if len(nameData) != 2 {
				t.Fatalf("expexted 2 names got %d", len(nameData))
			}
		})
	}
}
