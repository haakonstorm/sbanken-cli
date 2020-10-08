package sbanken

import (
	"context"
	"fmt"

	"github.com/engvik/sbanken-go"
	"github.com/urfave/cli/v2"
)

func (c *Connection) Transfer(cliCtx *cli.Context) error {
	ctx := context.Background()

	if err := c.ConnectClient(ctx, cliCtx); err != nil {
		return err
	}

	q := parseTransferQuery(cliCtx)

	if err := c.Client.Transfer(ctx, q); err != nil {
		return err
	}

	fmt.Printf("%d successfully transfered from %s to %s\n", cliCtx.Int("amount"), cliCtx.String("from"), cliCtx.String("to"))

	return nil
}

func parseTransferQuery(ctx *cli.Context) *sbanken.TransferQuery {
	q := &sbanken.TransferQuery{
		FromAccountID: ctx.String("from"),
		ToAccountID:   ctx.String("to"),
		Message:       ctx.String("message"),
		Amount:        float32(ctx.Int("amount")),
	}

	return q
}