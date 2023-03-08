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

## Multi chain run

In a particulaire case you can need to check multi same service. Exemple with ping.

For exemple I have a service on `hostA`. This service need à ping response to:

  1. `hostB` and
  2. `hostC`


For the first host `hostB` *(not real env var for ping healthcheck)*:
 ```bash
 export XXX_TYPE=ping
 export XXX_PING_HOST=hostB
 ```

For the second host, we specify an **`id`** in latest command line
 ```bash
 export XXX_MYID_TYPE=ping
 export XXX_MYID_PING_HOST=hostC
 ```

 This call healthcheck with this command:
 ```bash
 xxx && healthcheck && healthcheck -id MYID
 ```

> Pay attention to `-id` and changes in environment variables `MYID`

So une Docker, We have for exemple *(not real env var for ping healthcheck)*:

```Dockerfile
FROM traefik:v2.9.8

ADD https://github.com/hypolas/healthcheck/releases/download/v0.4.1/healthcheck_hpl-linux-amd64 /usr/bin/healthcheck
RUN chmod +x /usr/bin/healthcheck

ENV XXX_TYPE=ping
ENV XXX_PING_HOST=hostB

ENV XXX_MYID_TYPE=ping
ENV XXX_MYID_PING_HOST=hostC

HEALTHCHECK --start-period=10s --interval=10s --timeout=5s --retries=3 CMD healthcheck && healthcheck -id MYID
```
