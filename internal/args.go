package internal

import (
	"errors"
	"flag"
	"os"
)

// Read reads command line arguments and environment variables to constrcut Args.
func Read(args []string) (Args, error) {
	fs := flag.NewFlagSet("", flag.ExitOnError)

	verbose := fs.Bool("verbose", false, "Be verbose")
	token := fs.String("token", "", "Slack bot token")

	fs.Parse(args)

	tkn, ok := os.LookupEnv("ASHITABA_SLACK_TOKEN")

	res := Args{
		verbose: *verbose,
	}
	if ok {
		res.token = tkn
	} else {
		res.token = *token
	}

	if res.token == "" {
		return res, errors.New("Slack token is required, set it via -token flag or ASHITABA_SLACK_TOKEN environment variable")
	}

	return res, nil

}

type Args struct {
	token   string // Slack Token
	verbose bool   // Be verbose
}
