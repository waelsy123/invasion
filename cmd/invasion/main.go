package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/waelsy123/invasion/earth"
)

func main() {
	exitCode := cmd(os.Args)
	os.Exit(exitCode)
}

func cmd(args []string) int {
	var (
		n        int
		filename string
	)

	app := &cli.App{
		Name:            "invasion",
		Usage:           "alien invasion of your cities",
		HideHelpCommand: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "file",
				Aliases:     []string{"f"},
				Usage:       "`FILE` to load map from",
				Required:    true,
				Destination: &filename,
			},
			&cli.IntFlag{
				Name:        "aliens",
				Aliases:     []string{"n"},
				DefaultText: "10",
				Usage:       "Number of attacking aliens",
				Required:    true,
				Destination: &n,
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Printf("input: %d %s\n", n, filename)

			earth.CreateBoard(filename, n)

			return nil
		},
	}

	err := app.Run(args)
	if err != nil {
		return 1
	}

	return 0
}
