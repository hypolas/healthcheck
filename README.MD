# Universal healthcheck for DevOps :D

*The easiest healthcheck.*

This healthcheck is a hub of different healthcheck produced by community.

| For | Type | When | Source |
|---|---|---|---|
| http | http | done | https://github.com/hypolas/hypolashlckhttp
| API | http | done | https://github.com/hypolas/hypolashlckhttp
| ICMP | ping | done | https://github.com/hypolas/hypolashlckping
| PostgreSQL | "" | Soon | 
| MySQL | "" | Soon | 
| Redis | "" | Planed | 

____
Or what you want ! just ask it ;)
____

Type corresponde to `HYPOLAS_HEALTHCHECK_TYPE=\<type\>`

All configuration by environnement variables.

# Objective

- Include same synctax for all healcheck for amators or ingeniors
- Remove all complexity of script command
- Give possibility to do healthcheck without any driver installed on your OS.

# Advanced functions

## Dynamic variable

All environnement variable in environment varible will be resolve. Exemple with *ping* healthcheck.

For host with:
- hostmane = *`myhost`*
- IP = *`10.10.10.20`*

So with setting:
```bash
export XXX_PING_HOST=\$hostname
export XXX_PING_HOST="#CMDSTART# hostname -i #CMDEND#"
```

| Env var | Your value | Resolve value | Resolved by env | Resolved by CMD |
|---|---|---|---| --- |
| XXX_PING_HOST | \\$hostname | *`myhost`* | - | X |
| XXX_PING_HOST | `#CMDSTART#` hostname -i `#CMDEND#` | *`10.10.10.20`* | X | - |


### Exemple for Docker (Traefik)

```Dockerfile
FROM traefik:v2.9.8

ADD https://github.com/hypolas/healthcheck/releases/download/v0.4.1/healthcheck_hpl-linux-amd64 /usr/bin/healthcheck
RUN chmod +x /usr/bin/healthcheck

ENV HYPOLAS_HEALTHCHECK_TYPE=http
ENV HYPOLAS_HEALTHCHECK_HTTP_EXPECTED=OK
ENV HYPOLAS_HEALTHCHECK_HTTP_URL="http://#CMDSTART# hostname -i #CMDEND#:8082/ping"

HEALTHCHECK --start-period=10s --interval=10s --timeout=5s --retries=3 CMD healthcheck
```

## Multi run stack

I n particulaire case you can need to check multi same service. Exemple with ping.

For exemple I have a service on `hostA`. This service need à ping response to:

  1. `hostB` and
  2. `hostC`



You first config for `hostB`
 ```bash
 export XXX_TYPE=`ping`
 export XXX_PING_HOST=`hostB`
 ```

For second host we spécifie an **`id`** in latest command line
 ```bash
 export XXX_MYID_TYPE=`ping`
 export XXX_MYID_PING_HOST=`hostB`
 ```

 This call healthcheck with this command:
 ```bash
 xxx && healthcheck && healthcheck -id MYID
 ```

> :warning Pay attention to `-id` and changes in environment variables `MYID`


