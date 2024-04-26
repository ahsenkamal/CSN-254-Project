Netpulse is a Go library designed for peer-to-peer autodiscovery and RPC (Remote Procedure Call) between HTTP services within the same network. It operates without a central master and requires minimal configuration. Netpulse facilitates easy joining of a local network, whether as a client offering no services or as any service capable of communicating over HTTP. Its main application is for microservices within the same network, allowing them to communicate with each other seamlessly.

## Installation
Netpulse is dependent on `libzmq`, which can be installed either from source or from binaries.

1. Install WSL.
2. Install Ubuntu.
3. Install golang.
4. Install pkg-config.
5. Install libzmq3-dev.

```
export CGO_ENABLED=1
```

go build in netpulse directory

## Testing
In the `dummy` folder 4 services have been added. `service1` is a simple echo service.
```
cd ./dummy/
go run service1.go
```

The other services can be tested accordingly. Like `service2` which is addition service.
```
go run service2.go
```
