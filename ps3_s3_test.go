package ps3

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test(t *testing.T) {
	var ps3 Ps3
	ps3, err := Ps3_s3_new("us-west-1", "ps3-dev1")
	require.Nil(t, err)

	data := []byte(uuid.New().String())

	err = ps3.Put("test", data)
	require.Nil(t, err)

	bytes, err := ps3.Get("test")
	require.Nil(t, err)
	require.Equal(t, string(data), string(bytes))
}
