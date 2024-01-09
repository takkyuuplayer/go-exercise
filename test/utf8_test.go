package test

import (
	"testing"
	"unicode"
	"unicode/utf8"
)

func TestUTFMax(t *testing.T) {
	t.Parallel()
	if utf8.UTFMax != 4 { // 4 byte
		t.Errorf(`utf8.UTFMax = %#v, want %#v`, utf8.UTFMax, 4)
	}
}

func TestRuneLen(t *testing.T) {
	t.Parallel()
	if utf8.RuneLen(' ') != 1 {
		t.Errorf(`utf8.RuneLen(' ') = %#v, want %#v`, utf8.RuneLen(' '), 1)
	}

	if utf8.RuneLen('　') != 3 {
		t.Errorf(`utf8.RuneLen('　') = %#v, want %#v`, utf8.RuneLen('　'), 3)
	}

	if utf8.RuneLen('あ') != 3 {
		t.Errorf(`utf8.RuneLen('あ') = %#v, want %#v`, utf8.RuneLen('あ'), 3)
	}
}

func TestIsSpace(t *testing.T) {
	t.Parallel()
	if unicode.IsSpace(' ') != true {
		t.Errorf(`unicode.IsSpace(' ") = %#v, want %#v`, unicode.IsSpace(' '), true)
	}

	if unicode.IsSpace('\n') != true {
		t.Errorf(`unicode.IsSpace('\n') = %#v, want %#v`, unicode.IsSpace('\n'), true)
	}

	if unicode.IsSpace('　') != true {
		t.Errorf(`unicode.IsSpace('　') = %#v, want %#v`, unicode.IsSpace('　'), true)
	}
}
