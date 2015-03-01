package main

import (
	"flag"
	"io"
)

const (
	ExitCodeOk = iota
	ExitCodeParseFlagError
)

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) Run(args []string) int {

	flags := flag.NewFlagSet("gophue", flag.ContinueOnError)
	flags.SetOutput(c.errStream)

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	return ExitCodeOk
}
