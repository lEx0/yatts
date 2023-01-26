package request

import "testing"

func TestNewRequest(t *testing.T) {
	r := NewRequest()

	if r == nil {
		t.Error("request must be not nil")
		t.FailNow()
	}
}

func TestVoice(t *testing.T) {
	r := NewRequest()
	option := Voice(VoiceNigora)

	if option == nil {
		t.Error("option must be not nil")
		t.FailNow()
	}

	err := option(r)

	if err != nil {
		t.Error("error must be is empty")
		t.FailNow()
	}

	if r.Voice != "nigora" {
		t.Error("voice must be nigora")
		t.FailNow()
	}
}

func TestSpeed(t *testing.T) {
	r := NewRequest()
	option := Speed(1.5)

	if option == nil {
		t.Error("option must be not nil")
		t.FailNow()
	}

	err := option(r)

	if err != nil {
		t.Error("error must be is empty")
		t.FailNow()
	}

	if r.Speed != 1.5 {
		t.Error("speed must be 1.5")
		t.FailNow()
	}
}

func TestEmotion(t *testing.T) {
	r := NewRequest()
	option := Emotion(EmotionGood)

	if option == nil {
		t.Error("option must be not nil")
		t.FailNow()
	}

	err := option(r)

	if err != nil {
		t.Error("error must be is empty")
		t.FailNow()
	}

	if r.Emotion != "good" {
		t.Error("emotion must be good")
		t.FailNow()
	}
}

func TestOutputFormat(t *testing.T) {
	r := NewRequest()
	option := OutputFormat(OutputFormatMp3)

	if option == nil {
		t.Error("option must be not nil")
		t.FailNow()
	}

	err := option(r)

	if err != nil {
		t.Error("error must be is empty")
		t.FailNow()
	}

	if r.OutputFormat != "mp3" {
		t.Error("output format must be mp3")
		t.FailNow()
	}
}

func TestSampleRate(t *testing.T) {
	r := NewRequest()
	option := SampleRate(OutputSampleRate8k)

	if option == nil {
		t.Error("option must be not nil")
		t.FailNow()
	}

	err := option(r)

	if err != nil {
		t.Error("error must be is empty")
		t.FailNow()
	}

	if r.SampleRate != 8000 {
		t.Error("sample rate must be 8000")
		t.FailNow()
	}
}
