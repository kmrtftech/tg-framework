package oskit_test

import (
	"os"
	"testing"

	"github.com/kmrtftech/tg-framework/pkg/oskit"
	"github.com/stretchr/testify/require"
)

func TestMkdirAll(t *testing.T) {
	func() {
		defer oskit.MkdirAll("some-dir")()
		_, err := os.Stat("some-dir")
		require.False(t, os.IsNotExist(err))
	}()
	_, err := os.Stat("some-dir")
	require.True(t, os.IsNotExist(err))
}
