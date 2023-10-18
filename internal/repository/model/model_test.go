package model

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGet(t *testing.T) {
	t.Run("fail on bad type assertion", func(t *testing.T) {
		data := OrderData{}
		err := data.Get(123)
		require.Error(t, err)
		assert.Equal(t, "type assertion to []byte failed", err.Error())
	})

	t.Run("fail on json unmarshal", func(t *testing.T) {
		data := OrderData{}
		err := data.Get([]byte("not a json"))
		require.Error(t, err)
	})

	t.Run("success on correct input", func(t *testing.T) {
		expected := OrderData{
			OrderUid:    "test",
			TrackNumber: "test",
		}
		bytes, err := json.Marshal(expected)
		require.NoError(t, err)

		data := OrderData{}
		err = data.Get(bytes)
		require.NoError(t, err)

		assert.Equal(t, expected, data)
	})
}
