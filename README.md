# Josephcoin

Josephcoin is simple implementation of cryptocurrency. This coin does not have any value.

## Creator

-   Hyeonwoo KIM (@clo3olb)

## Table of Contents

-   [Getting Started](#getting-started)
    -   [Installing](#installing)
    -   [Running REST API](#running-rest-api)
    -   [Running CLI](#running-cli)
    -   [Setting port](#setting-port)

## Getting Started

### Installing

To start using Josephcoin, install Go and run `go install`:

```sh
$ go install github.com/clo3olb/josephcoin
```

if you want to see the code details and run on your go environment,

```sh
$ go get github.com/clo3olb/josephcoin
```

### Running REST API

From `$GOPATH`/bin

```sh
$ ./josephcoin -mode=rest
```

if you have fetched all the codes with `go get` then,

```sh
$ go run main.go -mode=rest
```

### Running Command Line Interface(CLI)

From `$GOPATH`/bin,

```sh
$ ./josephcoin -mode=cli
```

If you have fetched all the codes with `go get` then,

```sh
$ go run main.go -mode=cli
```

### Setting port

You can set port by indicating port number with `-port` flag.

```sh
$ ./josephcoin -mode=rest -port=4000
```
