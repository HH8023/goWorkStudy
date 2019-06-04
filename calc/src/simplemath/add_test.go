package simplemath

import "testing"

func TestAdd(t *testing.T) {
	r := Add(1,2)
	if r != 1 {
		t.Errorf("Add(1,2) faild. Got %d, expected 3.",r)
	}
}
