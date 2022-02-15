package stream

import (
	"fmt"
	"github.com/elgohr/go-stream/stream"
	"io/ioutil"
	"log"
	"strings"
)

func ExampleTestSplitterSuffix() {
	reader := strings.NewReader("123")
	r := stream.NewSuffixSplitter(reader, 2)
	c, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(c))
	s, err := ioutil.ReadAll(r.Suffix())
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(s))
	// Output:
	// 1
	// 23
}
