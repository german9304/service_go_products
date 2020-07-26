package model

import (
	//"log"
	"testing"
)

type PTests struct {
	Name string
	P    Product
}

func TestProduct(t *testing.T) {
	names := []string{"socks", "shoes", "computer"}
	tests := []PTests{
		PTests{"socks", New("socks", 1, 34)},
		PTests{names[1], New(names[1], 2, 100)},
		PTests{names[2], New(names[2], 3, 50)},
	}

	for i, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			if names[i] != v.P.Name {
				t.Errorf("expected product name is %s, but got %s \n", names[i], v.P.Name)
			}
		})
	}
}
