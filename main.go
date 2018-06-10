package main

import (
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/spf13/viper"

	"encoding/json"
	"io/ioutil"

	"github.com/maxkulish/hls_create/bitRate"
	"github.com/maxkulish/hls_create/templates"
)

// Station structure with all parameters
type Station struct {
	ID            int
	Name          string
	InputStream   string
	OutputStream  string
	Bitrate       int
	OutputBitRate map[string]int
	Volume        float32
}

// HLSPlaylist describe Playlist parameters and location
type HLSPlaylist struct {
	Station
	PlayListPath string // Folder for HLS playlists
	ExtPath      string // External server adress
	FFMPEGPath   string // ffmpeg output folder
	RunScript    string // Folder for bash scripts
}

// DefineBitrate calculate output bitrates for input
func (st Station) DefineBitrate() map[string]int {

	var result = map[string]int{}

	switch bitrate := st.Bitrate; bitrate {
	case 128:
		result = bitRate.Input128
	case 196:
		result = bitRate.Input192
	case 256:
		result = bitRate.Input256
	case 320:
		result = bitRate.Input320
	default:
		log.Printf("Can't define output bit rate: %d\n", bitrate)
	}
	return result
}

// Creates bash scripts for every station
func createStationScript(plst HLSPlaylist) {

	// Check if path /home/user/folder is exists
	if _, err := os.Stat(plst.RunScript); os.IsNotExist(err) {
		os.MkdirAll(plst.RunScript, 0755)
		log.Printf("[+] Creating folder: %s\n", plst.RunScript)
	}

	t := template.Must(template.New("bash").Parse(templates.BashTemplate))

	fileName := fmt.Sprintf("%s/%s.sh", plst.RunScript, plst.Name)
	f, err := os.Create(fileName)
	defer f.Close()

	if err != nil {
		log.Printf("[!] Can't create file: %v", err)
		os.Exit(1)
	}

	if err := os.Chmod(fileName, 0644); err != nil {
		log.Printf("[!] Can't chmod file: %v", err)
	}

	err = t.Execute(f, plst)
	if err != nil {
		log.Printf("[!] Can't execute template: %v", err)
		os.Exit(1)
	}
}

// createHLSPlaylist - creates stations.m3u8
func createHLSPlaylist(plst HLSPlaylist) {

	// Check if path is exists
	if _, err := os.Stat(plst.PlayListPath); os.IsNotExist(err) {
		os.MkdirAll(plst.PlayListPath, 0755)
		log.Printf("[+] Creating folder: %s\n", plst.PlayListPath)
	}

	t := template.Must(template.New("playlist").Parse(templates.PlaylistTemplate))

	f, err := os.Create(fmt.Sprintf("%s/%s.m3u8", plst.PlayListPath, plst.Name))
	defer f.Close()

	if err != nil {
		log.Printf("[!] Can't create file: %v", err)
		os.Exit(1)
	}

	err = t.Execute(f, plst)
	if err != nil {
		log.Printf("[!] Can't execute template: %v", err)
		os.Exit(1)
	}
}

// creates reloader.sh
func createReloader(stations []Station, scriptsFolder, reloaderPath string) {

	type ReloaderSettings struct {
		Stations     string
		ScriptsPath  string
		ReloaderPath string
	}

	parsedStations := func(stns []Station) string {
		var joined string
		for _, st := range stns {
			joined += fmt.Sprintf(" %s", st.Name)
		}

		return joined
	}(stations)

	// Check if path is exists
	if _, err := os.Stat(reloaderPath); os.IsNotExist(err) {
		os.MkdirAll(reloaderPath, 0755)
		log.Printf("[+] Creating folder: %s\n", reloaderPath)
	}

	reloader := ReloaderSettings{
		Stations:     parsedStations,
		ScriptsPath:  scriptsFolder,
		ReloaderPath: reloaderPath,
	}

	t := template.Must(template.New("reloader").Parse(templates.ReloaderScript))

	fileName := fmt.Sprintf("%s/reloader.sh", reloaderPath)
	f, err := os.Create(fileName)
	defer f.Close()

	if err != nil {
		log.Printf("[!] Can't create file: %v", err)
		os.Exit(1)
	}

	if err := os.Chmod(fileName, 0644); err != nil {
		log.Printf("[!] Can't chmod file: %v", err)
	}

	err = t.Execute(f, reloader)
	if err != nil {
		log.Printf("[!] Can't execute template: %v", err)
		os.Exit(1)
	}

}

// getStations reads stations from file stations.json
func getStations() []Station {

	raw, err := ioutil.ReadFile("./stations.json")
	if err != nil {
		log.Printf("[!] Can't read stations. Error: %v", err)
		os.Exit(1)
	}

	var stations []Station

	json.Unmarshal(raw, &stations)

	return stations
}

func main() {

	// Read project settings from config.json file
	viper.SetConfigFile("./config.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("[!] Can't read config file ./config.json")
		os.Exit(1)
	}

	// Folder to save scripts from config.json
	scriptsFolder := viper.GetString("scripts_folder")

	// ffmpeg output folder example: /hls/station/station-128.m3u8
	ffmpegFolder := viper.GetString("ffmpeg_folder")

	// External path to hls folder. Example: http://207.154.240.221/hls/
	extPath := viper.GetString("ext_path")

	// Folder to save HLS Playlist from config.json
	playListFolder := viper.GetString("playlist_folder")

	// Folder to save reloader.sh
	reloaderPath := viper.GetString("reloader_path")

	stations := getStations()
	if len(stations) == 0 {
		log.Fatal("[!] There are no stations in file. Nothing to do.")
		os.Exit(1)
	}

	// Create reloader.sh scripts
	createReloader(stations, scriptsFolder, reloaderPath)

	for _, st := range stations {

		st.OutputBitRate = st.DefineBitrate()

		plst := HLSPlaylist{
			Station:      st,
			PlayListPath: playListFolder,
			ExtPath:      extPath,
			FFMPEGPath:   ffmpegFolder,
			RunScript:    scriptsFolder,
		}

		// Create run bash scripts for every station
		createStationScript(plst)
		log.Printf("[$]\tscript for station: %s created", st.Name)

		// Create station.m3u8 file for every station
		createHLSPlaylist(plst)
		log.Printf("[$]\tHLS Playlist for station: %s created", st.Name)
	}

}
