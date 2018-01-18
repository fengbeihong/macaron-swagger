package main

import (
	"os"
	"github.com/fengbeihong/macaron-swagger/gen"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "v1.0.0"
	app.Usage = "Automatically generate RESTful API documentation with Swagger 2.0 for Go."

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "lang, l",
			Value: "spanish",
			Usage: "language for the greeting",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Usage:   "create docs.go",
			Action: func(c *cli.Context) error {
				searchDir := "."
				mainApiFile := "./main.go"
				if c.NArg() > 0 {
					mainApiFile = c.Args().Get(0)
				}
				gen.New().Build(searchDir, mainApiFile)

				return nil
			},
		},
	}

	app.Run(os.Args)
}
