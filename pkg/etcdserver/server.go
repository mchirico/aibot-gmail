package main

import (
	d "github.com/mchirico/aibot-etcd"


	"log"
	"net"
	"os"
)

const SockAddr = "/tmp/echo.sock"

func echoServer(c net.Conn) {
	for {
		buf := make([]byte, 1024)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[0:nr]
		log.Printf("data: %v\n",data)
		println("Server got:", string(data))
		d.D(string(data))

	}
}

func main() {


	if err := os.RemoveAll(SockAddr); err != nil {
		log.Fatal(err)
	}


	l, err := net.Listen("unix", SockAddr)
	if err != nil {
		log.Fatal("listen error:", err)
	}

	for {
		fd, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go echoServer(fd)
	}
}