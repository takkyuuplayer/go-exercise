package test

import (
	"reflect"
	"regexp"
	"testing"
)

func TestMatchString(t *testing.T) {
	var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

	if validID.MatchString("a") != false {
		t.Fatalf(`validID.MatchString("a") should be
			%#v
			but got
			%#v`, false, validID.MatchString("a"))
	}
}

func TestFindAllString(t *testing.T) {
	re := regexp.MustCompile("a.")

	if !reflect.DeepEqual([]string{"ar", "an", "al"}, re.FindAllString("paranormal", -1)) {
		t.Fatalf(`[]string{"ar", "an", "al"} should be
			%#v
			but got
			%#v`, []string{"ar", "an", "al"}, re.FindAllString("paranormal", -1))
	}

	if !reflect.DeepEqual([]string{"ar", "an"}, re.FindAllString("paranormal", 2)) {
		t.Fatalf(`[]string{"ar", "an"} should be
			%#v
			but got
			%#v`, []string{"ar", "an"}, re.FindAllString("paranormal", 2))
	}
}

func TestFindStringSubmatch(t *testing.T) {
	re := regexp.MustCompile("a(x*)b(y|z)c")

	if !reflect.DeepEqual(re.FindStringSubmatch("-axxxbyc-"), []string{"axxxbyc", "xxx", "y"}) {
		t.Errorf(`re.FindStringSubmatch("-axxxbyc-") = %#v, want %#v`, re.FindStringSubmatch("-axxxbyc-"), []string{"axxxbyc", "xxx", "y"})
	}

	if !reflect.DeepEqual(re.FindStringSubmatch("-abzc"), []string{"abzc", "", "z"}) {
		t.Errorf(`re.FindStringSubmatch("-abzc") = %#v, want %#v`, re.FindStringSubmatch("-abzc"), []string{"abzc", "", "z"})
	}
}
