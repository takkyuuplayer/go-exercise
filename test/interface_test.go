package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	Inf interface {
		MethodY() string
		MethodZ() string
	}
	strA struct{}
	strB struct {
		Inf
	}
)

func (a *strA) MethodY() string {
	return "A: MethodY"
}

func (a *strA) MethodZ() string {
	return "A: MethodZ"
}

func (b *strB) MethodZ() string {
	return "B: MethodZ"
}

func TestInterfaceEmbedding(t *testing.T) {
	a := &strA{}

	assert.Equal(t, "A: MethodY", a.MethodY())
	assert.Equal(t, "A: MethodZ", a.MethodZ())

	var b Inf = &strB{a}
	assert.Equal(t, "A: MethodY", b.MethodY())
	assert.Equal(t, "B: MethodZ", b.MethodZ())
}
