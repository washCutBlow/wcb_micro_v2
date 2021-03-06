# Wcb_micro_v2 Service

This is the Wcb_micro_v2 service

Generated with

```
micro new --namespace=go.micro --type=service wcb_micro_v2
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.service.wcb_micro_v2
- Type: service
- Alias: wcb_micro_v2

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
./wcb_micro_v2-service
```

Build a docker image
```
make docker
```