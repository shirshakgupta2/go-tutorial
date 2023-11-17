package main

import (
	"bytes"
	"crypto/rand"

	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"
)

type FileServer struct {
}

func (fs *FileServer) start() {
	listener, err := net.Listen("tcp", ":6600")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			print(err)
		}
		go fs.readLoop(conn)
	}

}

func (fs *FileServer) readLoop(conn net.Conn) {
	//buffer := make([]byte, 2048)
	buffer := new(bytes.Buffer)

	for {
		// n, err := conn.Read(buffer)
		// copy will copy until buffer is full or there is end of File  (EOF)
		//n,err :=io.Copy(buffer,conn)

		var size int64
		//Read reads structured binary data from conn into data. Data must be a pointer
		//to a fixed-size value or a slice of fixed-size values.
		binary.Read(conn, binary.LittleEndian, &size)

		//copies n bytes (or until an error) from conn to buffer. It returns the number of bytes copied

		n, err := io.CopyN(buffer, conn, size) // Copy N means only these amount of byte to be run
		if err != nil {
			panic(err)
		}

		//file := buffer[:n]
		//fmt.Println(file)
		fmt.Println(buffer.Bytes())
		fmt.Printf("received %d bytes from network connection\n", n)

	}

}
func sendFlie(size int) error {
	file := make([]byte, size)

	// someString := "hello world\nand hello go and more"
	// myReader := strings.NewReader(someString)
	
	 // ReadFull reads exactly len(buf) bytes from rand,reader into file. It returns the number of bytes copied
	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		panic(err)
	}

	conn, err := net.Dial("tcp", ":6600")
	if err != nil {
		panic(err)
	}
	// COpy will copy keep doing until the file ends or or we stop it
	// n,err:=io.Copy(conn,bytes.NewReader(file))// convert the file to bytes

	//Write writes the binary representation of data into conn. Data must be a fixed-size value or a slice of fixed-size values
	//little endian is a order of bytes
	binary.Write(conn, binary.LittleEndian, int64(size))

	//CopyN copies n bytes (or until an error) from bytes.NewReader(file) to conn. It returns the number of bytes copied
	n, err := io.CopyN(conn, bytes.NewReader(file), int64(size)) // convert the file to bytes
	//n, err := conn.Write(file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("written %d bytes from network connection\n", n)
	return nil
}
func main() {
	go func() {
		time.Sleep(4 * time.Second)
		sendFlie(300000)
	}()

	server := &FileServer{}
	server.start()

}
