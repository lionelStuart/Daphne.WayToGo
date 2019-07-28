package test

import (
	"io"
	"net/http"
	"os"
)

func Run() {
	http.HandleFunc("/upload", uploadHandler)

	http.Handle("/staticfile", http.StripPrefix("/staticfile", http.FileServer(http.Dir("/staticfile"))))

	http.ListenAndServe(":8080", nil)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		err := r.ParseMultipartForm(10000)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		m := r.MultipartForm
		files := m.File["uploadfile"]
		for i := range files {
			func() {
				file, err := files[i].Open()
				defer file.Close()
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				dst, err := os.Create("./upload/" + files[i].Filename)
				defer dst.Close()
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				if _, err := io.Copy(dst, file); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}()

		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
