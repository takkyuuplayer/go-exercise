package test

import (
	"strconv"
	"testing"
)

func TestByteOperation(t *testing.T) {
	b1, _ := strconv.ParseInt("1100", 2, 0)
	b2, _ := strconv.ParseInt("1010", 2, 0)

	if b1 != 12 {
		t.Errorf(`b1 = %#v, want %#v`, b1, 12)
	}

	if b1&b2 != 8 {
		t.Errorf(`b1 & b2 = %#v, want %#v`, b1&b2, 8)
	}

	if b1|b2 != 14 {
		t.Errorf(`b1|b2 = %#v, want %#v`, b1|b2, 14)
	}

	if b1^b2 != 6 { // XOR
		t.Errorf(`b1 ^ b2 = %#v, want %#v`, b1^b2, 6)
	}

	if b1&^b2 != 4 { // bit clear ( AND NOT )
		t.Errorf(`b1 &^ b2 = %#v, want %#v`, b1&^b2, 4)
	}
}

func TestByteArray(t *testing.T) {
	bt := []byte{1, 2, 3}

	t.Log(len(bt[1:1]) == 0)
}
