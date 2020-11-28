package startup

import (
	"log"
	"os/exec"
	"time"
)

func Start() {
	for {
		cmd := exec.Command("/project", "etcd")
		log.Printf("Running etcd")
		err := cmd.Run()
		log.Printf("Command finished with error: %v", err)
		time.Sleep(120 * time.Second)

	}

}
