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

package request

import (
	"errors"
)

var (
	ErrEmptyTextEntry = errors.New("empty text entry")
)

type (
	TextEntity interface {
		Process(req *request) error
	}

	SimpleTextEntity struct {
		Text string
	}

	AudioTemplateEntity struct {
		audioTemplate
	}
)

func NewAudioTemplateEntity(
	textTemplate string,
	textVariables map[string]string,
	defaultVariables map[string]AudioVariable,
	audioSource []byte,
	audioFormat outputFormat,
	audioSampleRate int,
) AudioTemplateEntity {
	return AudioTemplateEntity{
		audioTemplate: audioTemplate{
			textTemplate:     textTemplate,
			textVariables:    textVariables,
			defaultVariables: defaultVariables,
			audioSource:      audioSource,
			audioFormat:      audioFormat,
			audioSampleRate:  audioSampleRate,
		},
	}
}

func (e AudioTemplateEntity) Process(req *request) error {
	if len(e.audioSource) == 0 {
		return ErrNoSpeakEntity
	}

	req.AudioTemplate = &e.audioTemplate

	return nil
}

func (e SimpleTextEntity) Process(req *request) error {
	if e.Text == "" {
		return ErrEmptyTextEntry
	}

	req.Text = e.Text

	return nil
}
