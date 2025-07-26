package internal

import (
	"errors"
	"flag"
	"os"
)

var verbose = flag.Bool("verbose", false, "Be verbose")

func Read(args []string) (Command, error) {
	fs := flag.NewFlagSet("", flag.ExitOnError)

	fs.Bool("verbose", false, "Be verbose")

	fs.Parse(args)
	found, ok := os.LookupEnv("KASURI_SLACK_TOKEN")
	if !ok {
		return Command{}, errors.New("KASURI_SLACK_TOKEN not found")
	}
	return Command{token: found}, nil
}

type Command struct {
	token   string // Slack Token
	verbose bool   // Be verbose
}
