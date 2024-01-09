package test

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
)

func TestLen(t *testing.T) {
	t.Parallel()
	s := "hello, world"

	if len(s) != 12 {
		t.Errorf(`len(s) = %#v, want %#v`, len(s), 12)
	}

	if s[0] != 'h' {
		t.Errorf(`s[0] = %#v, want %#v`, s[0], 'h')
	}

	if s[0] != 104 {
		t.Errorf(`s[0] = %#v, want %#v`, s[0], 104)
	}
}

func TestSubString(t *testing.T) {
	t.Parallel()
	s := "hello, world"

	if s[:5] != "hello" {
		t.Errorf(`s[:5] = %#v, want %#v`, s[:5], "hello")
	}
}

func TestStringModification(t *testing.T) {
	t.Parallel()
	a := "left"
	b := a
	a += ", right"

	if a != "left, right" {
		t.Errorf(`a = %#v, want %#v`, a, "left")
	}

	if b != "left" {
		t.Errorf(`b = %#v, want %#v`, b, "left")
	}
}

func TestByteCode(t *testing.T) {
	t.Parallel()
	if "\xff" != "\377" {
		t.Errorf(`"\xff" = %#v, want %#v`, "\xff", "\377")
	}

	if "\xff" != "\377" {
		t.Errorf(`"\xff" = %#v, want %#v`, "\xff", "\377")
	}

}

func TestStringToByte(t *testing.T) {
	t.Parallel()
	org := "abcd"
	a := []byte(org)
	copy(a[1:2], []byte(" "))

	if !reflect.DeepEqual([]byte(org), []byte("abcd")) {
		t.Errorf(`[]byte(org) = %#v, want %#v`, []byte(org), []byte("abcd"))
	}

	if !reflect.DeepEqual(a, []byte("a cd")) {
		t.Errorf(`a = %#v, want %#v`, a, []byte("a cd"))
	}
}

func TestUnicode(t *testing.T) {
	t.Parallel()
	if "\xe4\xb8\x96\xe7\x95\x8c" != "世界" {
		t.Errorf(`"\xe4\xb8\x96\xe7\x95\x8c" = %#v, want %#v`, "\xe4\xb8\x96\xe7\x95\x8c", "世界")
	}

	if "\u4e16\u754c" != "世界" {
		t.Errorf(`"\u4e16\u754c" = %#v, want %#v`, "\u4e16\u754c", "世界")
	}

	if "\U00004e16\U0000754c" != "世界" {
		t.Errorf(`"\U00004e16\U0000754c" = %#v, want %#v`, "\U00004e16\U0000754c", "世界")
	}
}

func TestUTF8(t *testing.T) {
	t.Parallel()
	s := "Hello, 世界"

	if len(s) != 13 {
		t.Errorf(`len(s) = %#v, want %#v`, len(s), 13)
	}

	if utf8.RuneCountInString(s) != 9 {
		t.Errorf(`utf8.RuneCountInString(s) = %#v, want %#v`, utf8.RuneCountInString(s), 9)
	}

	s2 := ""
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		s2 += fmt.Sprintf("%c", r)
		i += size
	}

	if s != s2 {
		t.Errorf(`s = %#v, want %#v`, s, s2)
	}
}

func TestForLoop(t *testing.T) {
	t.Parallel()
	if fmt.Sprintf("%T", 'a') != "int32" {
		t.Errorf(`Type of 'a' = %T, want %s`, 'a', "int32")
	}

	s := "Hello 世界"

	for _, v := range s {
		if fmt.Sprintf("%T", v) != "int32" {
			t.Errorf(`Type of v = %T, want %s`, v, "int32")
		}
	}
}

func TestTrim(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "Hello World", strings.Trim(" Hello World  ", " "))
}

var name = strings.Repeat(".", 100)
var name2 = strings.Repeat(".", 100) + "あ"

func BenchmarkStringsRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Repeat(".", len(name)) == name
		_ = strings.Repeat(".", len(name2)) == name
	}
}

var reg = regexp.MustCompile(`\A\.+\z`)

func BenchmarkRegex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = reg.MatchString(name)
		_ = reg.MatchString(name2)
	}
}
