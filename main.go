package main

import (
	"confluence_cli/confluence_commands"
	"confluence_cli/log"
	"github.com/urfave/cli/v2"
	"os"
)

func init() {
	os.Setenv("APP_NAME", "backend_testing_dev")
	log.InitLogger(false)
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")
}

func main() {
	app := &cli.App{
		Name:  "confluence_cli",
		Usage: "Integrate with Confluence Website easily by Command",
		Commands: []*cli.Command{
			confluence_commands.CreatePageCommand(),
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
