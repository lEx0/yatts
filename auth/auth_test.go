// The MIT License (MIT)
//
// Copyright (c) 2021 Amangeldy Kadyl
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
