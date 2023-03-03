# Description

Ce binaire permet de vérifier le status d'un service:

- API => Done
- http => Done (Si le code de réponse HTML est attendu, cela est concidéré comme un check http)
- tcp => Next step

Cette outils peut être utilisé pour différents healthcheck. Il sort en Exit(0) si la status d'un service est valide ou Exit(1) si ce n'est pas le cas.

# Configuration

Utilise les variables d'environnement:

- HYPOLAS_HEALTHCHECK_HTTP_EXPECTED=down
- HYPOLAS_HEALTHCHECK_HTTP_HEADERS=firstHeader,firstValue\_\_Authorization,"Basic dGVzdDp0ZXN0MTIz"
- HYPOLAS_HEALTHCHECK_HTTP_JSON=headers\_\_Accept-Encoding
- HYPOLAS_HEALTHCHECK_HTTP_PROXY=http://exemple.com:8080
- HYPOLAS_HEALTHCHECK_HTTP_RESPONSES=200,201
- HYPOLAS_HEALTHCHECK_HTTP_TIMEOUT=5 # TimeOut in seconds
- HYPOLAS_HEALTHCHECK_HTTP_URL=https://httpbin.org/get
- HYPOLAS_HEALTHCHECK_TYPE=http

# Particularité => Les réponses aux API

## Exemple

Les API au format REST (json) sont prisent en compte. Vous pouvez allez chercher la valeur d'une réponse en profondeur.

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

HYPOLAS_HEALTHCHECK_HTTP_JSON est de type string flat avec comme séparateur un double "\_".

Je veux vérifier que la réponse du service "httpd" est "up". Pour cela je vais exporter 3 variables d'environnement:

```bash
export HYPOLAS_HEALTHCHECK_TYPE=http
export HYPOLAS_HEALTHCHECK_HTTP_JSON=services__1__status
export HYPOLAS_HEALTHCHECK_HTTP_EXPECTED=up
```

Le service index 1 étant `httpd`, le résultat sera un Exit(1).
