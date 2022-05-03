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

package stream_test

import (
	"fmt"
	"github.com/elgohr/go-stream/stream"
	"github.com/stretchr/testify/require"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

func TestTrimSuffix(t *testing.T) {
	for _, scenario := range []struct {
		given    string
		suffix   int
		expected string
	}{
		{
			given:    "1234",
			suffix:   1,
			expected: "123",
		},
		{
			given:    "1234",
			suffix:   3,
			expected: "1",
		},
		{
			given:    "1234",
			suffix:   0,
			expected: "1234",
		},
		{
			given:    "",
			suffix:   1,
			expected: "",
		},
		{
			given:    "12",
			suffix:   3,
			expected: "",
		},
		{
			given:    "000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			suffix:   3,
			expected: "000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		},
	} {
		t.Run(fmt.Sprintf("%s-%d-%s", scenario.given, scenario.suffix, scenario.expected), func(t *testing.T) {
			reader := strings.NewReader(scenario.given)
			r := stream.NewSuffixTrimmedReader(reader, scenario.suffix)
			c, err := ioutil.ReadAll(r)
			require.NoError(t, err)
			require.Equal(t, scenario.expected, string(c))
			require.Equal(t, 0, reader.Len())
		})
	}
}

func TestTrimSuffixWithInstantClosingReader(t *testing.T) {
	r := stream.NewSuffixTrimmedReader(instantClosingReader{}, 1)
	c, err := ioutil.ReadAll(r)
	require.NoError(t, err)
	require.Equal(t, "TES", string(c))
}

type instantClosingReader struct{}

func (i instantClosingReader) Read(p []byte) (n int, err error) {
	return copy(p, "TEST"), io.EOF
}

func FuzzTrimSuffix(f *testing.F) {
	f.Add(1, uint(3), "12a!34")
	f.Fuzz(func(t *testing.T, i int, b uint, s string) {
		r := stream.NewSuffixTrimmedReader(strings.NewReader(s), i)
		c, err := customReadAll(r, int(b))
		require.NoError(t, err)
		if i > len(s) {
			require.Equal(t, "", string(c))
		} else if i < 0 {
			require.Equal(t, string(c), string(c))
		} else {
			require.Equal(t, s[:len(s)-i], string(c))
		}
	})
}

func customReadAll(r io.Reader, bufferSize int) ([]byte, error) {
	b := make([]byte, 0, bufferSize)
	for {
		if len(b) == cap(b) {
			b = append(b, 0)[:len(b)]
		}
		n, err := r.Read(b[len(b):cap(b)])
		b = b[:len(b)+n]
		if err != nil {
			if err == io.EOF {
				return b, nil
			}
			return b, err
		}
	}
}
