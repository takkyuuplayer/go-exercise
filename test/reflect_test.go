package test_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyStruct struct {
	A int
}

func TestType(t *testing.T) {
	t.Parallel()
	var a *MyStruct

	f := func(models ...interface{}) {
		for _, model := range models {
			assert.Equal(t, 1, reflect.TypeOf(model).Elem().Elem().NumField())
		}
	}

	f(&a)
	a = &MyStruct{}
	f(&a)
}

func TestSetting(t *testing.T) {
	t.Parallel()
	a := &MyStruct{}

	wipePassed(&a)
	assert.Nil(t, a)

	a = &MyStruct{}
	setValue(&a)
	assert.Equal(t, 1, a.A)

	a = &MyStruct{}
	setValueWithInterface(&a)
	assert.Equal(t, 5, a.A)
}

func wipePassed(r interface{}) {
	v := reflect.ValueOf(r)
	p := v.Elem()
	p.Set(reflect.Zero(p.Type()))
}

func setValue(r interface{}) {
	v := reflect.ValueOf(r)
	p := v.Elem().Elem()

	for i := 0; i < p.NumField(); i++ {
		f := p.Field(i)
		f.SetInt(1)
	}
}

func setValueWithInterface(r interface{}) {
	v := reflect.ValueOf(r)
	p := v.Elem().Elem()

	a := interface{}(int(5))
	for i := 0; i < p.NumField(); i++ {
		f := p.Field(i)
		f.Set(reflect.ValueOf(&a).Elem().Elem())
	}
}
