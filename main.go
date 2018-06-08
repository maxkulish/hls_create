package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
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

const bashTemplate = `
#!/usr/bin/env bash

for stream in {{ .Name }}
  do mkdir -p /mnt/hls/$stream && cd /mnt/hls/$stream
     ffmpeg -i {{ .InputStream }} -c:a aac -b:a {{index .OutputBitrate "low" }}k -strict -2 -f hls -segment_list_size 10 -hls_time 3 -hls_flags delete_segments $stream-{{index .OutputBitrate "low" }}.m3u8 2> $stream-{{index .OutputBitrate "low" }}.log  &
     ffmpeg -i {{ .InputStream}} -c:a aac -b:a {{index .OutputBitrate "middle" }}k -strict -2 -f hls -segment_list_size 10 -hls_time 3 -hls_flags delete_segments $stream-{{index .OutputBitrate "middle" }}.m3u8 2> $stream-{{index .OutputBitrate "middle" }}.log &
     ffmpeg -i {{ .InputStream }} -c:a aac -b:a {{index .OutputBitrate "high" }}k -strict -2 -f hls -segment_list_size 10 -hls_time 3 -hls_flags delete_segments $stream-{{index .OutputBitrate "high" }}.m3u8 2> $stream-{{index .OutputBitrate "high" }}.log &
done
`

func createStationScript(name, inputStream string, bitrate int) bool {

	newStation := Station{
		Name:          name,
		InputStream:   inputStream,
		InputBitrate:  bitrate,
		OutputBitrate: defineBitrate[bitrate],
	}

	fmt.Println(newStation)

	t := template.Must(template.New("bash").Parse(bashTemplate))

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
