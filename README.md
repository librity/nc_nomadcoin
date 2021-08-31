# Nomad Coders - Nomad Coin

<p align="center">
    <img src=".github/nomadcoin.jpeg" />
</p>

## Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)
- [Notes](#notes)
- [Libs](#libs)
- [Docs](#docs)
- [Resources](#resources)

## About <a name = "about"></a>

A fully-featured blockchain and cryptocurrency using the Go programming language.

- [Take the course too!](https://nomadcoders.co/nomadcoin)
- [Certificate of Completion]()
- [Original code](https://github.com/nomadcoders/nomadcoin)

### Features

- [ ] Mining
- [ ] Transactions
- [ ] Database Backend
- [ ] Wallets
- [ ] REST API
- [ ] HTML Explorer
- [ ] P2P (Websockets)
- [ ] Unit Testing

## Getting Started <a name = "getting_started"></a>

### Prerequisites

You need to hace `GO` installed in your computer.

### Installing

```bash
$ git clone https://github.com/librity/nc_nomadcoin
$ cd nc_nomadcoin
$ go run main.go
```

Automagically analyze race conditions during compilation:

```bash
$ go build -race
```

## Usage <a name = "usage"></a>

## Notes <a name = "notes"></a>

### One-way hash functions

```go
hashFunction("sexy") => "dsdj21321wq0wjdw0jw9djcosaniqij0"
hashFunction("sexy") => "dsdj21321wq0wjdw0jw9djcosaniqij0"
hashFunction("sexyy") => "ri3j9rj2302j0ginvin0n00ivwn0inv0u"
hashFunction("ri3j9rj2302j0ginvin0n00ivwn0inv0u") => UNDEFINED
```

### Blockchain

```go
newBlockHash := hashFunction(data + previousBlockHash)
```

`data` could be anything. Any alteration to a previous block's data will
avalanche obvious changes to the next blocks' hashes.

## Libs <a name = "libs"></a>

- https://pkg.go.dev/fmt@go1.17#Printf
- https://pkg.go.dev/fmt@go1.17#Sprintf
- https://pkg.go.dev/fmt@go1.17#Fprint
- https://pkg.go.dev/sync@go1.17#Once
- https://pkg.go.dev/sync@go1.17#Do
- https://pkg.go.dev/net/http@go1.17#HandleFunc
- https://pkg.go.dev/net/http@go1.17#ListenAndServe
- https://pkg.go.dev/log@go1.17#Fatal

## Docs <a name = "docs"></a>

- https://golang.org/doc/
- https://pkg.go.dev/std
- https://golang.org/ref/spec#Variables
- https://golang.org/ref/spec#Pointer_types

## Resources <a name = "resources"></a>

- https://github.com/LarryRuane/minesim
- https://mining-simulator.netlify.app/
- https://www.youtube.com/playlist?list=PL7jH19IHhOLOJfXeVqjtiawzNQLxOgTdq
- https://www.digitalocean.com/community/tutorials/understanding-data-types-in-go
- https://en.wikipedia.org/wiki/Cryptographic_hash_function
- https://en.wikipedia.org/wiki/SHA-2
- https://en.wikipedia.org/wiki/Public-key_cryptography
