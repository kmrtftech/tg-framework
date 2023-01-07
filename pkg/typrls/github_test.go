package typrls_test

import (
	"os"
	"testing"

	"github.com/kmrtftech/tg-framework/pkg/typrls"
	"github.com/stretchr/testify/require"
)

func TestGithub_Publish(t *testing.T) {
	os.Unsetenv("GITHUB_TOKEN")
	github := &typrls.Github{}
	require.EqualError(t, github.Publish(nil), "github-release: missing $GITHUB_TOKEN")
}
