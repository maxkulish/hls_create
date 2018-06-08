package main

import (
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/maxkulish/hls_create/templates"
)

type Station struct {
	Name          string
	InputStream   string
	InputBitrate  int
	OutputBitrate map[string]int
}

var input128 = map[string]int{
	"low":    128,
	"middle": 64,
	"high":   48,
}

var input196 = map[string]int{
	"low":    196,
	"middle": 128,
	"high":   64,
}

var input320 = map[string]int{
	"low":    320,
	"middle": 128,
	"high":   64,
}

var defineBitrate = map[int]map[string]int{
	128: input128,
	196: input196,
	320: input320,
}

func createStationScript(name, inputStream string, bitrate int) bool {

	newStation := Station{
		Name:          name,
		InputStream:   inputStream,
		InputBitrate:  bitrate,
		OutputBitrate: defineBitrate[bitrate],
	}

	fmt.Println(newStation)

	t := template.Must(template.New("bash").Parse(templates.BashTemplate))

	f, err := os.Create(fmt.Sprintf("./%s.sh", name))
	defer f.Close()

	if err != nil {
		log.Println("create file: ", err)
		return false
	}

	err = t.Execute(f, newStation)
	if err != nil {
		log.Print("execute: ", err)
		return false
	}

	return true
}

func main() {
	res := createStationScript("test", "https://test.go", 128)

	fmt.Print(res)
}
