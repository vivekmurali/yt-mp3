package pkg

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/kkdai/youtube/v2"
)

func download(name, videoID string) {
	// videoID = "https://www.youtube.com/watch?v=PCicKydX5GE"
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		// panic(err)
		log.Println(err)
	}
	// fmt.Println("\n\n\n -----\nVideo FORMAT: \n ", len(video.Formats), "\n----")
	// fmt.Println("\n\n\n -----\nVideo FORMAT: \n ", video.Formats[1].AudioChannels, "\n----")
	fmt.Printf("\n\n\n ====\n Video Format: \n %#v \n", video.Formats)
	newAudio := video.Formats.AudioChannels(2)
	// fmt.Printf("\n\n\n ====\n Video Format: \n %#v \n", newAudio)

	stream, _, err := client.GetStream(video, &newAudio[0])
	if err != nil {
		log.Println("Error here")
		log.Println(err)
	}

	file, err := os.Create(name + ".mp3")
	if err != nil {
		log.Println(err)
	}

	_, err = io.Copy(file, stream)

	if err != nil {
		// panic(err)
		log.Println(err)
	}

	// convert(name)

	// defer func() {
	// 	err = os.Remove(name)

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()
	defer file.Close()

	// log.Println("Converted and removed the video file", name)
	// return name
}
