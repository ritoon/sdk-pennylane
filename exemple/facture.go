package main

import (
	"context"
	"fmt"
	"io"
	"net/mail"
	"strconv"

	pennylane "github.com/ritoon/sdk-pennylane"
)

func envoyerFacture(ctx context.Context, c *pennylane.Client, id, email string, out io.Writer) error {
	invoice, err := strconv.Atoi(id)
	if err != nil || invoice <= 0 {
		return fmt.Errorf("identifiant de facture invalide")
	}
	recipient, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("adresse du destinataire invalide")
	}
	recipients := []string{recipient.Address}
	r, err := c.SendByEmailCustomerInvoiceWithResponse(ctx, invoice, nil, pennylane.SendByEmailCustomerInvoiceJSONRequestBody{Recipients: &recipients})
	if err != nil {
		return err
	}
	if r.StatusCode() != 204 {
		return fmt.Errorf("envoi de la facture : HTTP %d", r.StatusCode())
	}
	_, err = fmt.Fprintln(out, "Demande d’envoi acceptée.")
	return err
}
