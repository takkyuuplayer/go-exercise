package test_test

import (
	"crypto/x509"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestErrorCast(t *testing.T) {
	var err error = &url.Error{
		Err: &x509.UnknownAuthorityError{},
	}

	e, ok := err.(*url.Error)

	assert.Equal(t, e.Err, &x509.UnknownAuthorityError{})
	assert.Equal(t, true, ok)

	e2, ok2 := e.Unwrap().(*x509.UnknownAuthorityError)

	assert.Equal(t, e2, &x509.UnknownAuthorityError{})
	assert.Equal(t, true, ok2)
}
