package templates

const BashTemplate = `#!/usr/bin/env bash
{{ $bitrate := .OutputBitRate }}
for stream in {{ .Name }}
  do mkdir -p /mnt/hls/$stream && cd /mnt/hls/$stream
     ffmpeg -i {{ .InputStream }} -c:a aac -b:a {{index $bitrate "low" }}k {{if gt .Volume 1.0}}-filter:a "volume={{.Volume}}"{{end}} -f hls -segment_list_size 10 -hls_time 3 -hls_flags delete_segments $stream-{{index $bitrate "low" }}.m3u8 2> $stream-{{index $bitrate "low" }}.log  &
     ffmpeg -i {{ .InputStream}} -c:a aac -b:a {{index $bitrate "middle" }}k {{if gt .Volume 1.0}}-filter:a "volume={{.Volume}}"{{end}} -f hls -segment_list_size 10 -hls_time 3 -hls_flags delete_segments $stream-{{index $bitrate "middle" }}.m3u8 2> $stream-{{index $bitrate "middle" }}.log &
     ffmpeg -i {{ .InputStream }} -c:a aac -b:a {{index $bitrate "high" }}k {{if gt .Volume 1.0}}-filter:a "volume={{.Volume}}"{{end}} -f hls -segment_list_size 10 -hls_time 3 -hls_flags delete_segments $stream-{{index $bitrate "high" }}.m3u8 2> $stream-{{index $bitrate "high" }}.log &
done
`
