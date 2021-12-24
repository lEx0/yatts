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
			name: "set turkish language",
			in: in{
				language: LangTr,
				request: request{
					SampleRate: 8000,
				},
			},
			want: nil,
			result: request{
				SampleRate: 8000,
				Language:   "tr-TR",
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
			name: "set oksana voice",
			in: in{
				voice:   VoiceOksana,
				request: request{},
			},
			want:   nil,
			result: request{Voice: "oksana"},
		},
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
			name: "set oksana:rc voice",
			in: in{
				voice:   VoiceOksanaRC,
				request: request{},
			},
			want:   nil,
			result: request{Voice: "oksana:rc"},
		},
		{
			name: "set jane:rc voice",
			in: in{
				voice:   VoiceJaneRC,
				request: request{},
			},
			want:   nil,
			result: request{Voice: "jane:rc"},
		},
		{
			name: "set omazh:rc voice",
			in: in{
				voice:   VoiceOmazhRC,
				request: request{},
			},
			want:   nil,
			result: request{Voice: "omazh:rc"},
		},
		{
			name: "set zahar:rc voice",
			in: in{
				voice:   VoiceZaharRC,
				request: request{},
			},
			want:   nil,
			result: request{Voice: "zahar:rc"},
		},
		{
			name: "set ermil:rc voice",
			in: in{
				voice:   VoiceErmilRC,
				request: request{},
			},
			want:   nil,
			result: request{Voice: "ermil:rc"},
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
