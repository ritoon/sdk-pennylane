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

## Exploration locale du code

- Privilégier `rg --files` pour repérer les fichiers et `rg -n` pour rechercher
  un symbole ou un motif. Si `rg` est indisponible, utiliser `find` ou `grep`.
- Commencer par les chemins et types de fichiers pertinents. Limiter le contexte
  affiché aux lignes utiles, puis élargir la recherche si nécessaire.
- Pour un comportement du SDK, commencer par le code maintenu à la main et les
  tests. Si le code généré est pertinent, rechercher le symbole exact dans
  `api.gen.go` et lire uniquement l’extrait concerné.
