# Instructions pour les contributions

Toute PR fusionnée dans `main` doit être suivie de la publication d’un nouveau
tag de version sémantique sur son commit de fusion, y compris les PR de
documentation, de tests et de maintenance.

Avant la fusion, indiquer dans la PR l’incrément retenu et sa justification.
Après la fusion, appliquer la procédure de [CONTRIBUTING.md](CONTRIBUTING.md).
Le travail de livraison n’est terminé qu’après vérification du tag sur GitHub.
Ne jamais déplacer ou écraser un tag publié.

Cette procédure est exécutée par la personne ou l’agent qui fusionne la PR ;
aucun workflow de publication automatique n’est actuellement configuré.

## Exploration du dépôt

- Utiliser en priorité `rg --files` pour repérer les fichiers et `rg -n` pour
  rechercher du texte ou des symboles. Commencer par les fichiers, packages ou
  noms précis qui concernent la tâche, puis élargir si nécessaire.
- Exclure `api.gen.go` des recherches générales avec `-g '!api.gen.go'`.
  Pour vérifier un comportement du client généré, chercher le symbole exact dans
  ce fichier et lire uniquement le passage utile.
- Limiter les sorties aux correspondances et extraits nécessaires. Si `rg` est indisponible,
  utiliser `git grep` ou l’outil local équivalent.

## Utiliser et maintenir CodeGraph

- Garder `rg` pour le texte et les chemins. Pour suivre les appelants, les
  dépendances ou le rayon d’impact d’un symbole, utiliser `codegraph callers`,
  `callees` ou `impact`. Pour suivre un flux, cibler les symboles avec
  `codegraph explore --max-files 3 "<symboles>"` ou une requête MCP ciblée.
- Depuis la racine du dépôt, utiliser `codegraph init` une fois si l’index est
  absent. Le watcher maintient normalement l’index à jour. Vérifier avec
  `codegraph status`, puis lancer `codegraph sync` si l’index a pris du retard.
  Si l’index est celui d’un autre checkout, vérifier les résultats importants
  sur la branche de travail. Garder les fichiers d’index hors de Git.
- Pour vérifier si une version du CLI est disponible, utiliser
  `codegraph upgrade --check`. Après un `codegraph upgrade`, redémarrer la
  session MCP si elle utilise encore l’ancien processus, puis lancer
  `codegraph index` et `codegraph status`. En dehors d’une mise à jour, réserver
  la reconstruction complète aux problèmes persistants après synchronisation.

Voir la [référence CLI CodeGraph](https://github.com/colbymchenry/codegraph/blob/main/site/src/content/docs/reference/cli.md).
