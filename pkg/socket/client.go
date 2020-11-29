package socket

// Ref:
//   https://stackoverflow.com/questions/2886719/unix-sockets-in-go

import (
	"io"
	"log"
	"net"
)

const SockAddr = "/sock/echo.sock"

func Reader(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			return
		}
		println("Client got:", string(buf[0:n]))
	}
}

func Client(msg string) {
	c, err := net.Dial("unix", SockAddr)
	if err != nil {
		log.Printf("err: %s\n", err)
	}
	defer c.Close()

	// go Reader(c)

	n, err := c.Write([]byte(msg))
	if err != nil {
		log.Printf("write error:%s\n", err)

	}
	log.Printf("wrote bytes: %v\n", n)

}
