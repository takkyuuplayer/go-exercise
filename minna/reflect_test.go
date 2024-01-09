package minna

import (
	"reflect"
	"testing"
)

func TestDeepEqual(t *testing.T) {
	t.Parallel()
	if !reflect.DeepEqual([]string{"1", "2"}, []string{"1", "2"}) {
		t.Fatalf(`[]string{"1","2"} should be
			%#v
			but got
			%#v`, []string{"1", "2"}, []string{"1", "2"})
	}
}
