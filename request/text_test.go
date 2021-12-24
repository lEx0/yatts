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
	"testing"
)

func TestSimpleTextEntity_process(t *testing.T) {
	type (
		in struct {
			text    string
			request request
		}
		testCase struct {
			name     string
			in       in
			out      error
			expected request
		}
	)

	tests := []testCase{
		{
			name: "check empty text",
			in: in{
				text:    "",
				request: request{},
			},
			out:      ErrEmptyTextEntry,
			expected: request{},
		},
		{
			name: "check non empty text",
			in: in{
				text:    "123",
				request: request{},
			},
			out: nil,
			expected: request{
				Text: "123",
			},
		},
		{
			name: "check truncate ssml",
			in: in{
				text: "bazz",
				request: request{
					SSML: "foo",
				},
			},
			out: nil,
			expected: request{
				Text: "bazz",
			},
		},
	}

	for _, entry := range tests {
		t.Run(entry.name, func(t *testing.T) {
			actual := SimpleTextEntity{
				Text: entry.in.text,
			}.Process(&entry.in.request)

			assert.Equal(t, entry.out, actual)
			assert.Equal(t, entry.expected, entry.in.request)
		})
	}
}

func TestSSMLTextEntity_process(t *testing.T) {
	type (
		in struct {
			text    string
			request request
		}
		testCase struct {
			name     string
			in       in
			out      error
			expected request
		}
	)

	tests := []testCase{
		{
			name: "check empty text",
			in: in{
				text:    "",
				request: request{},
			},
			out:      ErrEmptyTextEntry,
			expected: request{},
		},
		{
			name: "check invalid ssml",
			in: in{
				text: "laskdjf",
				request: request{
					Text: "foooo",
				},
			},
			out: ErrInvalidSSML,
			expected: request{
				Text: "foooo",
			},
		},
		{
			name: "check ssml with invalid root tag",
			in: in{
				text: "<foo>GMMN</foo>",
				request: request{
					Text: "foooo",
				},
			},
			out: ErrInvalidSSML,
			expected: request{
				Text: "foooo",
			},
		},
		{
			name: "check valid ssml",
			in: in{
				text:    "<speak>привет</speak>",
				request: request{},
			},
			out: nil,
			expected: request{
				SSML: "<speak>привет</speak>",
			},
		},
		{
			name: "check truncate text",
			in: in{
				text: "<speak>привет</speak>",
				request: request{
					Text: "foo",
				},
			},
			out: nil,
			expected: request{
				SSML: "<speak>привет</speak>",
			},
		},
	}

	for _, entry := range tests {
		t.Run(entry.name, func(t *testing.T) {
			actual := SSMLTextEntity{
				SSML: entry.in.text,
			}.Process(&entry.in.request)

			assert.Equal(t, entry.out, actual)
			assert.Equal(t, entry.expected, entry.in.request)
		})
	}
}
