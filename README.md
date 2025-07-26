# [orderedmap][repo-url] [![Go Reference][pkg-dev-img]][pkg-dev-url] [![CI Status][ci-img]][ci-url] [![MIT License][mit-img]][mit-url]

An ordered map for Golang.

This library provides the following functionalities:

- A map which is like a Go standard map, and provide same methods with Go [sync.Map](https://pkg.go.dev/sync#Map) except `CompareAndDelete` and `CompareAndSwap`. (Concurrent use is not supported.)
- `Front` and `Back` methods that iterates map entries in the order of key insertions.
- `Ldelete` and `LoadAndLdelete` methods for logical deletions, because `Store` and `Delete` are slower than Go standard map.
- `LoadOrStoreFunc` method which stores a result of a give function when an entry for the specified key is not present.
- `MarshalJSON` and `UnmarshalJSON` methods for JSON serialization and deserialization. These methods are implementations of `json.Marshaler` and `json.Unmarshaler` interfaces.

## Importing this package

```
import "github.com/sttk/orderedmap"
```

## Usage

The usage of this package is described on the overview in the go package document.

See https://pkg.go.dev/github.com/sttk/orderedmap#pkg-overview.

## Supporting Go versions

This framework supports Go 1.18 or later.

### Actual test results for each Go version:

```
% gvm-fav
Now using version go1.18.10
go version go1.18.10 darwin/amd64
ok  	github.com/sttk/orderedmap	0.298s	coverage: 99.3% of statements

Now using version go1.19.13
go version go1.19.13 darwin/amd64
ok  	github.com/sttk/orderedmap	0.363s	coverage: 99.3% of statements

Now using version go1.20.14
go version go1.20.14 darwin/amd64
ok  	github.com/sttk/orderedmap	0.314s	coverage: 99.3% of statements

Now using version go1.21.13
go version go1.21.13 darwin/amd64
ok  	github.com/sttk/orderedmap	0.310s	coverage: 99.3% of statements

Now using version go1.22.12
go version go1.22.12 darwin/amd64
ok  	github.com/sttk/orderedmap	0.314s	coverage: 99.3% of statements

Now using version go1.23.10
go version go1.23.10 darwin/amd64
ok  	github.com/sttk/orderedmap	0.324s	coverage: 99.3% of statements

Now using version go1.24.4
go version go1.24.4 darwin/amd64
ok  	github.com/sttk/orderedmap	0.331s	coverage: 99.3% of statements

Back to go1.24.4
Now using version go1.24.4
```

## License

Copyright (C) 2023-2025 Takayuki Sato

This program is free software under MIT License.<br>
See the file LICENSE in this distribution for more details.


[repo-url]: https://github.com/sttk/orderedmap
[pkg-dev-img]: https://pkg.go.dev/badge/github.com/sttk/orderedmap.svg
[pkg-dev-url]: https://pkg.go.dev/github.com/sttk/orderedmap
[ci-img]: https://github.com/sttk/orderedmap/actions/workflows/go.yml/badge.svg?branch=main
[ci-url]: https://github.com/sttk/orderedmap/actions?query=branch%3Amain
[mit-img]: https://img.shields.io/badge/license-MIT-green.svg
[mit-url]: https://opensource.org/licenses/MIT

