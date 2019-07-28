package controller

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type MyHandler struct {
}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	http.StripPrefix("/", http.FileServer(http.Dir("./upload"))).ServeHTTP(w, r)
}
func index(w http.ResponseWriter, r *http.Request) {
	title := home{Title: "Start Page"}
	t, _ := template.ParseFiles(TemplateDir + "index.html")
	t.Execute(w, title)
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle upload")
	if r.Method == http.MethodGet {
		title := home{Title: "Upload Files"}
		t, _ := template.ParseFiles(TemplateDir + "file.html")
		t.Execute(w, title)
	} else if r.Method == http.MethodPost {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Fprintf(w, "%v", "upload failure")
			return
		}
		fileext := filepath.Ext(handler.Filename)
		if check(fileext) == false {
			fmt.Fprintf(w, "%v", "upload type not allowed")
			return
		}
		filename := strconv.FormatInt(time.Now().Unix(), 10) + fileext
		f, _ := os.OpenFile(UploadDir+filename, os.O_CREATE|os.O_WRONLY, 0660)
		_, err = io.Copy(f, file)
		if err != nil {
			fmt.Fprintf(w, "%v", "upload failure on copy")
			return
		}
		fileDir, _ := filepath.Abs(UploadDir + filename)
		fmt.Fprintf(w, "%v", filename+"uploaded succe:"+fileDir)
	}
}

func Show(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	title := home{Title: "Show Page"}
	t, _ := template.ParseFiles(TemplateDir + "show.html")
	t.Execute(w, title)
}

func check(name string) bool {
	ext := []string{".exe", ".js", ".png"}
	for _, v := range ext {
		if v == name {
			return false
		}
	}
	return true
}

func StaticServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	http.StripPrefix("/file", http.FileServer(http.Dir("./upload/"))).ServeHTTP(w, r)
}

type home struct {
	Title string
}

type Router map[string]func(w http.ResponseWriter, r *http.Request)

var (
	mux Router
)

const (
	TemplateDir = "./view/"
	UploadDir   = "./upload/"
)

func init() {
	mux = make(Router)
	mux["/"] = index
	mux["/upload"] = upload
	mux["/file"] = StaticServer
	mux["/show"] = Show

}
