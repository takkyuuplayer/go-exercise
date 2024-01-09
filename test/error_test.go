package test_test

import (
	"crypto/x509"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorCast(t *testing.T) {
	t.Parallel()
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
