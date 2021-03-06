package cli

import (
	"github.com/urfave/cli/v2"
)

type accounts interface {
	ListAccounts(*cli.Context) error
	ReadAccount(*cli.Context) error
}

func getAccountsCommand(conn accounts) *cli.Command {
	return &cli.Command{
		Name:    "accounts",
		Usage:   "list and read accounts",
		Aliases: []string{"a"},
		Subcommands: []*cli.Command{
			{
				Name:    "list",
				Usage:   "list all accounts",
				Aliases: []string{"l"},
				Action:  conn.ListAccounts,
			},
			{
				Name:    "read",
				Usage:   "read a single account",
				Aliases: []string{"r"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "id",
						Usage:    "account id (or name) to read",
						Required: true,
					},
				},
				Action: conn.ReadAccount,
			},
		},
	}
}
