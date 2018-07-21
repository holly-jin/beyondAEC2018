package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type (
	Config struct {
		Orbit                int
		Zoom                 int
		PresetValue          int
		DisplayMode          int
		View                 int
		LayerTree            int
		LayerCirculation     int
		LayerSiteBuilding    int
		LayerProjectBuilding int
		LayerSunShadow       int
		LayerRoad            int
	}
)

func GetSettings(w http.ResponseWriter, r *http.Request, fileId int) {
	data, err := ioutil.ReadFile(fmt.Sprintf("configs/%d.json", fileId))
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
		log.Println(err.Error())
		return
	}
	err = ioutil.WriteFile(fmt.Sprintf("configs/%d.json", fileId), data, os.ModePerm)
	if err != nil {
		http.Error(w, "Unable to write file", 500)
		return
	}
	w.WriteHeader(201)
	fmt.Fprint(w, "File Saved")
}

func main() {
	http.HandleFunc("/settings", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		idNum, err := strconv.Atoi(id)
		if err != nil || idNum < 1 {
			http.Error(w, "Missing Id", 404)
			return
		}
		switch r.Method {
		case "POST":
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			SetSettings(w, r, idNum)
			return
		case "GET":
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			GetSettings(w, r, idNum)
			return
		case "OPTIONS":
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			return
		default:
			http.Error(w, "Not found man", 404)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":3720", nil))
}
