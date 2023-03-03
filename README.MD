# Description

Ce binaire permet de vérifier le status d'un service:

- API => Done
- http => Done (Si le code de réponse HTML est attendu, cela est concidéré comme un check http)
- tcp => Next step

Cette outils peut être utilisé pour différents healthcheck. Il sort en Exit(0) si la status d'un service est valide ou Exit(1) si ce n'est pas le cas.

# Configuration

Utilise les variables d'environnement:

- HYPOLAS_HEALCHECK_HTTP_EXPECTED=down
- HYPOLAS_HEALCHECK_HTTP_HEADERS=firstHeader,firstValue\_\_Authorization,"Basic dGVzdDp0ZXN0MTIz"
- HYPOLAS_HEALCHECK_HTTP_JSON=headers\_\_Accept-Encoding
- HYPOLAS_HEALCHECK_HTTP_PROXY=http://exemple.com:8080
- HYPOLAS_HEALCHECK_HTTP_RESPONSES=200,201
- HYPOLAS_HEALCHECK_HTTP_TIMEOUT=5 # TimeOut in seconds
- HYPOLAS_HEALCHECK_HTTP_URL=https://httpbin.org/get
- HYPOLAS_HEALCHECK_TYPE=http

# Particularité => Les réponses aux API

## Exemple

Les API au format REST (json) sont prisent en compte. Vous pouvez allez chercher la valeur d'une client en profondeur.

Prenons cette réponse de l'API:

```json
{
  "hostname": "myhost",
  "services": [
    {
      "nom": "nginx",
      "status": "up"
    },
    {
      "nom": "httpd",
      "status": "down"
    }
  ]
}
```

HYPOLAS_HEALCHECK_HTTP_JSON est de type string flat avec comme séparateur un double "\_".

Je veux vérifier que la réponse du service "httpd" est "up". Pour cela je vais exporter 2 variables d'environnement:

```bash
export HYPOLAS_HEALCHECK_HTTP_JSON=services__1__status
export HYPOLAS_HEALCHECK_HTTP_EXPECTED=up
```