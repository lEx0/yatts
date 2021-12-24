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

package yatts

import (
	"context"
	"fmt"
	"github.com/lEx0/yatts/auth"
	"github.com/lEx0/yatts/request"
	"io"
	"net/http"
)

type (
	// TTS is the interface for text to speech
	TTS interface {
		Speak(ctx context.Context, entity request.TextEntity, options ...request.Option) (io.ReadCloser, error)
	}

	// YaTTS is implementation of TTS based on Yandex TTS
	YaTTS struct {
		auth   auth.Authable
		client *http.Client
		url    string
	}
)

// DefaultYandexTTSEndpointURL is the default endpoint for the TTS service
const DefaultYandexTTSEndpointURL = "https://tts.api.cloud.yandex.net/speech/v1/tts:synthesize"

// NewYaTTS creates a new YaTTS instance
func NewYaTTS(authenticator auth.Authable, client *http.Client) *YaTTS {
	if client == nil {
		client = http.DefaultClient
	}

	return &YaTTS{
		auth:   authenticator,
		client: client,
		url:    DefaultYandexTTSEndpointURL,
	}
}

// SetTTSEndpointURL sets the endpoint url for the TTS service.
func (y *YaTTS) SetTTSEndpointURL(url string) {
	y.url = url
}

// Speak sends a request to the TTS endpoint and receives an audio stream.
func (y *YaTTS) Speak(ctx context.Context, entity request.TextEntity, options ...request.Option) (io.ReadCloser, error) {
	if req, err := y.buildRequest(ctx, entity, options...); err != nil {
		return nil, err
	} else if resp, err := y.client.Do(req); err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	} else {
		return resp.Body, nil
	}
}

// build http.Request to Yandex TTS API
func (y *YaTTS) buildRequest(ctx context.Context, entity request.TextEntity, options ...request.Option) (*http.Request, error) {
	r := request.NewRequest()

	for _, option := range options {
		if err := option(r); err != nil {
			return nil, err
		}
	}

	if err := entity.Process(r); err != nil {
		return nil, err
	}

	if body, err := r.Body(); err != nil {
		return nil, err
	} else if req, err := http.NewRequestWithContext(
		ctx, http.MethodPost, y.url, body,
	); err != nil {
		return nil, err
	} else if err := y.auth.Do(req); err != nil {
		return nil, err
	} else {
		return req, nil
	}
}
