package request

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
)

var (
	ErrNoSpeakEntity        = errors.New("no text or ssml")
	ErrInvalidSpeakingSpeed = errors.New("invalid speaking speed")
)

type request struct {
	Text         string
	SSML         string
	Language     string
	Voice        string
	Speed        float64
	SampleRate   int
	OutputFormat string
	FolderID     string
}

func (r request) Body() (io.Reader, error) {
	v := url.Values{}

	if r.Text != "" {
		v.Set("text", r.Text)
	} else if r.SSML != "" {
		v.Set("ssml", r.SSML)
	} else {
		return nil, ErrNoSpeakEntity
	}

	if r.Language != "" {
		v.Set("lang", r.Language)
	}

	if r.Voice != "" {
		v.Set("voice", r.Voice)
	}

	if r.Speed != 0 {
		v.Set("speed", fmt.Sprintf("%.1f", r.Speed))
	}

	if r.SampleRate != 0 {
		v.Set("sampleRateHertz", strconv.Itoa(r.SampleRate))
	}

	if r.FolderID != "" {
		v.Set("folderId", r.FolderID)
	}

	if r.OutputFormat != "" {
		v.Set("format", r.OutputFormat)
	}

	return strings.NewReader(v.Encode()), nil
}
