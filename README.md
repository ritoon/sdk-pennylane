# SDK Pennylane pour Go

Ce module expose les **186 opérations de la Company API v2** répertoriées sur le portail Pennylane le 16 novembre 2026. Les requêtes, les réponses par statut HTTP et les méthodes du client sont générées en Go à partir des définitions OpenAPI de chaque fiche. Les trois fiches d’événements webhook décrivent des notifications entrantes. Les API Firm et Firm Group disposent de portails distincts. [Index officiel](https://pennylane.readme.io/llms.txt) · [Familles d’API](https://pennylane.readme.io/docs/what-apis-are-available)

## Utiliser le client

Le constructeur accepte un token Company, un token Firm ou un access token OAuth déjà obtenu. Avec un token Firm, renseigner également `Config.CompanyID` pour les opérations ciblant une société. Les applications OAuth gèrent l’obtention et la rotation du token. Le client utilise un délai de 60 secondes par défaut. Le champ facultatif `HTTPClient` permet de fournir son propre transport ou délai.

```go
package connecteur

import (
    "context"
    "fmt"
    "os"

    pennylane "github.com/ritoon/sdk-pennylane"
)

func LireJournaux(ctx context.Context) error {
    client, err := pennylane.New(pennylane.Config{Token: os.Getenv("PENNYLANE_TOKEN")})
    if err != nil {
        return err
    }
    if err := client.Ping(ctx); err != nil {
        return err
    }

    params := &pennylane.GetJournalsParams{}
    for {
        page, err := client.GetJournalsWithResponse(ctx, params)
        if err != nil {
            return err
        }
        if page.JSON200 == nil {
            return fmt.Errorf("lecture des journaux : HTTP %d", page.StatusCode())
        }
        for _, journal := range page.JSON200.Items {
            fmt.Println(journal.Code, journal.Label)
        }
        if !page.JSON200.HasMore {
            return nil
        }
        cursor, err := page.JSON200.NextCursor.Get()
        if err != nil || cursor == "" {
            return fmt.Errorf("page de journaux sans curseur de continuation")
        }
        params.Cursor = &cursor
    }
}
```

`Ping(ctx)` appelle `GET /api/external/v2/me` et vérifie que Pennylane retourne une société identifiable. Les droits requis pour les appels suivants dépendent des scopes du token. [Profil authentifié](https://pennylane.readme.io/reference/getme)

Les méthodes `…WithResponse` décodent les réponses documentées. Vérifier le statut HTTP même lorsque `err == nil`. Le corps brut reste accessible dans `Body`, et les en-têtes, dont `Retry-After`, dans `HTTPResponse.Header`. Les méthodes HTTP brutes retournent un `*http.Response` dont l’appelant ferme le corps.

Pour une pièce jointe, initialiser le champ `File` du corps multipart avec `InitFromBytes` ou `InitFromMultipart`, puis appeler la méthode typée. Les méthodes `…WithBodyWithResponse` acceptent aussi un flux `io.Reader` et son type de contenu pour un envoi en streaming.

## Vérifier et actualiser le SDK

Go 1.24 ou ultérieur. Le contrat consolidé est dans `internal/spec/company.json`. Chaque opération contient le lien vers sa fiche officielle.

```sh
go generate ./...
go test ./...
```

Pour actualiser le contrat, télécharger l’index officiel et les fiches Markdown du répertoire `reference/` dans un dossier local, en conservant leurs noms. Assembler les fiches avant de régénérer le client.

```sh
go run ./internal/spec /chemin/llms.txt /chemin/fiches > company.next.json &&
    mv company.next.json internal/spec/company.json
go generate ./...
go test ./...
```
