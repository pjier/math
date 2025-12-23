package math

import "testing"

func Test_Add(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)
	}
}
