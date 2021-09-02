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

- [x] HTML Explorer
- [ ] REST API
- [ ] Database Backend
- [ ] Mining
- [ ] Transactions
- [ ] Wallets
- [ ] P2P (Websockets)
- [ ] Unit Testing

## Getting Started <a name = "getting_started"></a>

### Prerequisites

Install the latest version of `Go`, then install external dependencies:

```bash
go get -u github.com/gorilla/mux
```

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

- https://pkg.go.dev/fmt#Printf
- https://pkg.go.dev/fmt#Sprintf
- https://pkg.go.dev/fmt#Fprint
- https://pkg.go.dev/fmt#Stringer
- https://pkg.go.dev/sync#Once
- https://pkg.go.dev/sync#Do
- https://pkg.go.dev/net/http#HandleFunc
- https://pkg.go.dev/net/http#ListenAndServe
- https://pkg.go.dev/log#Fatal
- https://pkg.go.dev/text/template#hdr-Actions
- https://pkg.go.dev/html/template#ParseFiles
- https://pkg.go.dev/html/template#Must
- https://pkg.go.dev/path/filepath#Match
- https://pkg.go.dev/encoding/json#Marshal
- https://pkg.go.dev/encoding#TextMarshaler
- https://pkg.go.dev/strconv#Atoi
- https://github.com/gorilla/mux

## Docs <a name = "docs"></a>

- https://golang.org/doc/
- https://pkg.go.dev/std
- https://golang.org/ref/spec#Variables
- https://golang.org/ref/spec#Pointer_types
- https://tour.golang.org/methods/17

## Resources <a name = "resources"></a>

- https://github.com/LarryRuane/minesim
- https://mining-simulator.netlify.app/
- https://www.blockchain.com/explorer
- https://www.youtube.com/playlist?list=PL7jH19IHhOLOJfXeVqjtiawzNQLxOgTdq
- https://stackoverflow.com/questions/29762118/range-over-array-index-in-templates
- https://andybrewer.github.io/mvp/?ref=producthunt
- https://www.digitalocean.com/community/tutorials/how-to-add-a-favicon-to-your-website-with-html
- https://en.wikipedia.org/wiki/Marshalling_(computer_science)
- https://swagger.io/specification/
- https://en.wikipedia.org/wiki/Adapter_pattern

### Go

- https://www.gorillatoolkit.org/
- https://www.digitalocean.com/community/tutorials/understanding-data-types-in-go
- https://stackoverflow.com/questions/25161774/what-are-conventions-for-filenames-in-go
- https://jogendra.dev/import-cycles-in-golang-and-how-to-deal-with-them
- https://hackthedeveloper.com/golang-server-static-files/
- https://stackoverflow.com/questions/19239449/how-do-i-reverse-an-array-in-go
- https://stackoverflow.com/questions/54858529/golang-reverse-a-arbitrary-slice
- https://github.com/golang/go/wiki/SliceTricks#reversing
- https://golangdocs.com/golang-mux-router
- https://stackoverflow.com/questions/40478027/what-is-an-http-request-multiplexer
- https://golangdocs.com/maps-in-golang
- https://golang.org/src/strconv/atoi.go?h=Atoi

### Go templates

- https://gowebexamples.com/templates/
- https://blog.gopheracademy.com/advent-2017/using-go-templates/
- https://stackoverflow.com/questions/25689829/arithmetic-in-go-templates
- https://stackoverflow.com/questions/17843311/template-and-custom-function-panic-function-not-defined
- https://stackoverflow.com/questions/38686583/golang-parse-all-templates-in-directory-and-subdirectories

### Cryptograhy

- https://en.wikipedia.org/wiki/Cryptographic_hash_function
- https://en.wikipedia.org/wiki/SHA-2
- https://en.wikipedia.org/wiki/Public-key_cryptography

### Favicons

- https://favicon.io/emoji-favicons/coin/
- https://www.freefavicon.com/
- https://icons8.com/icons/set/coin--animated
- https://www.ionos.com/tools/favicon-generator
- https://www.favicon-generator.org/search/---/Coin
- https://www.favicon.cc/?action=icon&file_id=138923

<p align="center">
    <img src=".github/golang_multiplexer.png" />
</p>
