package request

import "testing"

func TestSimpleTextEntity_Process(t *testing.T) {
	r := NewRequest()
	e := SimpleTextEntity{
		Text: "hello",
	}
	err := e.Process(r)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if r.Text != "hello" {
		t.Error("text must be hello")
		t.FailNow()
	}
}
