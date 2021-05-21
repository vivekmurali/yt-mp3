package pkg

import (
	"bytes"
	"fmt"
	"os/exec"
)

func convert(name string) error {
	cmd := exec.Command("ffmpeg", "-i", name, name+".mp3")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}
	return err
	// err := cmd.Run()
	// return err
}
