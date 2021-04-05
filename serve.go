package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	yt "github.com/vivekmurali2k/ytmp3/pkg"
)

func serveFile(w http.ResponseWriter, r *http.Request) {

	var url string
	if r.Header.Get("content-type") == "application/json" {
		var data Body
		gg, err := ioutil.ReadAll(r.Body)
		if err != nil {
			e := ErrorJson{
				Error: err.Error(),
			}
			b, _ := json.Marshal(e)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(b)
			return
		}

		err = json.Unmarshal([]byte(gg), &data)
		if err != nil {
			e := ErrorJson{
				Error: err.Error(),
			}
			b, _ := json.Marshal(e)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(b)
			return
		}

		url = data.Url

	} else {

		err := r.ParseForm()
		if err != nil {
			e := ErrorJson{
				Error: err.Error(),
			}
			b, _ := json.Marshal(e)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(b)
			return
		}
		url = r.FormValue("url")
		// fmt.Println(url)

	}
	// fmt.Println(url)
	name := yt.Download(url)
	// fmt.Println("Done successfully!")
	http.ServeFile(w, r, name+".mp3")
}

type Body struct {
	Url string `json:"url"`
}

type ErrorJson struct {
	Error string `json:"error"`
}
