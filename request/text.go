package request

import (
	"encoding/xml"
	"errors"
)

var (
	ErrEmptyTextEntry = errors.New("empty text entry")
	ErrInvalidSSML    = errors.New("invalid SSML")
)

type (
	TextEntity interface {
		Process(req *request) error
	}

	SimpleTextEntity struct {
		Text string
	}
	SSMLTextEntity struct {
		SSML string
	}

	ssmlValidationStruct struct {
		XMLName xml.Name `xml:"speak"`
	}
)

func (e SimpleTextEntity) Process(req *request) error {
	if e.Text == "" {
		return ErrEmptyTextEntry
	}

	req.SSML = ""
	req.Text = e.Text

	return nil
}

func (e SSMLTextEntity) Process(req *request) error {
	var result ssmlValidationStruct

	if e.SSML == "" {
		return ErrEmptyTextEntry
	} else if err := xml.Unmarshal([]byte(e.SSML), &result); err != nil {
		return ErrInvalidSSML
	}

	req.Text = ""
	req.SSML = e.SSML

	return nil
}
