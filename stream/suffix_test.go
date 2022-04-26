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
	"io/ioutil"
	"strings"
	"testing"
)

func TestExtractSuffix(t *testing.T) {
	for _, scenario := range []struct {
		given    string
		suffix   int
		expected string
	}{
		{
			given:    "1234",
			suffix:   3,
			expected: "234",
		},
		{
			given:    "1234",
			suffix:   0,
			expected: "",
		},
		{
			given:    "",
			suffix:   1,
			expected: "",
		},
		{
			given:    "12",
			suffix:   3,
			expected: "12",
		},
	} {
		t.Run(fmt.Sprintf("%s-%d-%s", scenario.given, scenario.suffix, scenario.expected), func(t *testing.T) {
			reader := strings.NewReader(scenario.given)
			r := stream.NewSuffixReader(reader, scenario.suffix)
			c, err := ioutil.ReadAll(r)
			require.NoError(t, err)
			require.Equal(t, scenario.given, string(c))
			require.Equal(t, scenario.expected, string(r.Suffix()))
			require.Equal(t, 0, reader.Len())
		})
	}
}

func FuzzExtractSuffix(f *testing.F) {
	f.Add(1, 3, "1a3!4")
	f.Fuzz(func(t *testing.T, i int, b int, s string) {
		if b > 0 {
			r := stream.NewSuffixReader(strings.NewReader(s), i)
			c, err := customReadAll(r, b)
			require.NoError(t, err)
			require.Equal(t, s, string(c))
			if i < 0 {
				require.Nil(t, r.Suffix())
			} else if i <= len(s) {
				require.Equal(t, s[len(s)-i:], string(r.Suffix()))
			} else if i > len(s) {
				require.Equal(t, s, string(r.Suffix()))
			}
		}
	})
}
