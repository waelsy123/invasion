package invasion

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Main(args []string) int {
	var (
		n    int
		file string
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
				Destination: &file,
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
			fmt.Printf("input: %d %s", n, file)

			return nil
		},
	}

	err := app.Run(args)
	if err != nil {
		return 1
	}

	return 0
}
