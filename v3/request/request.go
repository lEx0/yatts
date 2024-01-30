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
	AudioTemplate *audioTemplate
	Text          string
	Voice         string
	Speed         float64
	Emotion       string
	SampleRate    int
	OutputFormat  outputFormat
}

func (r request) Build() (*tts.UtteranceSynthesisRequest, error) {
	result, err := r.buildRequest()

	if err != nil {
		return nil, err
	}

	result.OutputAudioSpec, err = buildAudioFormat(r.OutputFormat, r.SampleRate)

	return result, err
}

func (r request) buildRequest() (*tts.UtteranceSynthesisRequest, error) {
	var (
		result = tts.UtteranceSynthesisRequest{
			Utterance:       nil,
			Hints:           nil,
			OutputAudioSpec: nil,
			UnsafeMode:      false,
		}
		hints = make([]*tts.Hints, 0)
		err   error
	)

	if r.Text != "" {
		result.Utterance = &tts.UtteranceSynthesisRequest_Text{
			Text: r.Text,
		}

		if r.Voice == "" {
			return nil, ErrVoiceNotSpecified
		}

		hints = append(
			hints, &tts.Hints{
				Hint: &tts.Hints_Voice{Voice: r.Voice},
			},
		)

		if r.Speed != 0 {
			if r.Speed < 0.1 || r.Speed > 3 {
				return nil, ErrInvalidSpeakingSpeed
			}

			hints = append(
				hints, &tts.Hints{
					Hint: &tts.Hints_Speed{Speed: r.Speed},
				},
			)
		}

		hints = append(
			hints, &tts.Hints{
				Hint: &tts.Hints_Role{Role: r.Emotion},
			},
		)

		if len(hints) > 0 {
			result.Hints = hints
		}
	} else if r.AudioTemplate != nil {
		result.Utterance = r.AudioTemplate.TextTemplate()
		result.Model = "zsl"
		result.Hints, err = r.AudioTemplate.Hints()
	} else {
		return nil, ErrNoSpeakEntity
	}

	return &result, err
}

func buildAudioFormat(format outputFormat, sampleRate int) (*tts.AudioFormatOptions, error) {
	switch format {
	case OutputFormatLPCM:
		return &tts.AudioFormatOptions{
			AudioFormat: &tts.AudioFormatOptions_RawAudio{
				RawAudio: &tts.RawAudio{
					AudioEncoding:   tts.RawAudio_LINEAR16_PCM,
					SampleRateHertz: int64(sampleRate),
				},
			},
		}, nil
	case OutputFormatWav:
		return &tts.AudioFormatOptions{
			AudioFormat: &tts.AudioFormatOptions_ContainerAudio{
				ContainerAudio: &tts.ContainerAudio{
					ContainerAudioType: tts.ContainerAudio_WAV,
				},
			},
		}, nil
	case OutputFormatOggOpus:
		return &tts.AudioFormatOptions{
			AudioFormat: &tts.AudioFormatOptions_ContainerAudio{
				ContainerAudio: &tts.ContainerAudio{
					ContainerAudioType: tts.ContainerAudio_OGG_OPUS,
				},
			},
		}, nil
	case OutputFormatMp3:
		return &tts.AudioFormatOptions{
			AudioFormat: &tts.AudioFormatOptions_ContainerAudio{
				ContainerAudio: &tts.ContainerAudio{
					ContainerAudioType: tts.ContainerAudio_MP3,
				},
			},
		}, nil
	case "":
		return nil, ErrOutputFormatNotSpecified
	default:
		return nil, ErrInvalidOutputFormat
	}
}
