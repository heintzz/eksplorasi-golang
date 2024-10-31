package helper

import "testing"

func TestSum(t *testing.T) {
	total := Sum(10, 20, 30, 40, 50)
	expected := 150

	if total != expected {
		t.Errorf("error. expect : %v but got : %v", expected, total)
	}
}
