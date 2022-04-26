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
