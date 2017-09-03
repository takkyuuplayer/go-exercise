package minna

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeepEqual(t *testing.T) {
	if !reflect.DeepEqual([]string{"1", "2"}, []string{"1", "2"}) {
		t.Fatal(fmt.Sprintf(`[]string{"1","2"} should be
			%#v
			but got
			%#v`, []string{"1", "2"}, []string{"1", "2"}))
	}
}
