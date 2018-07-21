package main
import (
	"log"
	"net/http"
	"strconv"
	"io/ioutil"
)

func GetSettings(w http.ResponseWriter, r *http.Request, fileId int) {
}

func SetSettings(w http.ResponseWriter, r *http.Request, fileId int) {
}

func main () {
	http.HandleFunc("/settings", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		idNum, err := strconv.Atoi(id)
		if err != nil || idNum < 1 {
			http.Error(w, "Missing Id", 404)
			return
		}
		switch r.Method {
		case "POST":
			SetSettings(w, r)
			return
		case "GET":
			GetSettings(w, r)
			return
		default:
			http.Error(w, "Not found man", 404)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":3720", nil))
}
