# Universal `healthcheck` and `WAIT` for DevOps :D

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
Or what you want, don't hesitate to ask ;)
____

Type corresponde to `HYPOLAS_HEALTHCHECK_TYPE=\<type\>`

All configuration by environnement variables.

# Objective

- Include the same syntax for all care for hobbyists or engineers
- Remove all script command complexity
- Provide the ability to do a healthcheck without any driver installed on your operating system.

# Advanced functions

## Dynamic variable

All environment variables in the environment variable will be resolved. Exemple with *ping* healthcheck.

For an host with:
- MYENV_HOSTNAME=*`myhost`* exported
- IP = *`10.10.10.20`*

You have set bevore:
MYENV_HOSTNAME=*`myhost`*

With this settings:
```bash
export XXX_PING_HOST=\$MYENV_HOSTNAME
export XXX_PING_HOST="#CMDSTART# hostname -i #CMDEND#"
```

| Env var | Your value | Resolve value | Resolved by env | Resolved by CMD |
|---|---|---|---| --- |
| XXX_PING_HOST | \\$MYENV_HOSTNAME | *`myhost`* | - | X |
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

# Not again ready but in developpement =>

Using this `healthcheck` as `waiter`
____________________________________________
## Multi chain run

In a particulaire case you can need `wait` before start service. A databbase for exemple.

For exemple I have a service on `hostA`. This service need access to two database:

  1. `BDD-A` and
  2. `BDD-B`


For the first database `BDD-A` *(not real env var for ping healthcheck)*:
 ```bash
 export XXX_TYPE=mysql
 export XXX_WAIT=enable
 export XXX_MYSQL_HOST=BDD-A
 ```

For the second database, we specify an **`id`** in latest command line
 ```bash
export XXX_MYID_TYPE=mysql
export XXX_MYID_WAIT=enable
export XXX_MYID_MYSQL_HOST=BDD-B
 ```

 This call healthcheck with this command:
 ```bash
 xxx && healthcheck && healthcheck -id MYID && serviceStart.sh
 ```

> Pay attention to `-id` and changes in environment variables `MYID`

In this case, until the twice databases not responding, the services will wait.

Exemple with docker-compose.yml *(not real env var for MySQL healthcheck)*:

```yaml
version: '3'

services:
  web:
    image: myimage:1.0.0
    entrypoint: healthcheck && healthcheck -id MYID && serviceStart.sh
    environment:
      XXX_TYPE: mysql
      XXX_WAIT: enable
      XXX_MYSQL_HOST: BDD-A
      XXX_MYID_TYPE: mysql
      XXX_MYID_WAIT: enable
      XXX_MYID_MYSQL_HOST: BDD-B

```
