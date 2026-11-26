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
- Exclure les fichiers générés des recherches générales avec `-g '!*.gen.go'`.
  Pour vérifier un comportement du client généré, consulter `docs/code-genere.md`,
  chercher le symbole exact dans le fichier du domaine et lire uniquement le passage utile.
- Limiter les sorties aux correspondances et extraits nécessaires. Utiliser
  un index de code si la recherche textuelle ne suffit pas, par exemple pour
  résoudre des références entre packages ou dépôts. Si `rg` est indisponible,
  utiliser `git grep` ou l’outil local équivalent.
