package pkg_test

import (
	"testing"

	"github.com/ProImpact/fakeapi/pkg"
)

func TestSortMap(t *testing.T) {
	tests := []struct {
		name   string // description of this test case
		values []map[string]any
		key    string
	}{
		{
			name: "test int",
			values: []map[string]any{
				{
					"age": 21,
				},
				{
					"age": -3,
				},
			},
			key: "age",
		},
		{
			name: "test float",
			values: []map[string]any{
				{
					"age": 21.34,
				},
				{
					"age": -3.21,
				},
			},
			key: "age",
		},
		{
			name: "test string",
			values: []map[string]any{
				{
					"age": "baa",
				},
				{
					"age": "aaa",
				},
			},
			key: "age",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pkg.SortMap(tt.values, tt.key)
			t.Log(tt.values)
		})
	}
}
