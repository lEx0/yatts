// The MIT License (MIT)
//
// Copyright (c) 2024 Amangeldy Kadyl
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
	"github.com/yandex-cloud/go-genproto/yandex/cloud/ai/tts/v3"
	"time"
)

type (
	audioTemplate struct {
		textTemplate     string
		textVariables    map[string]string
		defaultVariables map[string]AudioVariable
		audioSource      []byte
		audioFormat      outputFormat
		audioSampleRate  int
	}
	AudioVariable struct {
		Value  string
		Start  time.Duration
		Length time.Duration
	}
)

func (t audioTemplate) TextTemplate() *tts.UtteranceSynthesisRequest_TextTemplate {
	utterance := &tts.UtteranceSynthesisRequest_TextTemplate{
		TextTemplate: &tts.TextTemplate{
			TextTemplate: t.textTemplate,
		},
	}

	for key, value := range t.textVariables {
		utterance.TextTemplate.Variables = append(
			utterance.TextTemplate.Variables,
			&tts.TextVariable{
				VariableName:  key,
				VariableValue: value,
			},
		)
	}

	return utterance

}

func (t audioTemplate) Hints() ([]*tts.Hints, error) {
	hint := tts.Hints_AudioTemplate{
		AudioTemplate: &tts.AudioTemplate{
			Audio: &tts.AudioContent{
				AudioSource: &tts.AudioContent_Content{
					Content: t.audioSource,
				},
			},
			TextTemplate: &tts.TextTemplate{
				TextTemplate: t.textTemplate,
				Variables:    []*tts.TextVariable{},
			},
			Variables: []*tts.AudioVariable{},
		},
	}

	for key, value := range t.defaultVariables {
		hint.AudioTemplate.TextTemplate.Variables = append(
			hint.AudioTemplate.TextTemplate.Variables,
			&tts.TextVariable{
				VariableName:  key,
				VariableValue: value.Value,
			},
		)
		hint.AudioTemplate.Variables = append(
			hint.AudioTemplate.Variables,
			&tts.AudioVariable{
				VariableName:     key,
				VariableStartMs:  value.Start.Milliseconds(),
				VariableLengthMs: value.Length.Milliseconds(),
			},
		)
	}

	var err error

	hint.AudioTemplate.Audio.AudioSpec, err = buildAudioFormat(
		t.audioFormat,
		t.audioSampleRate,
	)

	return []*tts.Hints{{Hint: &hint}}, err
}
