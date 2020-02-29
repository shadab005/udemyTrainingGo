package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	fmt.Println("Hello from reader!")
	b := make([]byte, 4)
	b[0] = 'A'
	fmt.Println(b[:len(b)])
	testReader()
}

func testReader() {
	reader := strings.NewReader("Clear is better than clever")
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			fmt.Println("\nEOF")
			break
		}
		fmt.Println(string(p[:n]))
	}

}

//TODO
// 1. Creater reader object given the file name
// 2. Read file line by line
// 3. Read file in single string
// 4. Decode file using json decoder. File content has list of encoded value
// 5. Read file in chunks of size in byte array
