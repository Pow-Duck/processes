package processes

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestBaseCommand(t *testing.T) {
	command := exec.Command("a.exe")
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Start()
	if err != nil {
		log.Fatalln(err)
	}
}

func TestQuery(t *testing.T) {
	command, err := RunCommand(`C:\Users\Github/AppData/Local/chia-blockchain/app-1.1.6/resources/app.asar.unpacked/daemon/chia.exe version`)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(command)
}
