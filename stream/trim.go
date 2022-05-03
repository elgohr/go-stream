// Copyright 2022 Lars Gohr
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

	var err error

	if openSlots := bufferSize - len(r.buffer); openSlots > 0 {
		add := make([]byte, openSlots)
		var n int
		n, err = r.reader.Read(add)
		r.buffer = append(r.buffer, add[:n]...)
	}

	bufLen := len(r.buffer)
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
