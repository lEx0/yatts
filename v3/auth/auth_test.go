package auth

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
	"testing"
)

func TestNewAPITokenAuth(t *testing.T) {
	t.Run("with x-folder-id", func(t *testing.T) {
		a := NewAPITokenAuth("token", "123123")

		ctx, err := a.Auth(context.Background())

		if err != nil {
			t.Error("error must be is empty")
			t.FailNow()
		}

		md, exists := metadata.FromOutgoingContext(ctx)

		if !exists {
			t.Error("metadata must exists")
			t.FailNow()
		}

		authValue := md.Get("authorization")

		if len(authValue) != 1 {
			t.Error("authorization value must exists")
			t.FailNow()
		}

		if authValue[0] != "Api-Key token" {
			t.Error("authorization value must be Api-Key token")
			t.FailNow()
		}

		xFolderIDValue := md.Get("x-folder-id")

		if len(xFolderIDValue) == 0 {
			t.Error("x-folder-id value must exists")
			t.FailNow()
		}

		if xFolderIDValue[0] != "123123" {
			t.Error("x-folder-id value must be 123123")
			t.FailNow()
		}
	})
	t.Run("without x-folder-id", func(t *testing.T) {
		a := NewAPITokenAuth("token", "")

		ctx, err := a.Auth(context.Background())

		if err != nil {
			t.Error("error must be is empty")
			t.FailNow()
		}

		md, exists := metadata.FromOutgoingContext(ctx)

		if !exists {
			t.Error("metadata must exists")
			t.FailNow()
		}

		xFolderIDValue := md.Get("x-folder-id")

		if len(xFolderIDValue) != 0 {
			t.Error("x-folder-id value must be empty")
			t.FailNow()
		}
	})
}

func TestNewIAMTokenAuth(t *testing.T) {
	t.Run("with empty GetIamTokenFunc", func(t *testing.T) {
		_, err := NewIAMTokenAuth(nil, "123123")

		if err == nil {
			t.Error("error must be not empty")
			t.FailNow()
		}
	})
	t.Run("GetIamTokenFunc return error", func(t *testing.T) {
		a, err := NewIAMTokenAuth(func() (string, error) {
			return "", errors.New("error")
		}, "123123")

		if err != nil {
			t.Error("error must be is empty")
			t.FailNow()
		}

		_, err = a.Auth(context.Background())

		if err == nil {
			t.Error("error must be not empty")
			t.FailNow()
		}
	})
	t.Run("GetIamTokenFunc return token and with x-folder-id", func(t *testing.T) {
		a, err := NewIAMTokenAuth(func() (string, error) {
			return "iam-token", nil
		}, "123123")

		if err != nil {
			t.Error("error must be is empty")
			t.FailNow()
		}

		ctx, err := a.Auth(context.Background())

		if err != nil {
			t.Error("error must be is empty")
			t.FailNow()
		}

		md, exists := metadata.FromOutgoingContext(ctx)

		if !exists {
			t.Error("metadata must exists")
			t.FailNow()
		}

		authValue := md.Get("authorization")

		if len(authValue) != 1 {
			t.Error("authorization value must exists")
			t.FailNow()
		}

		if authValue[0] != "Bearer iam-token" {
			t.Error("authorization value must be Bearer token")
			t.FailNow()
		}

		xFolderIDValue := md.Get("x-folder-id")

		if len(xFolderIDValue) == 0 {
			t.Error("x-folder-id value must exists")
			t.FailNow()
		}

		if xFolderIDValue[0] != "123123" {
			t.Error("x-folder-id value must be 123123")
			t.FailNow()
		}
	})
}
