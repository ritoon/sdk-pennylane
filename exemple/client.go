package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	pennylane "github.com/ritoon/sdk-pennylane"
)

func lireJSON(in io.Reader, body any) error {
	dec := json.NewDecoder(in)
	dec.DisallowUnknownFields()
	if err := dec.Decode(body); err != nil {
		return err
	}
	if err := dec.Decode(new(json.RawMessage)); err != io.EOF {
		return fmt.Errorf("un seul document JSON est attendu")
	}
	return nil
}

func creerClient(ctx context.Context, c *pennylane.Client, in io.Reader, out io.Writer) error {
	var body pennylane.PostCustomerJSONBody0
	if err := lireJSON(in, &body); err != nil {
		return err
	}
	address := body.BillingAddress
	if body.CustomerType != "company" || body.Name == "" || address.Address == "" || address.City == "" || address.PostalCode == "" || address.CountryAlpha2 == "" {
		return fmt.Errorf("renseigner customer_type=company, name et billing_address complète")
	}
	// Le corps de cette route est une union. Encoder la variante entreprise conserve ses champs.
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}
	r, err := c.PostCustomerWithBodyWithResponse(ctx, nil, "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}
	if r.JSON201 == nil {
		return fmt.Errorf("création du client : HTTP %d", r.StatusCode())
	}
	customer, err := r.JSON201.AsGetCustomerStatus2000()
	if err != nil {
		return err
	}
	if customer.Id <= 0 {
		return fmt.Errorf("identifiant du client absent de la réponse")
	}
	_, err = fmt.Fprintln(out, customer.Id)
	return err
}
