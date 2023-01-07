package typgo_test

import (
	"strings"
	"testing"

	"github.com/kmrtftech/tg-framework/pkg/typgo"
	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	var out strings.Builder
	c := typgo.Logger{
		Stdout:  &out,
		Headers: typgo.LogHeaders("cmd-1", "sub-1"),
	}
	c.Info("some-info")
	c.Infof("some-info %d\n", 9999)
	c.Warn("some-warning")
	c.Warnf("some-warning %d\n", 9999)
	require.Equal(t, "cmd-1:sub-1> some-info\ncmd-1:sub-1> some-info 9999\ncmd-1:sub-1> some-warning\ncmd-1:sub-1> some-warning 9999\n", out.String())
}

func TestLogger_NoPanic(t *testing.T) {
	c := typgo.Logger{}
	c.Info("some-info")
	c.Infof("some-info %d\n", 9999)
	c.Warn("some-warning")
	c.Warnf("some-warning %d\n", 9999)
}
