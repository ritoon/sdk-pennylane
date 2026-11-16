# Payloads d’erreur

[Retour aux opérations](../operations.md)

## GetJournalsStatus400

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | anyOf | racine |  |
| `$ (anyOf 1)` | object | selon variante | additionalProperties `false` |
| `$ (anyOf 1).error` | string | obligatoire |  |
| `$ (anyOf 1).status` | integer | obligatoire |  |
| `$ (anyOf 2)` | object | selon variante | additionalProperties `false` |
| `$ (anyOf 2).message` | string | obligatoire |  |
| `$ (anyOf 3)` | object | selon variante | additionalProperties `false` |
| `$ (anyOf 3).code` | string | obligatoire | enum `["InvalidDateFormat","InvalidDateTimeFormat","InvalidEmailFormat","InvalidPattern","InvalidUUIDFormat","LessThanExclusiveMinimum","LessThanMinimum","LessThanMinItems","LessThanMinLength","MoreThanExclusiveMaximum","MoreThanMaximum","MoreThanMaxItems","MoreThanMaxLength","NotAMultipartFile","NotAnyOf","NotEnumInclude","NotExistContentTypeDefinition","NotExistDiscriminatorMappedSchema","NotExistDiscriminatorPropertyName","NotExistPropertyDefinition","NotExistRequiredKey","NotExistStatusCodeDefinition","NotNullError","NotOneOf","ValidateError"]` |
| `$ (anyOf 3).message` | string | obligatoire |  |
| `$ (anyOf 4)` | object | selon variante | additionalProperties `false` |
| `$ (anyOf 4).code` | string | obligatoire | enum `["InvalidDateFormat","InvalidDateTimeFormat","InvalidEmailFormat","InvalidPattern","InvalidUUIDFormat","LessThanExclusiveMinimum","LessThanMinimum","LessThanMinItems","LessThanMinLength","MoreThanExclusiveMaximum","MoreThanMaximum","MoreThanMaxItems","MoreThanMaxLength","NotAMultipartFile","NotAnyOf","NotEnumInclude","NotExistContentTypeDefinition","NotExistDiscriminatorMappedSchema","NotExistDiscriminatorPropertyName","NotExistPropertyDefinition","NotExistRequiredKey","NotExistStatusCodeDefinition","NotNullError","NotOneOf","ValidateError"]` |
| `$ (anyOf 4).field` | string | obligatoire |  |
| `$ (anyOf 4).message` | string | obligatoire |  |
| `$ (anyOf 5)` | object | selon variante | additionalProperties `false` |
| `$ (anyOf 5).code` | string | obligatoire | enum `["InvalidDateFormat","InvalidDateTimeFormat","InvalidEmailFormat","InvalidPattern","InvalidUUIDFormat","LessThanExclusiveMinimum","LessThanMinimum","LessThanMinItems","LessThanMinLength","MoreThanExclusiveMaximum","MoreThanMaximum","MoreThanMaxItems","MoreThanMaxLength","NotAMultipartFile","NotAnyOf","NotEnumInclude","NotExistContentTypeDefinition","NotExistDiscriminatorMappedSchema","NotExistDiscriminatorPropertyName","NotExistPropertyDefinition","NotExistRequiredKey","NotExistStatusCodeDefinition","NotNullError","NotOneOf","ValidateError"]` |
| `$ (anyOf 5).message` | string | obligatoire |  |
| `$ (anyOf 5).payload` | string | obligatoire |  |
| `$ (anyOf 6)` | object | selon variante | additionalProperties `false` |
| `$ (anyOf 6).code` | string | obligatoire | enum `["InvalidDateFormat","InvalidDateTimeFormat","InvalidEmailFormat","InvalidPattern","InvalidUUIDFormat","LessThanExclusiveMinimum","LessThanMinimum","LessThanMinItems","LessThanMinLength","MoreThanExclusiveMaximum","MoreThanMaximum","MoreThanMaxItems","MoreThanMaxLength","NotAMultipartFile","NotAnyOf","NotEnumInclude","NotExistContentTypeDefinition","NotExistDiscriminatorMappedSchema","NotExistDiscriminatorPropertyName","NotExistPropertyDefinition","NotExistRequiredKey","NotExistStatusCodeDefinition","NotNullError","NotOneOf","ValidateError"]` |
| `$ (anyOf 6).field` | string | obligatoire |  |
| `$ (anyOf 6).message` | string | obligatoire |  |
| `$ (anyOf 6).payload` | string | obligatoire |  |

## GetJournalsStatus401

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.error` | string | obligatoire | example `"The access token is invalid"` |
| `$.status` | integer | obligatoire | example `401` |

## GetJournalsStatus403

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.error` | string | obligatoire | example `"Access to this resource requires scope \"ledger\"."` |
| `$.status` | integer | obligatoire | example `403` |

## GetJournalsStatus404

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.error` | string | obligatoire | example `"Not Found"` |
| `$.status` | integer | obligatoire | example `404` |

## ImportCustomerInvoicesStatus503

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.error` | string | obligatoire | example `"The service is temporarily unavailable"` |
| `$.status` | integer | obligatoire | example `503` |

## PostJournalsStatus422

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.error` | string | obligatoire | example `"Entry lines are not balanced"` |
| `$.status` | integer | obligatoire | example `422` |

## SendByEmailCommercialDocumentStatus409

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.error` | string | obligatoire | example `"Conflict with the current state of the target resource. Resource already exists."` |
| `$.status` | integer | obligatoire | example `409` |
