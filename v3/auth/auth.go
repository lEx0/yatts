// The MIT License (MIT)
//
// Copyright (c) 2023 Amangeldy Kadyl
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
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
)

type (
	Authable interface {
		Auth(ctx context.Context) (context.Context, error)
	}

	getIamTokenFunc func() (string, error)

	IAMTokenAuth struct {
		getIAMTokenFunc getIamTokenFunc
		xFolderID       string
	}

	APITokenAuth struct {
		token     string
		xFolderID string
	}
)

func NewIAMTokenAuth(getIAMTokenFunc getIamTokenFunc, xFolderID string) (*IAMTokenAuth, error) {
	if getIAMTokenFunc == nil {
		return nil, errors.New("invalid getIAMTokenFunc")
	}

	return &IAMTokenAuth{
		getIAMTokenFunc: getIAMTokenFunc,
		xFolderID:       xFolderID,
	}, nil
}

func NewAPITokenAuth(token, xFolderID string) *APITokenAuth {
	return &APITokenAuth{
		token:     token,
		xFolderID: xFolderID,
	}
}

func (a *IAMTokenAuth) Auth(ctx context.Context) (context.Context, error) {
	token, err := a.getIAMTokenFunc()

	if err != nil {
		return nil, err
	}

	kv := []string{
		"authorization",
		"Bearer " + token,
	}

	if a.xFolderID != "" {
		kv = append(kv, "x-folder-id", a.xFolderID)
	}

	return metadata.AppendToOutgoingContext(ctx, kv...), nil
}

func (a *APITokenAuth) Auth(ctx context.Context) (context.Context, error) {
	kv := []string{
		"authorization",
		"Api-Key " + a.token,
	}

	if a.xFolderID != "" {
		kv = append(kv, "x-folder-id", a.xFolderID)
	}

	return metadata.AppendToOutgoingContext(ctx, kv...), nil
}
