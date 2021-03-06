package cli

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/engvik/sbanken-cli/internal/sbanken"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

type sbankenConn interface {
	ConnectClient(context.Context, *cli.Context, string) error
	SetConfig(*sbanken.Config)
	ListAccounts(*cli.Context) error
	ReadAccount(*cli.Context) error
	ListCards(*cli.Context) error
	ListEfakturas(*cli.Context) error
	PayEfaktura(*cli.Context) error
	ListNewEfakturas(*cli.Context) error
	ReadEfaktura(*cli.Context) error
	ListPayments(*cli.Context) error
	ReadPayment(*cli.Context) error
	ListStandingOrders(*cli.Context) error
	ListTransactions(*cli.Context) error
	Transfer(*cli.Context) error
	GetCustomer(*cli.Context) error
}

type tableWriter interface {
	SetStyle(string)
	SetColors(bool)
}

// New creates a new cli app.
func New(ctx context.Context, conn sbankenConn, tw tableWriter, version string) *cli.App {
	flags := getGlobalFlags()

	app := &cli.App{
		Name:    "sbanken",
		Usage:   "provides an easy way to interact with your bank from the terminal",
		Version: version,
		Before: func(c *cli.Context) error {
			configPath := c.String("config")

			if configPath == "" {
				configDir, err := os.UserConfigDir()
				if err != nil {
					return err
				}

				configPath = fmt.Sprintf("%s/sbanken/config.yaml", configDir)
			}

			var hasConfig bool

			loadConfigFunc := altsrc.InitInputSourceWithContext(
				flags,
				func(context *cli.Context) (altsrc.InputSourceContext, error) {
					isc, err := altsrc.NewYamlSourceFromFile(configPath)

					if err == nil && isc.Source() != "" {
						hasConfig = true
					}

					return isc, err
				},
			)

			if err := loadConfigFunc(c); err != nil {
				if c.String("client-id") == "" {
					return errors.New("client-id is a required parameter")
				}

				if c.String("client-secret") == "" {
					return errors.New("client-secret is a required parameter")
				}

				if c.String("customer-id") == "" {
					return errors.New("customer-id is a required parameter")
				}
			}

			if err := conn.ConnectClient(ctx, c, version); err != nil {
				return err
			}

			if hasConfig {
				// Explicitly ignore error, execution should continue without config file
				cfg, _ := sbanken.LoadConfig(configPath)
				conn.SetConfig(cfg)
			}

			tw.SetStyle(c.String("style"))
			tw.SetColors(c.Bool("colors"))

			return nil
		},
		Flags: flags,
		Commands: []*cli.Command{
			getAccountsCommand(conn),
			getCardsCommand(conn),
			getEfakturasCommand(conn),
			getPaymentsCommand(conn),
			getStandingOrdersCommand(conn),
			getTransactionsCommand(conn),
			getTransfersCommand(conn),
			getCustomerCommand(conn),
		},
	}

	app.EnableBashCompletion = true

	return app
}
