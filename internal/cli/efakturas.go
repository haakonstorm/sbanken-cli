package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func getEfakturasCommand(conn sbankenConn) *cli.Command {
	return &cli.Command{
		Name:  "efakturas",
		Usage: "interact with efakturas",
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "list all efakturas",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "new",
						Usage: "list only new efakturas",
					},
					&cli.StringFlag{
						Name:  "start-date",
						Usage: "start date to filter on",
					},
					&cli.StringFlag{
						Name:  "end-date",
						Usage: "end date to filter on",
					},
					&cli.StringFlag{
						Name:  "status",
						Usage: "status to filter on",
					},
					&cli.StringFlag{
						Name:  "index",
						Usage: "index to filter on",
					},
					&cli.StringFlag{
						Name:  "length",
						Usage: "length to filter on",
					},
				},
				Action: func(c *cli.Context) error {
					n := c.Args().Get(0)
					if n == "new" {
						fmt.Println("list new efakturas")
						return nil
					}

					fmt.Println("list all efakturas")
					return nil
				},
			},
			{
				Name:  "pay",
				Usage: "pay efaktura",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "id",
						Usage:    "efaktura id to pay",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "account-id",
						Usage:    "account id to pay from",
						Required: true,
					},
					&cli.BoolFlag{
						Name:     "pay-minimum",
						Usage:    "pay only minimum",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Println("pay efakturas")
					return nil
				},
			},
			{
				Name:  "read",
				Usage: "read a single efaktura",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "id",
						Usage:    "efaktura id to read",
						Required: true,
					},
					&cli.StringFlag{
						Name:  "index",
						Usage: "index to filter on",
					},
					&cli.StringFlag{
						Name:  "length",
						Usage: "length to filter on",
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Println("read efaktura")
					return nil
				},
			},
		},
	}
}