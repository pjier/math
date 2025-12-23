package math

import "testing"

func Test_Add(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)
	}
}

func Test_AddFloat(t *testing.T) {
	result := Add(2.0, 3.0)
	if result != 5.0 {
		t.Error("incorrect result: expected 5, got", result)
	}
}
