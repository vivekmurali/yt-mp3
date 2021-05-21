package pkg

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
)

func Filename(w http.ResponseWriter, r *http.Request) {
	name := uuid.NewString()
	name += ".mp4"
	vid, err := getVideoID(r)
	if err != nil {
		ErrorHandler(w, err)
	}
	_ = vid
	go download(name, vid)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(name))
}

type Body struct {
	Url string `json:"url"`
}

func getVideoID(r *http.Request) (string, error) {
	var data Body
	gg, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal([]byte(gg), &data)
	if err != nil {
		return "", err
	}

	return data.Url, nil

}

func ErrorHandler(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
}
