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
	"go-stream/stream"
	"io/ioutil"
	"log"
	"strings"
)

func ExampleTestTrimSuffix() {
	reader := strings.NewReader("123")
	r := stream.NewSuffixTrimmedReader(reader, 2)
	c, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(c))
	// Output:
	// 1
}
