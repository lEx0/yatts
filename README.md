# YaTTS
[![Test](https://github.com/lEx0/yatts/actions/workflows/go.yml/badge.svg)](https://github.com/lEx0/yatts/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/lEx0/yatts?status.svg)](https://godoc.org/github.com/lEx0/yatts)
[![Go Report](https://goreportcard.com/badge/github.com/lEx0/yatts)](https://goreportcard.com/report/github.com/lEx0/yatts)
[![codecov](https://codecov.io/gh/lEx0/yatts/branch/master/graph/badge.svg)](https://codecov.io/gh/lEx0/yatts)

Golang's library for synthesizing speech from text using Yandex.Speech API V1

## Features
 - Multiple authantication methods (iam, api token)
 - Support SSML
 - Return lpcm, Ogg/Opus, mp3 (v3)

## Install
 - speechkit v1 (rest): `go get -u github.com/lEx0/yatts`
 - speechkit v3 (grpc): `go get -u github.com/lEx0/yatts/v3`

## Example

```go
package main

import (
	"github.com/lEx0/yatts"
	"github.com/lEx0/yatts/auth"
	"github.com/lEx0/yatts/request"
	
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := yatts.NewYaTTS(auth.NewAPITokenAuth("YOUR_API_TOKEN"), nil)
	r, err := client.Speak(
		ctx,
		request.SSMLTextEntity{
			SSML: "<speak>Привет мир <break time=\"2s\"/></speak>",
		},

		request.OutputFormat(request.OutputFormatLPCM),
		request.SampleRate(request.OutputSampleRate48k),
		request.Language(request.LangRu),
		request.Voice(request.VoiceOksanaRC),
	)

	if err != nil {
		log.Fatalln(err)
	}

	defer r.Close()

	buff := bytes.Buffer{}
	_, _ = io.Copy(&buff, r)

	fmt.Println(buff.Len())
}

```
