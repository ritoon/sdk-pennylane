package main

import (
	"context"
	"fmt"
	"io"
	"strconv"

	pennylane "github.com/ritoon/sdk-pennylane"
)

func creerMandat(ctx context.Context, c *pennylane.Client, id string, in io.Reader, out io.Writer) error {
	customer, err := strconv.ParseInt(id, 10, 64)
	if err != nil || customer <= 0 {
		return fmt.Errorf("identifiant client invalide")
	}
	var body pennylane.PostSepaMandatesJSONRequestBody
	if err := lireJSON(in, &body); err != nil {
		return err
	}
	if body.CustomerId != 0 && body.CustomerId != customer {
		return fmt.Errorf("customer_id diffère du client choisi")
	}
	body.CustomerId = customer
	if body.Identifier == "" || body.Iban == "" || body.Bic == "" || body.SignedAt.IsZero() {
		return fmt.Errorf("renseigner identifier, iban, bic et signed_at")
	}
	r, err := c.PostSepaMandatesWithResponse(ctx, nil, body)
	if err != nil {
		return err
	}
	if r.JSON201 == nil {
		return fmt.Errorf("création du mandat SEPA : HTTP %d", r.StatusCode())
	}
	if r.JSON201.Id <= 0 || r.JSON201.Customer.Id != customer {
		return fmt.Errorf("réponse du mandat sans identifiant ou avec un autre client")
	}
	_, err = fmt.Fprintln(out, r.JSON201.Id)
	return err
}
