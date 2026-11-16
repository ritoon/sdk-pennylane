// Command exemple vérifie la connexion à la Company API avec le SDK.
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	pennylane "github.com/ritoon/sdk-pennylane"
)

func main() {
	cfg := pennylane.Config{Token: os.Getenv("PENNYLANE_TOKEN")}
	if value := os.Getenv("PENNYLANE_COMPANY_ID"); value != "" {
		id, err := strconv.ParseInt(value, 10, 64)
		if err != nil || id <= 0 {
			log.Fatal("PENNYLANE_COMPANY_ID doit être un entier positif")
		}
		cfg.CompanyID = id
	}

	client, err := pennylane.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	if err := run(context.Background(), client, os.Args[1:], os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, c *pennylane.Client, args []string, in io.Reader, out io.Writer) error {
	if len(args) == 0 {
		if err := c.Ping(ctx); err != nil {
			return err
		}
		_, err := fmt.Fprintln(out, "Connexion à Pennylane réussie.")
		return err
	}
	switch args[0] {
	case "balance":
		if len(args) == 3 {
			return balance(ctx, c, args[1], args[2], out)
		}
	case "envoyer-facture":
		if len(args) == 3 {
			return envoyerFacture(ctx, c, args[1], args[2], out)
		}
	case "creer-client":
		if len(args) == 1 {
			return creerClient(ctx, c, in, out)
		}
	case "creer-mandat":
		if len(args) == 2 {
			return creerMandat(ctx, c, args[1], in, out)
		}
	}
	return fmt.Errorf("usage : exemple [balance DEBUT FIN | envoyer-facture ID EMAIL | creer-client < client.json | creer-mandat ID_CLIENT < mandat.json]")
}
