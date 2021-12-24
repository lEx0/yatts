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
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestNewRequest(t *testing.T) {
	assert.Equal(t, *(NewRequest()), request{})
}

func TestRequest_Body(t *testing.T) {
	t.Run("with all fields", func(t *testing.T) {
		r := request{
			Text:         "text",
			SSML:         "<speak>123</speak>",
			Language:     "ru-RU",
			Voice:        "oksana",
			Speed:        1,
			SampleRate:   8000,
			OutputFormat: "lpcm",
			FolderID:     "123123",
		}

		body, err := r.Body()
		assert.Nil(t, err)

		data, err := ioutil.ReadAll(body)
		assert.Nil(t, err)
		assert.Equal(t,
			`folderId=123123&format=lpcm&lang=ru-RU&sampleRateHertz=8000&speed=1.0&text=text&voice=oksana`,
			string(data),
		)
	})
	t.Run("without any fields", func(t *testing.T) {
		r := request{}

		body, err := r.Body()
		assert.ErrorIs(t, err, ErrNoSpeakEntity)
		assert.Nil(t, body)
	})
	t.Run("with only ssml field", func(t *testing.T) {
		r := request{
			SSML: "<speak>123</speak>",
		}

		body, err := r.Body()
		assert.Nil(t, err)

		data, err := ioutil.ReadAll(body)
		assert.Nil(t, err)
		assert.Equal(t,
			`ssml=%3Cspeak%3E123%3C%2Fspeak%3E`,
			string(data),
		)
	})
}
