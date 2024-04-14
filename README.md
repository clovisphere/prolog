# prolog

> A commit log prototype

This example application is from the excellent [Distributed Services with Go](https://pragprog.com/titles/tjgo/distributed-services-with-go) 📖 by Travis Jeffrey.

## What do you need to run the application?

- [Go 1.13+](https://go.dev/dl/)
- [curl](https://curl.se) or any other fancy tool that you (can) use to make HTTP calls, i.e. [httpie](https://httpie.io), [bruno](https://www.usebruno.com), etc.

### Test the API

> Start the service

```console
$ go run ./cmd/server/main.go
```

> Add a record

```console
$ curl -X POST localhost:8080 -d '{"record": {"value": "TGV0J3MgR28gIzEK"}}'
```

> Read/Retrieve record

```console
$ curl -X GET localhost:8080 -d '{"offset": 0}'
```
