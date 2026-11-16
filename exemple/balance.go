package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	pennylane "github.com/ritoon/sdk-pennylane"
)

func balance(ctx context.Context, c *pennylane.Client, start, end string, out io.Writer) error {
	var p pennylane.GetTrialBalanceParams
	if err := p.PeriodStart.UnmarshalText([]byte(start)); err != nil {
		return err
	}
	if err := p.PeriodEnd.UnmarshalText([]byte(end)); err != nil {
		return err
	}
	if p.PeriodEnd.Before(p.PeriodStart.Time) {
		return fmt.Errorf("la fin de période précède le début")
	}
	encoder := json.NewEncoder(out)
	for {
		r, err := c.GetTrialBalanceWithResponse(ctx, &p)
		if err != nil {
			return err
		}
		if r.JSON200 == nil {
			return fmt.Errorf("balance comptable : HTTP %d", r.StatusCode())
		}
		for _, item := range r.JSON200.Items {
			if err := encoder.Encode(item); err != nil {
				return err
			}
		}
		if !r.JSON200.HasMore {
			return nil
		}
		cursor, err := r.JSON200.NextCursor.Get()
		if err != nil || cursor == "" || (p.Cursor != nil && cursor == *p.Cursor) {
			return fmt.Errorf("curseur de pagination manquant ou répété")
		}
		p.Cursor = &cursor
	}
}
