package pkg

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func Song(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "songname")
	if name == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error"))
		return
	}
	name += "_done.mp3"
	info, _ := os.Stat(name)
	if info != nil {
		w.WriteHeader(http.StatusOK)
		http.ServeFile(w, r, name)
		os.Remove(name)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("File not ready yet"))
}
