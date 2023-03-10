package app

import (
	"os"

	"github.com/kmrtftech/tg-framework/pkg/typgo"
	"github.com/urfave/cli/v2"
)

// Main function the typical-go
func Main() (err error) {
	return App().Run(os.Args)
}

// App application
func App() *cli.App {
	app := cli.NewApp()
	app.Name = typgo.ProjectName
	app.Version = typgo.ProjectVersion
	app.Usage = ""       // NOTE: intentionally blank
	app.Description = "" // NOTE: intentionally blank
	app.Commands = []*cli.Command{
		cmdRun(),
		cmdSetup(),
	}
	return app
}
