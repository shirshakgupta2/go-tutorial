package main

import (
	"fmt"
	"io"
	"os"
)

// alphaReader is now capable of reading from any reader implementation.
type randReader struct {
	reader io.Reader
}

func newRandReader(reader io.Reader) *randReader {
	return &randReader{reader: reader}
}

func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

func (a *randReader) Read(p []byte) (int, error) {
	n, err := a.reader.Read(p)
	if err != nil {
		return n, err
	}
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if char := alpha(p[i]); char != 0 {
			buf[i] = char
		}
	}
	/*The copy built-in function copies elements from a source slice(buf) into a destination slice(p).
 	(As a special case, it also will copy bytes from a string to a slice of bytes.)
 	The source and destination may overlap. Copy returns the number of elements copied,
 	which will be the minimum of len(src) and len(dst).*/
	copy(p, buf)
	return n, nil
}

//	func main() {
//		// use an io.Reader as source for randReader
//		reader := newRandReader(strings.NewReader("Hello! It's 9am, where is the sun?"))
//		p := make([]byte, 4)
//		for {
//			n, err := reader.Read(p)
//			if err == io.EOF {
//				break
//			}
//			fmt.Print(string(p[:n]))
//		}
//		fmt.Println()
//	}
func main() {
	// use an os.File as source for alphaReader
	file, err := os.Open("./main.go")// this will treat the the main .go as an string and will 
	// fetch out all the datta that is to be read from the main .go file
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	reader := newRandReader(file)
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
	fmt.Println()
}
