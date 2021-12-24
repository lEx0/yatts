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
	"errors"
	"net/http"
)

type (
	Authable interface {
		Do(req *http.Request) error
	}

	IAMTokenAuth struct {
		getIAMTokenFunc func() (string, error)
	}

	APITokenAuth struct {
		token string
	}
)

func NewIAMTokenAuth(getIAMTokenFunc func() (string, error)) (*IAMTokenAuth, error) {
	if getIAMTokenFunc == nil {
		return nil, errors.New("invalid getIAMTokenFunc")
	}

	return &IAMTokenAuth{
		getIAMTokenFunc: getIAMTokenFunc,
	}, nil
}

func NewAPITokenAuth(token string) *APITokenAuth {
	return &APITokenAuth{
		token: token,
	}
}

func (a *APITokenAuth) Do(req *http.Request) error {
	req.Header.Set("Authorization", "Api-Key "+a.token)

	return nil
}

func (a *IAMTokenAuth) Do(req *http.Request) error {
	token, err := a.getIAMTokenFunc()

	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	return nil
}
