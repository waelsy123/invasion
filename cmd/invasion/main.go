package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/urfave/cli/v2"
	"github.com/waelsy123/invasion/invasion"
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

	const (
		MAX_ITERATION = 10000
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
			rand.Seed(time.Now().UnixNano())

			board := invasion.Board{}
			board.Init(filename, n)

			// log.Printf("board: %+v\n", board)

			board.DestoryPhase()

			for i := 0; i < MAX_ITERATION; i++ {
				log.Printf("i: %v\n", i)
				if len(board.GetConnections()) < 2 || len(board.GetAlienLocations()) < 2 {
					// no more connections or aliens, exit
					break
				}

				board.MovingPhase()
				board.DestoryPhase()
			}

			board.Print()

			return nil
		},
	}

	// figure out what to factor
	// write tests and fixtures
	// write build and ci
	// write readme

	err := app.Run(args)
	if err != nil {
		return 1
	}

	return 0
}
