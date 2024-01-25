# TDD in Go

This repository contains examples of use cases on how to use TDD in Go. It is based on the [quii GitBook](https://quii.gitbook.io/learn-go-with-tests/) website.

### Useful commands

Run a specific test:
```bash
$ go test -run <test_name> # example: go test -run TestHello
```

Run the tests of a specific directory:
```bash
$ cd <directory> # example: cd hello_world
$ go test
```

Run the tests of a specific directory and show the coverage:
```bash
$ cd <directory> # example: cd hello_world
$ go test -cover
```

Run the main function of a specific file:
```bash
$ go run <file_name>.go
```

Run a benchmark test for a specific directory
```bash
$ cd <directory> # example: cd hello_world
$ go test -bench=.
```

Run all the tests of the project:
```bash
$ go test ./...
```