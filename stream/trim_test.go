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
	"github.com/stretchr/testify/require"
	"go-stream/stream"
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
