# Utiliser le SDK Go Pennylane

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

## Renseigner un corps de requête et lire le résultat

Dans la fiche de [création d’un compte bancaire](endpoints/postBankAccount.md), le champ JSON obligatoire `name` correspond au champ Go `Name`. Cet exemple crée un compte dans la société du client et nécessite le scope `bank_accounts:all`.

```go
func CreerCompte(ctx context.Context, client *pennylane.Client, nom string) (int64, error) {
    body := pennylane.PostBankAccountJSONRequestBody{Name: nom}
    resp, err := client.PostBankAccountWithResponse(ctx, nil, body)
    if err != nil {
        return 0, err
    }
    if resp.JSON201 == nil {
        return 0, fmt.Errorf("création du compte : HTTP %d", resp.StatusCode())
    }
    return resp.JSON201.Id, nil
}
```

Le client ajoute l’authentification et la société configurée. Ici, `nil` laisse les paramètres HTTP à leurs valeurs par défaut. Le résultat de la création se lit dans `JSON201` car le contrat prévoit le statut 201. La fiche indique aussi les statuts d’erreur et leurs champs documentés.
