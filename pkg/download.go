// Package pkg provides ...
package pkg

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/kkdai/youtube/v2"
)

func Download(videoID string) string {
	// videoID := "https://www.youtube.com/watch?v=PCicKydX5GE"
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	resp, err := client.GetStream(video, &video.Formats[0])
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	now := time.Now()

	name := fmt.Sprintf("%d%d.mp4", now.Second(), now.Nanosecond())

	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(file, resp.Body)

	if err != nil {
		panic(err)
	}

	convert(name)

	defer func() {
		err = os.Remove(name)

		if err != nil {
			log.Fatal(err)
		}
	}()
	defer file.Close()

	log.Println("Converted and removed the video file", name)
	return name
}
