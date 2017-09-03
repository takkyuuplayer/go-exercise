package tests

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestMatchString(t *testing.T) {
	var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

	if validID.MatchString("a") != false {
		t.Fatal(fmt.Sprintf(`validID.MatchString("a") should be
			%#v
			but got
			%#v`, false, validID.MatchString("a")))
	}
}

func TestFindAllString(t *testing.T) {
	re := regexp.MustCompile("a.")

	if !reflect.DeepEqual([]string{"ar", "an", "al"}, re.FindAllString("paranormal", -1)) {
		t.Fatal(fmt.Sprintf(`[]string{"ar", "an", "al"} should be
			%#v
			but got
			%#v`, []string{"ar", "an", "al"}, re.FindAllString("paranormal", -1)))
	}

	if !reflect.DeepEqual([]string{"ar", "an"}, re.FindAllString("paranormal", 2)) {
		t.Fatal(fmt.Sprintf(`[]string{"ar", "an"} should be
			%#v
			but got
			%#v`, []string{"ar", "an"}, re.FindAllString("paranormal", 2)))
	}
}
