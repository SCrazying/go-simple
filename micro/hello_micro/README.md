# Hello_micro Service

This is the Hello_micro service

Generated with

```
micro new go-simple/micro/hello_micro --namespace=com.skyfin --alias=hello_micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: com.skyfin.srv.hello_micro
- Type: srv
- Alias: hello_micro

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./hello_micro-srv
```

Build a docker image
```
make docker
```