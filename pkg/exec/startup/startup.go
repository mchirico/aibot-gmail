package startup

import (
	"log"
	"os/exec"
)

func Start() {
	for {
		cmd := exec.Command("/server")
		log.Printf("Running etcd")
		err := cmd.Run()
		log.Printf("Start exited: %v", err)
		log.Printf("looping...\n\n")

	}

}
