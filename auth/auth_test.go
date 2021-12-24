package auth

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestNewAPITokenAuth(t *testing.T) {
	t.Run("with token and folder id", func(t *testing.T) {
		auth := NewAPITokenAuth("test")
		assert.Implements(t, (*Authable)(nil), auth)

		req := http.Request{Header: make(http.Header)}
		err := auth.Do(&req)

		assert.NoError(t, err)
		assert.Equal(t, "Api-Key test", req.Header.Get("Authorization"))
	})
	t.Run("without folder id", func(t *testing.T) {
		auth := NewAPITokenAuth("test")
		req := http.Request{Header: make(http.Header)}
		err := auth.Do(&req)

		assert.NoError(t, err)
		assert.Equal(t, "Api-Key test", req.Header.Get("Authorization"))
	})
}

func TestNewIAMTokenAuth(t *testing.T) {
	t.Run("with token func and folder id", func(t *testing.T) {
		auth, err := NewIAMTokenAuth(func() (string, error) {
			return "test", nil
		})

		assert.Implements(t, (*Authable)(nil), auth)
		assert.NoError(t, err)

		req := http.Request{Header: make(http.Header)}
		err = auth.Do(&req)

		assert.NoError(t, err)
		assert.Equal(t, "Bearer test", req.Header.Get("Authorization"))
	})
	t.Run("with token func and empty folderID", func(t *testing.T) {
		auth, err := NewIAMTokenAuth(func() (string, error) {
			return "test", nil
		})

		assert.Implements(t, (*Authable)(nil), auth)
		assert.NoError(t, err)

		req := http.Request{Header: make(http.Header)}
		err = auth.Do(&req)

		assert.NoError(t, err)
		assert.Equal(t, "Bearer test", req.Header.Get("Authorization"))
	})
	t.Run("returned error", func(t *testing.T) {
		auth, err := NewIAMTokenAuth(func() (string, error) {
			return "", io.EOF
		})

		assert.NoError(t, err)
		assert.Implements(t, (*Authable)(nil), auth)

		req := http.Request{Header: make(http.Header)}
		err = auth.Do(&req)

		assert.Error(t, err)
	})
	t.Run("with nil func", func(t *testing.T) {
		auth, err := NewIAMTokenAuth(nil)

		assert.Error(t, err)
		assert.Nil(t, auth)
	})
}
