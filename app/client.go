package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	for i := 0; i < 5; i++ {
		conn, err := net.Dial("tcp", "0.0.0.0:6379")
		if err != nil {
			return
		}

		io.Copy(os.Stdout, conn)
		conn.Close()
	}

	fmt.Println("done")

}
