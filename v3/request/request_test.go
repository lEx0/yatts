package request

import "testing"

func TestRequest_Build(t *testing.T) {
	r := NewRequest()
	r.Emotion = "good"
	r.Speed = 1.5
	r.Voice = "nigora"
	r.Text = "hello"
	r.OutputFormat = "lpcm"
	r.SampleRate = 48000

	req, err := r.Build()

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if req == nil {
		t.Error("request must be not nil")
		t.FailNow()
	}
}
