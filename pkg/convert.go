// Package pkg provides ...
package pkg

import (
	"log"
	"os/exec"
)

func convert(name string) {
	cmd := exec.Command("ffmpeg", "-i", name, name+".mp3")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
