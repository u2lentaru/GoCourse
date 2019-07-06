package qe

import "testing"

func TestQuadraticEquation(t *testing.T) {
	var v1, v2 complex128
	v1, v2 = QuadraticEquation(1, 2, 3)
	if v1 != 1 && v2 != 1 {
		t.Error("Expected 1,2 got", v1, v2)
	}
}
