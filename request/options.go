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

type (
	Option func(req *request) error

	lang             string
	voice            string
	outputFormat     string
	outputSampleRate int
	emotion          string
)

//goland:noinspection GoExportedFuncWithUnexportedType
func NewRequest() *request {
	return &request{}
}

// voice details
// https://cloud.yandex.ru/docs/speechkit/tts/voices
const (
	LangRu lang = "ru-RU"
	LangEn lang = "en-US"
	LangKK lang = "kk-KK"
	LangDE lang = "de-DE"
	LangUZ lang = "uz-UZ"

	VoiceJane    voice = "jane"
	VoiceOmazh   voice = "omazh"
	VoiceZahar   voice = "zahar"
	VoiceErmil   voice = "ermil"
	VoiceAlena   voice = "alena"
	VoiceFilipp  voice = "filipp"
	VoiceAmira   voice = "amira"
	VoiceMadi    voice = "madi"
	VoiceMadiRus voice = "madirus"
	VoiceNigora  voice = "nigora"
	VoiceLea     voice = "lea"
	VoiceJohn    voice = "john"

	EmotionNone    emotion = ""
	EmotionNeutral emotion = "neutral"
	EmotionGood    emotion = "good"
	EmotionEvil    emotion = "evil"

	OutputFormatLPCM    outputFormat = "lpcm"
	OutputFormatOggOpus outputFormat = "oggopus"

	OutputSampleRate8k  outputSampleRate = 8000
	OutputSampleRate16k outputSampleRate = 16000
	OutputSampleRate48k outputSampleRate = 48000
)

func Language(name lang) Option {
	return func(req *request) error {
		req.Language = string(name)
		return nil
	}
}

func Voice(name voice) Option {
	return func(req *request) error {
		req.Voice = string(name)
		return nil
	}
}

func Speed(speed float64) Option {
	return func(req *request) error {
		if speed < 0.1 || speed > 3 {
			return ErrInvalidSpeakingSpeed
		}

		req.Speed = speed

		return nil
	}
}

func Emotion(name emotion) Option {
	return func(req *request) error {
		req.Emotion = string(name)
		return nil
	}
}

func OutputFormat(name outputFormat) Option {
	return func(req *request) error {
		req.OutputFormat = string(name)
		return nil
	}
}

func SampleRate(rate outputSampleRate) Option {
	return func(req *request) error {
		req.SampleRate = int(rate)

		return nil
	}
}

func FolderID(id string) Option {
	return func(req *request) error {
		req.FolderID = id

		return nil
	}
}
