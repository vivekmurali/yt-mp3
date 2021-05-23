package pkg

import (
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
	// fmt.Printf("\n\n\n ====\n Video Format: \n %#v \n", video.Formats)
	newAudio := video.Formats.FindByItag(140)
	// fmt.Printf("\n\n\n ====\n Video Format: \n %#v \n", newAudio)

	// stream, _, err := client.GetStream(video, &newAudio[0])
	stream, _, err := client.GetStream(video, newAudio)
	if err != nil {
		log.Println("Error here")
		log.Println(err)
	}

	file, err := os.Create(name)
	if err != nil {
		log.Println(err)
	}

	_, err = io.Copy(file, stream)

	if err != nil {
		// panic(err)
		log.Println(err)
	}

	err = convert(name)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = os.Remove(name)

		if err != nil {
			log.Fatal(err)
		}
	}()
	defer file.Close()

	// log.Println("Converted and removed the video file", name)
	// return name
}
