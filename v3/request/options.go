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

type (
	Option func(req *request) error

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
	VoiceJane      voice = "jane"
	VoiceOmazh     voice = "omazh"
	VoiceZahar     voice = "zahar"
	VoiceErmil     voice = "ermil"
	VoiceAlena     voice = "alena"
	VoiceFilipp    voice = "filipp"
	VoiceAmira     voice = "amira"
	VoiceMadi      voice = "madi"
	VoiceMadiRu    voice = "madi_ru"
	VoiceNigora    voice = "nigora"
	VoiceLea       voice = "lea"
	VoiceJohn      voice = "john"
	VoiceSaule     voice = "saule"
	VoiceZhanar    voice = "zhanar"
	VoiceDasha     voice = "dasha"
	VoiceJulia     voice = "julia"
	VoiceLera      voice = "lera"
	VoiceMasha     voice = "masha"
	VoiceMarina    voice = "marina"
	VoiceAlexander voice = "alexander"
	VoiceKirill    voice = "kirill"
	VoiceAnton     voice = "anton"
	VoiceSauleRu   voice = "saule_ru"
	VoiceZamiraRu  voice = "zamira_ru"
	VoiceZhanarRu  voice = "zhanar_ru"
	VoiceYulduzRu  voice = "yulduz_ru"
	VoiceZamira    voice = "zamira"
	VoiceYulduz    voice = "yulduz"

	EmotionNone     emotion = ""
	EmotionNeutral  emotion = "neutral"
	EmotionGood     emotion = "good"
	EmotionEvil     emotion = "evil"
	EmotionWhisper  emotion = "whisper"
	EmotionFriendly emotion = "friendly"
	EmotionStrict   emotion = "strict"

	OutputFormatLPCM    outputFormat = "lpcm"
	OutputFormatOggOpus outputFormat = "oggopus"
	OutputFormatMp3     outputFormat = "mp3"
	OutputFormatWav     outputFormat = "wav"

	OutputSampleRate8k  outputSampleRate = 8000
	OutputSampleRate16k outputSampleRate = 16000
	OutputSampleRate48k outputSampleRate = 48000
)

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
		req.OutputFormat = name
		return nil
	}
}

func SampleRate(rate outputSampleRate) Option {
	return func(req *request) error {
		req.SampleRate = int(rate)

		return nil
	}
}
