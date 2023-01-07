package app

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kmrtftech/tg-framework/pkg/tmplkit"
	"github.com/kmrtftech/tg-framework/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func cmdSetup() *cli.Command {
	return &cli.Command{
		Name:  "setup",
		Usage: "Setup typical-go",
		Flags: []cli.Flag{
			projectPkgFlag,
			typicalBuildFlag,
			typicalTmpFlag,
			&cli.BoolFlag{Name: "go-mod", Usage: "Initiate go.mod before setup"},
			&cli.BoolFlag{Name: "new", Usage: "Setup new project with standard layout and typical-build"},
		},
		Action: func(c *cli.Context) error {
			return Setup(&typgo.Context{Context: c})
		},
	}
}

// Setup typical-go
func Setup(c *typgo.Context) error {
	if c.Bool("go-mod") {
		if err := initGoMod(c); err != nil {
			return err
		}
	}

	p, err := GetParam(c)
	if err != nil {
		return err
	}

	if c.Bool("new") {
		newProject(p)
	}
	return createWrapper(p)
}

// initGoMod initiate gomodob
func initGoMod(c *typgo.Context) error {
	fmt.Fprintf(Stdout, "Initiate go.mod\n")
	pkg := c.String(ProjectPkgParam)
	if pkg == "" {
		return errors.New("project-pkg is empty")
	}
	dir := filepath.Base(pkg)
	os.Mkdir(dir, 0777)
	var stderr strings.Builder
	if err := c.ExecuteCommand(&typgo.Command{
		Name:   "go",
		Args:   []string{"mod", "init", pkg},
		Stderr: &stderr,
		Dir:    dir,
	}); err != nil {
		return fmt.Errorf("%s: %s", err.Error(), stderr.String())
	}
	return nil
}

func createWrapper(p *Param) error {
	path := fmt.Sprintf("%s/typicalw", p.SetupTarget)
	fmt.Fprintf(Stdout, "Create '%s'\n", path)
	return tmplkit.WriteFile(path, typicalwTmpl, p)
}

func newProject(p *Param) {
	mainPkg := p.SetupTarget + "/cmd/" + p.ProjectName
	main := mainPkg + "/main.go"
	fmt.Fprintf(Stdout, "Create '%s'\n", main)
	os.MkdirAll(mainPkg, 0777)
	tmplkit.WriteFile(main, mainTmpl, p)

	appPkg := p.SetupTarget + "/internal/app"
	appStart := appPkg + "/start.go"
	fmt.Fprintf(Stdout, "Create '%s'\n", appStart)
	os.MkdirAll(appPkg, 0777)
	os.WriteFile(appStart, []byte(appStartSrc), 0777)

	generatedPkg := p.SetupTarget + "/internal/generated/ctor"
	generatedDoc := generatedPkg + "/ctor.go"
	fmt.Fprintf(Stdout, "Create '%s'\n", generatedDoc)
	os.MkdirAll(generatedPkg, 0777)
	os.WriteFile(generatedDoc, []byte(generatedDocSrc), 0777)

	typicalBuildPkg := p.SetupTarget + "/tools/typical-build"
	typicalBuild := typicalBuildPkg + "/typical-build.go"
	fmt.Fprintf(Stdout, "Create '%s'\n", typicalBuild)
	os.MkdirAll(typicalBuildPkg, 0777)
	tmplkit.WriteFile(typicalBuild, typicalBuildTmpl, p)

	os.WriteFile(p.SetupTarget+"/.gitignore", []byte(gitignore), 0777)
}

const typicalwTmpl = `#!/bin/bash

set -eu

PROJECT_PKG="{{.ProjectPkg}}"
BUILD_TOOL="{{.TypicalBuild}}"
TYPTMP={{.TypicalTmp}}
TYPGO=$TYPTMP/bin/typical-go
TYPGO_SRC=github.com/kmrtftech/tg-framework

if ! [ -s $TYPGO ]; then
	echo "Build $TYPGO_SRC to $TYPGO"
	go build -o $TYPGO $TYPGO_SRC
fi

$TYPGO run \
	-project-pkg=$PROJECT_PKG \
	-typical-build=$BUILD_TOOL \
	-typical-tmp=$TYPTMP \
	$@
`

const generatedDocSrc = `package ctor
`

const appStartSrc = `package app

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Start app
func Start() {
	// TODO: change start app implementation
	fmt.Println("Hello world!")
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

// Stop app
func Stop() {
	// TODO: change graceful shutdown implementation
	fmt.Printf("Stop app at %s", time.Now())
}
`
const mainTmpl = `package main

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"fmt"
	"log"

	"{{.ProjectPkg}}/internal/app"
	_ "{{.ProjectPkg}}/internal/generated/ctor"
	"github.com/kmrtftech/tg-framework/pkg/typapp"
	"github.com/kmrtftech/tg-framework/pkg/typgo"
)

func main() {
	fmt.Printf("%s %s\n", typgo.ProjectName, typgo.ProjectVersion)
	if err := typapp.StartApp(app.Start, app.Stop); err != nil {
		log.Fatal(err)
	}
}
`

const typicalBuildTmpl = `package main

import (
	"time"
	
	"github.com/kmrtftech/tg-framework/pkg/typgen"
	"github.com/kmrtftech/tg-framework/pkg/typapp"
	"github.com/kmrtftech/tg-framework/pkg/typgo"
	"github.com/kmrtftech/tg-framework/pkg/typmock"
)

var descriptor = typgo.Descriptor{
	ProjectName:    "{{.ProjectName}}",
	ProjectVersion: "0.0.1",

	Tasks: []typgo.Tasker{
		// generate
		&typgen.CodeGenerator{
			Annotators: []typgen.Annotator{
				&typapp.CtorAnnot{},
			},
		},
		// build
		&typgo.GoBuild{},
		// test
		&typgo.GoTest{
			Timeout:  30 * time.Second,
			Includes: []string{"internal/*"},
			Excludes: []string{"internal/generated"},
		},
		// run
		&typgo.RunBinary{Before: typgo.TaskNames{"generate", "build"}},
		// mock
		&typmock.GoMock{},
	},
}

func main() {
	typgo.Start(&descriptor)
}
`

const gitignore = `/bin
/release
/.typical-tmp
/vendor 
.envrc
.env
*.test
*.out`
