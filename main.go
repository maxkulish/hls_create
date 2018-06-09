package main

import (
	"fmt"
	"html/template"
	"log"
	"os"

	"encoding/json"
	"io/ioutil"

	"github.com/maxkulish/hls_create/bitRate"
	"github.com/maxkulish/hls_create/templates"
)

type Station struct {
	ID            int
	Name          string
	InputStream   string
	OutputStream  string
	Bitrate       int
	OutputBitRate map[string]int
	Volume        float32
}

func (Station) DefineBitrate(btrt int) map[string]int {

	var result = map[string]int{}

	switch btrt {
	case 128:
		result = bitRate.Input128
	case 196:
		result = bitRate.Input192
	case 256:
		result = bitRate.Input256
	case 320:
		result = bitRate.Input320
	default:
		log.Printf("Can't define output bit rate: %d\n", btrt)
	}
	return result
}

func createStationScript(name, inputStream, path string, btrt int, vol float32) {

	newStation := Station{
		Name:        name,
		InputStream: inputStream,
		Bitrate:     btrt,
		Volume:      vol,
	}

	// Check if path /home/user/folder is exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
		log.Printf("Creating folder: %s\n", path)
	}

	newStation.OutputBitRate = newStation.DefineBitrate(btrt)

	log.Println(newStation)

	t := template.Must(template.New("bash").Parse(templates.BashTemplate))

	f, err := os.Create(fmt.Sprintf("%s/%s.sh", path, name))
	defer f.Close()

	if err != nil {
		log.Printf("Can't create file: ", err)
		os.Exit(1)
	}

	err = t.Execute(f, newStation)
	if err != nil {
		log.Printf("Can't execute template: ", err)
		os.Exit(1)
	}
}

func createHLSPlaylist() {

}

func getStations() []Station {

	raw, err := ioutil.ReadFile("./stations.json")
	if err != nil {
		log.Printf("Can't read stations. Error: %v", err)
		os.Exit(1)
	}

	var stations []Station

	json.Unmarshal(raw, &stations)

	return stations
}

func main() {

	// Folder to save scripts
	scriptsFolder := "./scripts"

	stations := getStations()

	for _, s := range stations {
		createStationScript(
			s.Name,
			s.InputStream,
			scriptsFolder,
			s.Bitrate,
			s.Volume)

		log.Printf("[%s] script for station created with result", s.Name)
	}

	//res := createStationScript(
	//	"test",
	//	"https://test.go",
	//	"/home/max/go/src/github.com/maxkulish/scripts",
	//	320,
	//	1.3)
	//
	//fmt.Println(res)
}
