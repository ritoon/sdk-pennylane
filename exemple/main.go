// Command exemple vérifie la connexion à la Company API avec le SDK.
package main

import (
	"context"
	"fmt"
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
	if err := client.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connexion à Pennylane réussie.")
}
