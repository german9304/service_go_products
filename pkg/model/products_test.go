package model

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
		PTests{"socks", New("12334", "socks", 34)},
		PTests{names[1], New("330303", names[1], 100)},
		PTests{names[2], New("112222", names[2], 50)},
	}

	for i, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			if names[i] != v.P.Name {
				t.Errorf("expected product name is %s, but got %s \n", names[i], v.P.Name)
			}
		})
	}
}
