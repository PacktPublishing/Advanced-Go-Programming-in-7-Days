package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"io"
)

type LowercaseReader struct {
	text string
}

func NewLowercaseReader(text string) *LowercaseReader {
	return &LowercaseReader{text: text}
}

func (r *LowercaseReader)Read(p []byte) (int, error)  {
	buf := make([]byte, len(r.text))

	for i := 0; i < len(buf); i++ {
		buf[i] = r.text[i] | 0x20
	}

	n := copy(p, buf)
	return n, io.EOF
}

func main() {
	// Example of implementing a custom io.Reader
	r := NewLowercaseReader("ALL CAPITALS")
	resp, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal("error: something was wrong when converting to lowecase")
	}

	fmt.Println(string(resp))
}
