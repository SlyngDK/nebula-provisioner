# Running example

## Generating certificate

### Server

```
.../examples/ > openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 -keyout server.key -out server.crt -addext "subjectAltName = DNS:localhost,IP:127.0.0.1"
```

### Client

```
.../examples/ > openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 -keyout agent.key -out agent.crt
```

## Setup server

```
make
mkdir -p /tmp/nebula-provisioner/db
./bin/server -config examples/server.yml
```

Let the server running

## Initializing Server with encryption

```
./bin/server-client init
```

Save secrets from output.

## Unseal Server using secrets

Run as many times as you required to unseal, with different parts.

```
./bin/server-client unseal
```
