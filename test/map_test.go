package test

import (
	"reflect"
	"testing"
)

func TestMapZeroValue(t *testing.T) {
	m := map[string]int{
		"alice": 32,
		"bob":   30,
	}

	if !reflect.DeepEqual(m, map[string]int{"alice": 32, "bob": 30}) {
		t.Errorf(`m = %#v, want %#v`, m, map[string]int{"alice": 32, "bob": 30})
	}

	if m["charlie"] != 0 {
		t.Errorf(`m["charlie"] = %#v, want %#v`, m["charlie"], 0)
	}

	charlie, ok := m["charlie"]

	if charlie != 0 {
		t.Errorf(`charlie = %#v, want %#v`, charlie, 0)
	}

	if ok != false {
		t.Errorf(`ok = %#v, want %#v`, ok, false)
	}

}

func TestMultiMap(t *testing.T) {
	m := map[string]map[string]bool{}

	if m["a"]["b"] != false {
		t.Errorf(`m["a"]["b"] = %#v, want %#v`, m["a"]["b"], false)
	}
}
