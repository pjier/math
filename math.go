// The math package contains a few utility functions to do basic arithmetic.
package math

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add adds two numbers and returns the result. A number can be any int or float type.
//
// Read more about addition at [mathisfun].
//
// [mathsisfun]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](i, j T) T {
	return i + j
}
