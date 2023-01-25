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

package yatts

import (
	"context"
	"crypto/tls"
	"github.com/lEx0/yatts/v3/auth"
	"github.com/lEx0/yatts/v3/request"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/ai/tts/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
)

type (
	// TTS is the interface for text to speech
	TTS interface {
		Speak(ctx context.Context, entity request.TextEntity, options ...request.Option) (io.Reader, error)
	}

	// YaTTS is implementation of TTS based on Yandex TTS
	YaTTS struct {
		auth     auth.Authable
		endpoint string
		options  []request.Option
		client   tts.SynthesizerClient
	}
)

// DefaultYandexTTSEndpoint is the default endpoint for the TTS service
const DefaultYandexTTSEndpoint = "tts.api.cloud.yandex.net:443"

// NewYaTTS creates a new YaTTS instance
func NewYaTTS(
	authenticator auth.Authable,
	options ...request.Option,
) (*YaTTS, error) {
	client := &YaTTS{
		auth:     authenticator,
		endpoint: DefaultYandexTTSEndpoint,
		options:  options,
	}

	conn, err := grpc.Dial(client.endpoint, grpc.WithTransportCredentials(
		//nolint:gosec
		credentials.NewTLS(&tls.Config{}),
	))

	if err != nil {
		return nil, err
	}

	client.client = tts.NewSynthesizerClient(conn)

	return client, nil
}

// SetTTSEndpointURL sets the endpoint for the TTS service.
// for cases when you need to use a hybrid
func (y *YaTTS) SetTTSEndpointURL(url string) {
	y.endpoint = url
}

// Speak sends a request to the TTS endpoint and receives an audio stream.
func (y *YaTTS) Speak(ctx context.Context, entity request.TextEntity, options ...request.Option) (io.Reader, error) {
	req, err := y.buildRequest(entity, options...)

	if err != nil {
		return nil, err
	}

	authCtx, err := y.auth.Auth(ctx)

	if err != nil {
		return nil, err
	}

	cctx, cancel := context.WithCancel(authCtx)
	client, err := y.client.UtteranceSynthesis(cctx, req)

	if err != nil {
		cancel()
		return nil, err
	}

	pr, pw := io.Pipe()

	go func() {
		defer cancel()

		for {
			select {
			case <-cctx.Done():
				_ = pw.Close()

				return
			default:
			}

			if resp, err := client.Recv(); err != nil {
				_ = pw.CloseWithError(err)

				return
			} else if _, err := pw.Write(resp.AudioChunk.Data); err != nil {
				_ = pw.Close()

				return
			}
		}
	}()

	return pr, nil
}

// build tts.UtteranceSynthesisRequest to Yandex TTS API
func (y *YaTTS) buildRequest(entity request.TextEntity, options ...request.Option) (*tts.UtteranceSynthesisRequest, error) {
	r := request.NewRequest()

	for _, option := range append(y.options, options...) {
		if err := option(r); err != nil {
			return nil, err
		}
	}

	if err := entity.Process(r); err != nil {
		return nil, err
	}

	return r.Build()
}
