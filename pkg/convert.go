package pkg

import (
	"log"
	"os"
	"os/exec"
)

func convert(name string) error {
	cmd := exec.Command("ffmpeg", "-i", name, name+".mp3")
	// var out bytes.Buffer
	// var stderr bytes.Buffer
	// cmd.Stdout = &out
	// cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Println(err)
		return err
	}
	newName := name + "_done.mp3"
	err = os.Rename(name+".mp3", newName)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
