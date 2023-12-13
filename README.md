# OrderApi Service

This is the OrderApi service

Generated with

```
micro new --namespace=go.micro --type=api order-api
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.api.order-api
- Type: api
- Alias: order-api

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./order-api-api
```

Build a docker image
```
make docker
```