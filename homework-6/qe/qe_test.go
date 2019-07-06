package qe

import "testing"

func TestQuadraticEquation(t *testing.T) {
	var v1, v2 complex128
	v1, v2 = QuadraticEquation(1, 2, 3)
	if real(v1) != 1 && imag(v1) != 1 && real(v2) != 1 && imag(v2) != -1 {
		t.Error("Expected 1+i,1-i got", v1, v2)
	}
}

//Running tool: C:\Go\bin\go.exe test -timeout 30s GoCourse\homework-6\qe -run ^(TestQuadraticEquation)$
//
//--- FAIL: TestQuadraticEquation (0.00s)
//c:\Golang_work\src\GoCourse\homework-6\qe\qe_test.go:9: Expected 1+i,1-i got (-1+1.4142135623730951i) (-1-1.4142135623730951i)
//FAIL
//FAIL	GoCourse/homework-6/qe	1.186s
//Error: Tests failed.
