package types

import "io"

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
	Gatherer  IGather
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}


type IGather interface {
	Fetch(url string) (io.Reader, error)
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}

