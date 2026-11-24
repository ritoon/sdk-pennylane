# Contributions et versions

## Règle de publication

Chaque PR fusionnée dans `main` donne lieu à un nouveau tag annoté
`vMAJOR.MINOR.PATCH`, publié après la fusion sur le commit résultant de cette PR.
Les tags existants sont immuables. La suppression des versions antérieures à
`v1.0.0` est un nettoyage exceptionnel, pas une règle pour les versions futures.

Choisir l’incrément selon le changement le plus important de la PR :

| Incrément | Changement | Exemple |
|---|---|---|
| Major | Rupture de compatibilité de l’API publique Go | `v1.2.3` → `v2.0.0` |
| Minor | Fonctionnalité ou API ajoutée sans rupture | `v1.2.3` → `v1.3.0` |
| Patch | Correction compatible, documentation, tests, CI ou maintenance | `v1.2.3` → `v1.2.4` |

Un major remet minor et patch à zéro ; un minor remet patch à zéro.
Pour Go, une version majeure à partir de v2 nécessite aussi l’adaptation du
chemin du module et des imports avec le suffixe `/v2` (ou le major correspondant).

## Avant la fusion

1. Indiquer l’incrément et sa justification dans la description de la PR.
2. Vérifier les checks CI requis et résoudre les conflits.
3. Vérifier les tags distants pour déterminer la prochaine version stable.
   Le numéro prévu reste provisoire si une autre PR est publiée entre-temps.
4. Mettre à jour les exemples d’installation concernés par la livraison.

## Après la fusion

La personne ou l’agent qui fusionne est responsable de la publication. Cette
procédure est manuelle ; elle n’est pas exécutée par un workflow GitHub Actions.

1. Récupérer les références avec `git fetch origin --tags` et consulter
   `git ls-remote --tags origin` pour vérifier l’état distant.
2. Récupérer le SHA résultant de la fusion de la PR sur GitHub. Vérifier qu’il
   appartient à `origin/main` et que la CI de ce commit est réussie.
3. Recalculer la version depuis le dernier tag stable publié. Sérialiser les
   livraisons dans l’ordre des merges : terminer une publication avant la suivante.
4. Créer et pousser uniquement le nouveau tag, en remplaçant les valeurs :

   ```sh
   git tag -a vX.Y.Z <SHA_DE_FUSION> -m 'Release vX.Y.Z'
   git push origin refs/tags/vX.Y.Z
   git ls-remote origin 'refs/tags/vX.Y.Z*'
   ```

5. Vérifier que la référence déréférencée `refs/tags/vX.Y.Z^{}` correspond au
   SHA de fusion et communiquer le lien du tag.

Si la CI ou la publication échoue, signaler la livraison comme incomplète et
reprendre la procédure une fois le problème résolu. Si le tag existe déjà sur
le bon SHA, vérifier sa présence distante sans le recréer ; s’il désigne un
autre commit, recalculer la version sans forcer le push. Ne pas republier les
anciens tags locaux avec `git push --tags`.
