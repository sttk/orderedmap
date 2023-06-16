# [orderedmap][repo-url] [![Go Reference][pkg-dev-img]][pkg-dev-url] [![CI Status][ci-img]][ci-url] [![MIT License][mit-img]][mit-url]

An ordered map for Golang.

This library provides the following functionalities:

- A map which is like a Go standard map, and provide same methods with Go [sync.Map](https://pkg.go.dev/sync#Map) except `CompareAndDelete` and `CompareAndSwap`. (Concurrent use is not supported.)
- `Front` and `Back` methods that iterates map entries in the order of key insertions.
- `Ldelete` and `LoadAndLdelete` methods for logical deletions, because Store and Delete are slower than Go standard map.
- `LoadOrStoreFunc` method which stores a result of a give function when an entry for the specified key is not present.


## Importing this package

```
import "github.com/sttk-go/orderedmap"
```

## Usage

The usage of this framework is described on the overview in the go package document.

See https://pkg.go.dev/github.com/sttk-go/orderedmap#pkg-overview.

## Supporting Go versions

This framework supports Go 1.18 or later.

### Actual test results for each Go version:

```
% gvm-fav
Now using version go1.18.10
go version go1.18.10 darwin/amd64
ok  	github.com/sttk-go/orderedmap	0.127s	coverage: 100.0% of statements

Now using version go1.19.5
go version go1.19.5 darwin/amd64
ok  	github.com/sttk-go/orderedmap	0.129s	coverage: 100.0% of statements

Now using version go1.20
go version go1.20 darwin/amd64
ok  	github.com/sttk-go/orderedmap	0.129s	coverage: 100.0% of statements

Back to go1.20
Now using version go1.20
```

## License

Copyright (C) 2023 Takayuki Sato

This program is free software under MIT License.<br>
See the file LICENSE in this distribution for more details.


[repo-url]: https://github.com/sttk-go/orderedmap
[pkg-dev-img]: https://pkg.go.dev/badge/github.com/sttk-go/orderedmap.svg
[pkg-dev-url]: https://pkg.go.dev/github.com/sttk-go/orderedmap
[ci-img]: https://github.com/sttk-go/orderedmap/actions/workflows/go.yml/badge.svg?branch=main
[ci-url]: https://github.com/sttk-go/orderedmap/actions
[mit-img]: https://img.shields.io/badge/license-MIT-green.svg
[mit-url]: https://opensource.org/licenses/MIT

