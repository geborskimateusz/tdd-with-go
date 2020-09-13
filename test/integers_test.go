package test

import (
	"testing"

	integers "geborskimateusz.com/m/integers"
)

func TestAdder(t *testing.T) {
	a, b, expected := 2, 2, 4

	sum := integers.Add(a, b)

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
