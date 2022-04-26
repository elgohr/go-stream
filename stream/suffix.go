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

// SuffixReader extracts the suffix from a stream
type SuffixReader struct {
	reader     io.Reader
	suffix     []byte
	suffixSize int
}

func NewSuffixReader(reader io.Reader, suffixSize int) *SuffixReader {
	var buf []byte
	if suffixSize > 0 {
		buf = make([]byte, suffixSize)
	}
	return &SuffixReader{
		reader:     reader,
		suffix:     buf,
		suffixSize: suffixSize,
	}
}

func (r *SuffixReader) Read(p []byte) (int, error) {
	n, err := r.reader.Read(p)
	if r.suffixSize > 0 {
		r.suffixSize -= n
	}
	if r.suffix != nil {
		r.suffix = append(r.suffix, p[:n]...)[n:]
	}
	return n, err
}

// Suffix returns the last suffixSize bytes
func (r *SuffixReader) Suffix() []byte {
	if r.suffixSize > 0 {
		return r.suffix[r.suffixSize:]
	}
	return r.suffix
}
