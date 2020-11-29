package main

import (
	d "github.com/mchirico/aibot-etcd"


	"log"
	"net"
	"os"
)

const SockAddr = "/sock/echo.sock"

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

	d.D("Removing old socket")
	if err := os.RemoveAll(SockAddr); err != nil {
		d.D("Error removing socket")

	}

	d.D("Starting net.Listen..")
	l, err := net.Listen("unix", SockAddr)
	if err != nil {
		d.D("Listening error... this is bad 0")
		log.Fatal("listen error:", err)
	}

	for {
		fd, err := l.Accept()
		if err != nil {
			d.D("l.Accept()... this is bad 0")
			log.Fatal("accept error:", err)
		}

		d.D("go echoServer ... good")
		go echoServer(fd)
		d.D("After go echoServer ... very good")
	}
}