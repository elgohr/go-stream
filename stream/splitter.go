package stream

import (
	"bytes"
	"io"
)

type SuffixSplitter struct {
	reader       io.Reader
	suffixReader *SuffixReader
}

func NewSuffixSplitter(reader io.Reader, suffixSize int) *SuffixSplitter {
	suffixReader := NewSuffixReader(reader, suffixSize)
	return &SuffixSplitter{
		reader:       NewSuffixTrimmedReader(suffixReader, suffixSize),
		suffixReader: suffixReader,
	}
}

func (r *SuffixSplitter) Read(p []byte) (int, error) {
	return r.reader.Read(p)
}

func (r *SuffixSplitter) Suffix() io.Reader {
	return bytes.NewReader(r.suffixReader.Suffix())
}
