# Description from DevOps for DevOps :D

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

# Fonction avancée

## Description

Vous pouvez passé des commandes dans vos variable d'environnement. Par exemple dans Docker, un check sur un host static n'a que peut de sens. L'adresse IP et le hostname de votre conainer Docker pouvant changer régulièrement.

Cette problématique est résolue ici en donnant la possibilité d'appeler des commandes externes. Par exemple, l'IP de l'hôte (ou du Docker) est très utile pour un healthcheck locale. La commande `hostname -i` permet de récupérer l'IP alors que les variables d'environnements sont statiques.

Toutes les valeurs entre les balises `#CMDSTART#` et `#CMDEND#` seront lancées dans le shell de votre système. Le retour de votre commande dans le terminal mettra à jour votre variable d'environnement. Exemple:

- export HYPOLAS_HEALTHCHECK_HTTP_URL=http://#CMDSTART# hostname -i #CMDEND#:8082/ping

Engendra l'execution par votre système de `hostname -i`. Votre terminal lui répondra par exemple 10.10.10.20. La variable d'environnement sera au final:

- export HYPOLAS_HEALTHCHECK_HTTP_URL=http://10.10.10.20:8082/ping

## Exemple avec Docker (Traefik)

```Dockerfile
FROM traefik:v2.9.8

ADD https://github.com/hypolas/healthcheck/releases/download/v0.2.0/healthcheck_hpl-linux-amd64 /usr/bin/healthcheck
RUN chmod +x /usr/bin/healthcheck

ENV HYPOLAS_HEALTHCHECK_TYPE=http
ENV HYPOLAS_HEALTHCHECK_HTTP_EXPECTED=OK
ENV HYPOLAS_HEALTHCHECK_HTTP_URL="http://#CMDSTART# hostname -i #CMDEND#:8082/ping"

HEALTHCHECK --start-period=10s --interval=10s --timeout=5s --retries=3 CMD healthcheck
```

# Debug

Si vous constatez des dysfonctionnements, vous pouvez lancer en mode `debug`.

```bash
healthcheck -d
```

Si vous me faites une issue, ce log me sera grandement utile ;-).
