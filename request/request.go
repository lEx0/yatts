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

package request

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
)

var (
	ErrNoSpeakEntity        = errors.New("no text or ssml")
	ErrInvalidSpeakingSpeed = errors.New("invalid speaking speed")
)

type request struct {
	Text         string
	SSML         string
	Language     string
	Voice        string
	Speed        float64
	SampleRate   int
	OutputFormat string
	FolderID     string
}

func (r request) Body() (io.Reader, error) {
	v := url.Values{}

	if r.Text != "" {
		v.Set("text", r.Text)
	} else if r.SSML != "" {
		v.Set("ssml", r.SSML)
	} else {
		return nil, ErrNoSpeakEntity
	}

	if r.Language != "" {
		v.Set("lang", r.Language)
	}

	if r.Voice != "" {
		v.Set("voice", r.Voice)
	}

	if r.Speed != 0 {
		v.Set("speed", fmt.Sprintf("%.1f", r.Speed))
	}

	if r.SampleRate != 0 {
		v.Set("sampleRateHertz", strconv.Itoa(r.SampleRate))
	}

	if r.FolderID != "" {
		v.Set("folderId", r.FolderID)
	}

	if r.OutputFormat != "" {
		v.Set("format", r.OutputFormat)
	}

	return strings.NewReader(v.Encode()), nil
}
