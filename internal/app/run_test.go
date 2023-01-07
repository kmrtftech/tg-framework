package app_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/kmrtftech/tg-framework/internal/app"
	"github.com/kmrtftech/tg-framework/pkg/oskit"
	"github.com/kmrtftech/tg-framework/pkg/typgo"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	var out strings.Builder
	app.Stdout = &out
	defer func() { app.Stdout = &out }()

	defer oskit.MkdirAll(".typical-tmp")()

	c := &typgo.Context{
		Context: cliContext([]string{
			"-project-pkg=some-pkg",
		}),
	}
	defer c.PatchBash([]*typgo.MockCommand{
		{CommandLine: "go build -ldflags \"-X github.com/kmrtftech/tg-framework/pkg/typgo.ProjectName=some-pkg -X github.com/kmrtftech/tg-framework/pkg/typgo.ProjectPkg=some-pkg -X github.com/kmrtftech/tg-framework/pkg/typgo.TypicalTmp=.typical-tmp\" -o .typical-tmp/bin/typical-build ./tools/typical-build"},
		{CommandLine: ".typical-tmp/bin/typical-build"},
	})(t)

	require.NoError(t, app.Run(c))
	require.Equal(t, "Build tools/typical-build to .typical-tmp/bin/typical-build\n", out.String())
}

func TestRun_GetParamError(t *testing.T) {
	defer oskit.MkdirAll(".typical-tmp")()
	c := &typgo.Context{
		Context: cliContext([]string{}),
	}
	defer c.PatchBash([]*typgo.MockCommand{
		{CommandLine: "go list -m", ReturnError: errors.New("some-error")},
	})(t)

	err := app.Run(c)
	require.EqualError(t, err, "some-error: ")
}
