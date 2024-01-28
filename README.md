# go-collatz

[![Go Report Card](https://goreportcard.com/badge/github.com/rmaidveo/go-collatz)](https://goreportcard.com/report/github.com/rmaidveo/go-collatz)

A tool for checking the [Collatz conjecture](https://en.wikipedia.org/wiki/Collatz_conjecture).

## Installation

```
$ go install github.com/rmaidveo/go-collatz@latest
```

## Usage

```
$ go-collatz -h | -help | --help
$ go-collatz [options]
```

Options:

- `-h`, `-help`, `--help` &mdash; show the help message and exit;
- `-n INTEGER` &mdash; start number (default: 23).

## Output Example

```
23
70
35
106
53
160
80
40
20
10
5
16
8
4
2
1
```

## License

The MIT License (MIT)

Copyright &copy; 2024 rmaidveo
