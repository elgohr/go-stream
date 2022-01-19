package stream

import (
	"bufio"
	"io"
)

// SuffixTrimmedReader trims the last suffixSize bytes from a stream
type SuffixTrimmedReader struct {
	reader     *bufio.Reader
	suffixSize int
}

func NewSuffixTrimmedReader(reader io.Reader, suffixSize int) io.Reader {
	return &SuffixTrimmedReader{
		reader:     bufio.NewReader(reader),
		suffixSize: suffixSize,
	}
}

func (r *SuffixTrimmedReader) Read(p []byte) (int, error) {
	if r.suffixSize <= 0 {
		return r.reader.Read(p)
	}

	peekRead, err := r.reader.Peek(r.suffixSize * 2)
	if err != nil {
		suffixIndex := len(peekRead) - r.suffixSize
		if suffixIndex < 0 {
			return 0, err
		}
		return copy(p, peekRead[:suffixIndex]), err
	}

	return io.LimitReader(r.reader, int64(r.suffixSize)).Read(p)
}
