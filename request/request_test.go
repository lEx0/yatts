package request

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestNewRequest(t *testing.T) {
	assert.Equal(t, *(NewRequest()), request{})
}

func TestRequest_Body(t *testing.T) {
	t.Run("with all fields", func(t *testing.T) {
		r := request{
			Text:         "text",
			SSML:         "<speak>123</speak>",
			Language:     "ru-RU",
			Voice:        "oksana",
			Speed:        1,
			SampleRate:   8000,
			OutputFormat: "lpcm",
			FolderID:     "123123",
		}

		body, err := r.Body()
		assert.Nil(t, err)

		data, err := ioutil.ReadAll(body)
		assert.Nil(t, err)
		assert.Equal(t,
			`folderId=123123&format=lpcm&lang=ru-RU&sampleRateHertz=8000&speed=1.0&text=text&voice=oksana`,
			string(data),
		)
	})
	t.Run("without any fields", func(t *testing.T) {
		r := request{}

		body, err := r.Body()
		assert.ErrorIs(t, err, ErrNoSpeakEntity)
		assert.Nil(t, body)
	})
	t.Run("with only ssml field", func(t *testing.T) {
		r := request{
			SSML: "<speak>123</speak>",
		}

		body, err := r.Body()
		assert.Nil(t, err)

		data, err := ioutil.ReadAll(body)
		assert.Nil(t, err)
		assert.Equal(t,
			`ssml=%3Cspeak%3E123%3C%2Fspeak%3E`,
			string(data),
		)
	})
}
