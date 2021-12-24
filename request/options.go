package request

type (
	Option func(req *request) error

	lang             string
	voice            string
	outputFormat     string
	outputSampleRate int
)

func NewRequest() *request {
	return &request{}
}

const (
	LangRu lang = "ru-RU"
	LangEn lang = "en-US"
	LangTr lang = "tr-TR"

	VoiceOksana   voice = "oksana"
	VoiceJane     voice = "jane"
	VoiceOmazh    voice = "omazh"
	VoiceZahar    voice = "zahar"
	VoiceErmil    voice = "ermil"
	VoiceAlena    voice = "alena"
	VoiceFilipp   voice = "filipp"
	VoiceOksanaRC voice = "oksana:rc"
	VoiceJaneRC   voice = "jane:rc"
	VoiceOmazhRC  voice = "omazh:rc"
	VoiceZaharRC  voice = "zahar:rc"
	VoiceErmilRC  voice = "ermil:rc"

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
