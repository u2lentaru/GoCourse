package qe

import "math/cmplx"

//QuadraticEquation solves the quadratic equation in complex values
func QuadraticEquation(a, b, c complex128) (x1, x2 complex128) {
	sqd := cmplx.Sqrt(complex128(b*b - 4*a*c))
	x1 = (-b + sqd) / (2 * a)
	x2 = (-b - sqd) / (2 * a)
	return x1, x2
}
