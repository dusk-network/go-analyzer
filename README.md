# Dusk go-analyzer

A custom linter tool for enforcing CI checks on Dusk Golang repositories.

[![Go Report Card](https://goreportcard.com/badge/github.com/dusk-network/go-analyzer?style=flat-square)](https://goreportcard.com/report/github.com/dusk-network/go-analyzer)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/dusk-network/go-analyzer)](https://pkg.go.dev/github.com/dusk-network/go-analyzer)

## About

This tool should provide an easily extensible custom linter, which is easy to use. For the time being, its only purpose is to ensure that Golang source files in Dusk repositories include the correct license header.

## Usage

To use `go-analyzer` on any repo, make sure it is installed, and available in your `$PATH`:

```bash
$ go install github.com/dusk-network/go-analyzer
```

From there, you can get a full list of commands by simply running:

```bash
$ go-analyzer
```

To simply run all available lints, run:

```bash
$ go-analyzer -a
```

## License

This code is licensed under the MIT license. Please see [LICENSE](LICENSE) for further info.
