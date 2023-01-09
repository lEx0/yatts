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

func TestLanguage(t *testing.T) {
	type (
		in struct {
			language lang
			request  request
		}
		testCase struct {
			name   string
			in     in
			want   error
			result request
		}
	)

	tests := []testCase{
		{
			name: "set russian language",
			in: in{
				language: LangRu,
				request: request{
					SampleRate: 8000,
				},
			},
			want: nil,
			result: request{
				SampleRate: 8000,
				Language:   "ru-RU",
			},
		},
		{
			name: "set english language",
			in: in{
				language: LangEn,
				request: request{
					SampleRate: 8000,
				},
			},
			want: nil,
			result: request{
				SampleRate: 8000,
				Language:   "en-US",
			},
		},
		{
			name: "set kazakh language",
			in: in{
				language: LangKK,
				request: request{
					SampleRate: 8000,
				},
			},
			want: nil,
			result: request{
				SampleRate: 8000,
				Language:   "kk-KK",
			},
		},
		{
			name: "set deutsche language",
			in: in{
				language: LangDE,
				request: request{
					SampleRate: 8000,
				},
			},
			want: nil,
			result: request{
				SampleRate: 8000,
				Language:   "de-DE",
			},
		},
		{
			name: "set deutsche language",
			in: in{
				language: LangUZ,
				request: request{
					SampleRate: 8000,
				},
			},
			want: nil,
			result: request{
				SampleRate: 8000,
				Language:   "uz-UZ",
			},
		},
	}

	for _, entry := range tests {
		t.Run(entry.name, func(t *testing.T) {
			got := Language(entry.in.language)(&entry.in.request)

			assert.Equal(t, entry.want, got)
			assert.Equal(t, entry.result, entry.in.request)
		})
	}
}

func TestSampleRate(t *testing.T) {
	type (
		in struct {
			sr      outputSampleRate
			request request
		}
		testCase struct {
			name   string
			in     in
			want   error
			result request
		}
	)

	tests := []testCase{
		{
			name: "set 8kHz sample rate",
			in: in{
				sr:      OutputSampleRate8k,
				request: request{},
			},
			want:   nil,
			result: request{SampleRate: 8000},
		},
		{
			name: "set 16kHz sample rate",
			in: in{
				sr:      OutputSampleRate16k,
				request: request{},
			},
			want:   nil,
			result: request{SampleRate: 16000},
		},
		{
			name: "set 48kHz sample rate",
			in: in{
				sr:      OutputSampleRate48k,
				request: request{},
			},
			want:   nil,
			result: request{SampleRate: 48000},
		},
	}

	for _, entry := range tests {
		t.Run(entry.name, func(t *testing.T) {
			got := SampleRate(entry.in.sr)(&entry.in.request)

			assert.Equal(t, entry.want, got)
			assert.Equal(t, entry.result, entry.in.request)
		})
	}
}

func TestVoice(t *testing.T) {
	type (
		in struct {
			voice   voice
			request request
		}
		testCase struct {
			name   string
			in     in
			want   error
			result request
		}
	)

	tests := []testCase{
		{
			name: "set jana voice",
			in: in{
				voice:   VoiceJane,
				request: request{},
			},
			want:   nil,
			result: request{Voice: "jane"},
		},
		{
			name: "set omazh voice",
			in: in{
				voice:   VoiceOmazh,
				request: request{},
			},
			want:   nil,
			result: request{Voice: "omazh"},
		},
		{
			name: "set zahar voice",
			in: in{
				voice:   VoiceZahar,
				request: request{},
			},
			want:   nil,
			result: request{Voice: "zahar"},
		},
		{
			name: "set ermil voice",
			in: in{
				voice:   VoiceErmil,
				request: request{},
			},
			want:   nil,
			result: request{Voice: "ermil"},
		},
		{
			name: "set alena voice",
			in: in{
				voice:   VoiceAlena,
				request: request{},
			},
			want:   nil,
			result: request{Voice: "alena"},
		},
		{
			name: "set filipp voice",
			in: in{
				voice:   VoiceFilipp,
				request: request{},
			},
			want:   nil,
			result: request{Voice: "filipp"},
		},
		{
			name: "set Amira voice",
			in: in{
				voice:   VoiceAmira,
				request: request{},
			},
			want:   nil,
			result: request{Voice: "amira"},
		},
		{
			name: "set Madi voice",
			in: in{
				voice:   VoiceMadi,
				request: request{},
			},
			want:   nil,
			result: request{Voice: "madi"},
		},
		{
			name: "set MadiRus voice",
			in: in{
				voice:   VoiceMadiRus,
				request: request{},
			},
			want:   nil,
			result: request{Voice: "madirus"},
		},
		{
			name: "set Nigora voice",
			in: in{
				voice:   VoiceNigora,
				request: request{},
			},
			want:   nil,
			result: request{Voice: "nigora"},
		},
		{
			name: "set Lea voice",
			in: in{
				voice:   VoiceLea,
				request: request{},
			},
			want:   nil,
			result: request{Voice: "lea"},
		},
		{
			name: "set John voice",
			in: in{
				voice:   VoiceJohn,
				request: request{},
			},
			want:   nil,
			result: request{Voice: "john"},
		},
	}

	for _, entry := range tests {
		t.Run(entry.name, func(t *testing.T) {
			got := Voice(entry.in.voice)(&entry.in.request)

			assert.Equal(t, entry.want, got)
			assert.Equal(t, entry.result, entry.in.request)
		})
	}
}

func TestEmotion(t *testing.T) {
	type testCase struct {
		name string
		in   emotion
		want string
	}

	tests := []testCase{
		{
			name: "check neutral emotion",
			in:   EmotionNeutral,
			want: "neutral",
		},
		{
			name: "check good emotion",
			in:   EmotionGood,
			want: "good",
		},
		{
			name: "check evil emotion",
			in:   EmotionEvil,
			want: "evil",
		},
	}

	for _, emotion := range tests {
		t.Run(emotion.name, func(t *testing.T) {
			req := request{}
			_ = Emotion(emotion.in)(&req)
			assert.Equal(t, emotion.want, req.Emotion)
		})
	}
}

func TestSpeed(t *testing.T) {
	type (
		in struct {
			speed   float64
			request request
		}
		testCase struct {
			name   string
			in     in
			want   error
			result request
		}
	)

	tests := []testCase{
		{
			name: "check with speed < 0.1",
			in: in{
				speed: 0.01,
				request: request{
					Text: "123",
				},
			},
			want: ErrInvalidSpeakingSpeed,
			result: request{
				Text: "123",
			},
		},
		{
			name: "check with speed > 3",
			in: in{
				speed: 3.1,
				request: request{
					Text: "123",
				},
			},
			want: ErrInvalidSpeakingSpeed,
			result: request{
				Text: "123",
			},
		},
		{
			name: "check with normal speed",
			in: in{
				speed: 1,
				request: request{
					Text: "123",
				},
			},
			result: request{
				Text:  "123",
				Speed: 1,
			},
		},
	}

	for _, entry := range tests {
		t.Run(entry.name, func(t *testing.T) {
			got := Speed(entry.in.speed)(&entry.in.request)

			assert.Equal(t, entry.want, got)
			assert.Equal(t, entry.result, entry.in.request)
		})
	}
}

func TestFolderID(t *testing.T) {
	t.Run("test with empty folderID", func(t *testing.T) {
		req := request{}
		_ = FolderID("")(&req)

		assert.Equal(t, req.FolderID, "")
	})
	t.Run("test with non empty folderID", func(t *testing.T) {
		req := request{}
		_ = FolderID("123123")(&req)

		assert.Equal(t, req.FolderID, "123123")
	})
}

func TestOutputFormat(t *testing.T) {
	t.Run("set LPCM", func(t *testing.T) {
		req := request{}
		_ = OutputFormat(OutputFormatLPCM)(&req)

		assert.Equal(t, req.OutputFormat, "lpcm")
	})
	t.Run("set oggopus", func(t *testing.T) {
		req := request{}
		_ = OutputFormat(OutputFormatOggOpus)(&req)

		assert.Equal(t, req.OutputFormat, "oggopus")
	})
}
