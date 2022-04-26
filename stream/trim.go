package stream

import (
	"io"
)

// SuffixTrimmedReader trims the last suffixSize bytes from a stream
type SuffixTrimmedReader struct {
	buffer     []byte
	reader     io.Reader
	suffixSize int
}

func NewSuffixTrimmedReader(reader io.Reader, suffixSize int) io.Reader {
	return &SuffixTrimmedReader{
		buffer:     []byte{},
		reader:     reader,
		suffixSize: suffixSize,
	}
}

func (r *SuffixTrimmedReader) Read(p []byte) (int, error) {
	if r.suffixSize <= 0 {
		return r.reader.Read(p)
	}

	bufferSize := len(p) + r.suffixSize
	bufLen := len(r.buffer)

	var err error

	if openSlots := bufferSize - bufLen; openSlots > 0 {
		add := make([]byte, openSlots)
		var n int
		n, err = r.reader.Read(add)
		r.buffer = append(r.buffer, add[:n]...)
	}

	if err != nil {
		var n int
		if bufLen >= r.suffixSize {
			n = copy(p, r.buffer[:bufLen-r.suffixSize])
			if err == io.EOF {
				if bufLen-n == r.suffixSize {
					return n, io.EOF
				}
				return n, nil
			}
		}
		return n, err
	}

	if bufLen == bufferSize {
		n := copy(p, r.buffer[:bufLen-r.suffixSize])
		r.buffer = r.buffer[n:]
		return n, err
	}

	return 0, nil
}
