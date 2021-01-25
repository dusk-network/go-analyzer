package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/dusk-network/go-analyzer/pkg/runner"
	"github.com/urfave/cli/v2"
)

var CLIFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "all",
		Aliases: []string{"a"},
		Usage:   "Run all custom lints.",
		Value:   false,
	},
}

func main() {
	app := &cli.App{
		Name:      "go-analyzer",
		Usage:     "Performs custom lint checks on Golang repositories.",
		Copyright: "Copyright (c) 2021 DUSK",
		Version:   "0.1.0",
		Flags:     CLIFlags,
		Action:    action,
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}

func action(ctx *cli.Context) error {
	if arguments := ctx.Args(); arguments.Len() == 0 {
		return errors.New("no lints selected. exiting...")
	}

	if ctx.Bool("all") {
		if errs := runner.RunAll(); errs != nil {
			for err := range errs {
				fmt.Println(err)
			}

			return errors.New("lint failed")
		}
	}

	return nil
}
