package parser_test

import (
	"testing"

	"github.com/ProImpact/fakeapi/parser"
)

func TestGetENdpoints(t *testing.T) {
	p, err := parser.Open("./example.json")
	if err != nil {
		t.Fatal(err)
	}
	if len(*p) != 2 {
		t.Fatalf("expected 2 key got %d", len(*p))
	}
}
