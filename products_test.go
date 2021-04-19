package server

import (
	"testing"
)

type PTests struct {
	Name string
	P    Product
}

func TestProduct(t *testing.T) {
	names := []string{"socks", "shoes", "computer"}
	tests := []PTests{
		{"socks", Product{"12334", "socks", 34}},
		{names[1], Product{"330303", names[1], 100}},
		{names[2], Product{"112222", names[2], 50}},
	}

	for i, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			if names[i] != v.P.Name {
				t.Errorf("expected product name is %s, but got %s \n", names[i], v.P.Name)
			}
		})
	}
}
