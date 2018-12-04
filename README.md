# [go-digesteve][digesteve]: cryptographic digest utilities

[![IPFN project][badge-ipfn]][org-ipfn]
[![IPFN Documentation][badge-docs]][docs]
[![See COPYING.txt][badge-copying]][COPYING]
[![GoDoc][badge-godoc]][godoc-ipfn]
[![Travis CI][badge-ci]][ci]
[![Coverage Status][coverage-badge]][coverage-status]

Go common hashing algorithms exported from [digest][] for dependencies isolation.

## Implementations

* [blake2b-simd][]
* [keccakpg][]
* [sha256-simd][]

## License

See [COPYING][COPYING] file for licensing details.

## Project

This source code is part of [IPFN](https://github.com/ipfn) – interplanetary functions project.

[ci]: https://travis-ci.org/ipfn/go-digesteve
[docs]: https://docs.ipfn.io/
[COPYING]: https://github.com/ipfn/go-digesteve/blob/master/COPYING
[badge-ci]: https://travis-ci.org/ipfn/go-digesteve.svg?branch=master
[badge-copying]: https://img.shields.io/badge/license-Apache%202.0-blue.svg?style=flat-square
[badge-docs]: https://img.shields.io/badge/documentation-IPFN-blue.svg?style=flat-square
[badge-godoc]: https://godoc.org/github.com/ipfn/go-digesteve/digesteve?status.svg
[badge-ipfn]: https://img.shields.io/badge/project-IPFN-blue.svg?style=flat-square
[coverage-badge]: https://coveralls.io/repos/github/ipfn/go-digesteve/badge.svg?branch=master
[coverage-status]: https://coveralls.io/github/ipfn/go-digesteve?branch=master
[org-ipfn]: https://github.com/ipfn
[godoc-ipfn]: https://godoc.org/github.com/ipfn/go-digesteve/digesteve
[digest]: https://github.com/ipfn/go-digest/
[digesteve]: https://github.com/ipfn/go-digesteve/
[sha256-simd]: https://github.com/crackcomm/sha256-simd
[blake2b-simd]: https://github.com/minio/blake2b-simd
[keccakpg]: https://github.com/gxed/hashland/tree/master/keccakpg
