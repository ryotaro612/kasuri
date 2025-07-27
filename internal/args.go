package internal

import (
	"flag"
)

var verbose = flag.Bool("verbose", false, "Be verbose")

func Read(args []string) (Command, error) {
	fs := flag.NewFlagSet("", flag.ExitOnError)

	verbose := fs.Bool("verbose", false, "Be verbose")

	fs.Parse(args)

	return Command{
		verbose: *verbose,
	}, nil

	// found, ok := os.LookupEnv("KASURI_SLACK_TOKEN")
	// if !ok {
	// 	return Command{}, errors.New("KASURI_SLACK_TOKEN not found")
	// }
	// return Command{token: found}, nil
}

type Command struct {
	token   string // Slack Token
	verbose bool   // Be verbose
}
