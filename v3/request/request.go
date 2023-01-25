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
	"github.com/yandex-cloud/go-genproto/yandex/cloud/ai/tts/v3"
)

var (
	ErrNoSpeakEntity            = errors.New("no text or ssml")
	ErrVoiceNotSpecified        = errors.New("voice not specified")
	ErrInvalidSpeakingSpeed     = errors.New("invalid speaking speed")
	ErrOutputFormatNotSpecified = errors.New("output format not specified")
	ErrInvalidOutputFormat      = errors.New("invalid output format")
)

type request struct {
	Text         string
	Voice        string
	Speed        float64
	Emotion      string
	SampleRate   int
	OutputFormat string
}

func (r request) Build() (*tts.UtteranceSynthesisRequest, error) {
	result := tts.UtteranceSynthesisRequest{
		Utterance:       nil,
		Hints:           nil,
		OutputAudioSpec: nil,
		UnsafeMode:      false,
	}

	hints := make([]*tts.Hints, 0)

	if r.Text != "" {
		result.Utterance = &tts.UtteranceSynthesisRequest_Text{
			Text: r.Text,
		}
	} else {
		return nil, ErrNoSpeakEntity
	}

	if r.Voice == "" {
		return nil, ErrVoiceNotSpecified
	}

	hints = append(hints, &tts.Hints{
		Hint: &tts.Hints_Voice{Voice: r.Voice},
	})

	if r.Speed != 0 {
		if r.Speed < 0.1 || r.Speed > 3 {
			return nil, ErrInvalidSpeakingSpeed
		}

		hints = append(hints, &tts.Hints{
			Hint: &tts.Hints_Speed{Speed: r.Speed},
		})
	}

	if r.Emotion != "" {
		hints = append(hints, &tts.Hints{
			Hint: &tts.Hints_Role{Role: r.Emotion},
		})
	}

	if len(hints) > 0 {
		result.Hints = hints
	}

	if r.OutputFormat == "" {
		return nil, ErrOutputFormatNotSpecified
	}

	switch outputFormat(r.OutputFormat) {
	case OutputFormatLPCM:
		result.OutputAudioSpec = &tts.AudioFormatOptions{
			AudioFormat: &tts.AudioFormatOptions_RawAudio{
				RawAudio: &tts.RawAudio{
					AudioEncoding:   tts.RawAudio_LINEAR16_PCM,
					SampleRateHertz: int64(r.SampleRate),
				},
			},
		}
	case OutputFormatOggOpus:
		result.OutputAudioSpec = &tts.AudioFormatOptions{
			AudioFormat: &tts.AudioFormatOptions_ContainerAudio{
				ContainerAudio: &tts.ContainerAudio{
					ContainerAudioType: tts.ContainerAudio_OGG_OPUS,
				},
			},
		}
	case OutputFormatMp3:
		result.OutputAudioSpec = &tts.AudioFormatOptions{
			AudioFormat: &tts.AudioFormatOptions_ContainerAudio{
				ContainerAudio: &tts.ContainerAudio{
					ContainerAudioType: tts.ContainerAudio_MP3,
				},
			},
		}
	default:
		return nil, ErrInvalidOutputFormat
	}

	return &result, nil
}
