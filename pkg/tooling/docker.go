package tooling

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

// DockerLogin uses a docker client to login to metahub
func DockerLogin(pass, user, registry string) (err error) {
	log.Println("Use docker login via client")
	cmd := fmt.Sprintf("docker --log-level=debug --debug login --password %s --username %s  %s", pass, user, registry)
	out, err := exec.Command("ash", "-c", cmd).Output()
	if err != nil {
		log.Printf("CMD: %s", cmd)
		log.Printf("Err: %s", err.Error())
		return fmt.Errorf("Failed to execute command: %s", cmd)
	}
	time.Sleep(2)
	// TODO: That's rather ugly, but the login data is stored per user...
	err = os.Chmod("/root/.docker/config.json", 0644)
	if err != nil {
		return fmt.Errorf("Failed to chmod /root/.docker/config.json: %s", err.Error())
	}
	log.Println(string(out))
	return
}
