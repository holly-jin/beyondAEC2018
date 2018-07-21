package main
import (
	"log"
	"net/http"
	"fmt"
	"os"
	"strconv"
	"io/ioutil"
	"encoding/json"
)

type (
	Config struct {
		Version string
		BoolSetting bool
		StringSetting string
		NumSetting int
	}
)
func GetSettings(w http.ResponseWriter, r *http.Request, fileId int) {
	data, err := ioutil.ReadFile(fmt.Sprintf("configs/%d.json",fileId))
	if err != nil {
		http.Error(w, "Missing Data", 404)
		return
	}
	var c Config
	err = json.Unmarshal(data, &c)
	if err != nil {
		http.Error(w, "Bad file?", 404)
		return
	}
	jsonData, err := json.Marshal(c)
	if err != nil {
		http.Error(w, "Json encoding error", 404)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	return
}

func SetSettings(w http.ResponseWriter, r *http.Request, fileId int) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read body", 404)
		return
	}
	var c Config
	err = json.Unmarshal(data, &c)
	if err != nil {
		http.Error(w, "JSON decoding issue", 404)
		return
	}
	err = ioutil.WriteFile(fmt.Sprintf("configs/%d.json",fileId), data, os.ModePerm)
	if err != nil {
		http.Error(w, "Unable to write file", 500)
		return
	}
	w.WriteHeader(201)
	fmt.Fprint(w, "File Saved")
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
			SetSettings(w, r, idNum)
			return
		case "GET":
			GetSettings(w, r, idNum)
			return
		default:
			http.Error(w, "Not found man", 404)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":3720", nil))
}
