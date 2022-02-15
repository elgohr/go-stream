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

func TestSuffixSplitter(t *testing.T) {
	for _, scenario := range []struct {
		given          string
		suffix         int
		expectedRead   string
		expectedSuffix string
	}{
		{
			given:          "1234",
			suffix:         3,
			expectedRead:   "1",
			expectedSuffix: "234",
		},
		{
			given:          "1234",
			suffix:         0,
			expectedRead:   "1234",
			expectedSuffix: "",
		},
		{
			given:          "",
			suffix:         1,
			expectedRead:   "",
			expectedSuffix: "",
		},
		{
			given:          "12",
			suffix:         3,
			expectedRead:   "",
			expectedSuffix: "12",
		},
	} {
		t.Run(fmt.Sprintf("%s-%d-%s", scenario.given, scenario.suffix, scenario.expectedRead), func(t *testing.T) {
			reader := strings.NewReader(scenario.given)
			r := stream.NewSuffixSplitter(reader, scenario.suffix)
			c, err := ioutil.ReadAll(r)
			require.NoError(t, err)
			require.Equal(t, scenario.expectedRead, string(c))
			s, err := ioutil.ReadAll(r.Suffix())
			require.NoError(t, err)
			require.Equal(t, scenario.expectedSuffix, string(s))
			require.Equal(t, 0, reader.Len())
		})
	}
}
