package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	exitCode := cmd(os.Args)
	os.Exit(exitCode)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func cmd(args []string) int {
	var (
		n        int
		filename string
	)

	app := &cli.App{
		Name:            "generate",
		Usage:           "generate dataset",
		HideHelpCommand: true,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "aliens",
				Aliases:     []string{"n"},
				DefaultText: "10",
				Usage:       "number of connections - each connection 2 cities",
				Required:    true,
				Destination: &n,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Usage:       "`FILE` to write",
				Required:    true,
				Destination: &filename,
			},
		},
		Action: func(c *cli.Context) error {
			log.Printf("n: %v\n", n)

			f, err := os.Create(filename)
			check(err)
			defer f.Close()

			w := bufio.NewWriter(f)

			for i := 0; i < n; i++ {
				A := "A" + fmt.Sprint(i)
				B := "A" + fmt.Sprint(i+1)

				_, err = fmt.Fprintf(w, "%s north=%s\n", A, B)
			}

			check(err)
			w.Flush()

			return nil
		},
	}

	err := app.Run(args)
	if err != nil {
		return 1
	}

	return 0
}
