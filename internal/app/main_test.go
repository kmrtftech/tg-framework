package app_test

import (
	"testing"

	"github.com/kmrtftech/tg-framework/internal/app"
	"github.com/kmrtftech/tg-framework/pkg/typgo"
	"github.com/stretchr/testify/require"
)

func TestApp(t *testing.T) {
	typgo.ProjectName = "some-name"
	typgo.ProjectVersion = "some-version"
	defer func() {
		typgo.ProjectName = ""
		typgo.ProjectVersion = ""
	}()
	app := app.App()
	require.Equal(t, "some-name", app.Name)
	require.Equal(t, "some-version", app.Version)
	require.Equal(t, "run", app.Commands[0].Name)
	require.Equal(t, "setup", app.Commands[1].Name)
}
