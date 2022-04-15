# go-stream

[![Tests](https://github.com/elgohr/go-stream/workflows/Test/badge.svg)](https://github.com/elgohr/go-stream/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/elgohr/go-stream/branch/main/graph/badge.svg)](https://codecov.io/gh/elgohr/go-stream)
[![CodeQL](https://github.com/elgohr/go-stream/workflows/CodeQL/badge.svg)](https://github.com/elgohr/go-stream/actions/workflows/codeql-analysis.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/elgohr/go-stream)](https://goreportcard.com/report/github.com/elgohr/go-stream)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/elgohr/go-stream)](https://pkg.go.dev/github.com/elgohr/go-stream)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/gojp/goreportcard/blob/master/LICENSE)

A collection of `io.Reader` and `io.Writer` for working with big amounts of data.  
The collection aims for read-once solutions, without the need for buffering.  

## Installation

```
go get github.com/elgohr/go-stream
```

## Example on steroids
```go
w := stream.NewSizeWriter()
content := "content"
splitter := stream.NewSuffixSplitter(strings.NewReader(content), 3)
_, err := io.Copy(os.Stderr, io.TeeReader(splitter, w))
if err != nil {
   log.Fatalln(err)
}
originalBufferSize := w.Size()
_, err = io.Copy(os.Stdout, splitter.Suffix())
if err != nil {
   log.Fatalln(err)
}
```
Outputs: 
- "ent" (suffix) to os.Stdout
- "cont" to os.Stderr
- 7 as originalBufferSize

## Usage

Please see examples.
