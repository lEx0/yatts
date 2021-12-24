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
