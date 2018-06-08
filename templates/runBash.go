package templates

const BashTemplate = `
#!/usr/bin/env bash

for stream in {{ .Name }}
  do mkdir -p /mnt/hls/$stream && cd /mnt/hls/$stream
     ffmpeg -i {{ .InputStream }} -c:a aac -b:a {{index .OutputBitrate "low" }}k -strict -2 -f hls -segment_list_size 10 -hls_time 3 -hls_flags delete_segments $stream-{{index .OutputBitrate "low" }}.m3u8 2> $stream-{{index .OutputBitrate "low" }}.log  &
     ffmpeg -i {{ .InputStream}} -c:a aac -b:a {{index .OutputBitrate "middle" }}k -strict -2 -f hls -segment_list_size 10 -hls_time 3 -hls_flags delete_segments $stream-{{index .OutputBitrate "middle" }}.m3u8 2> $stream-{{index .OutputBitrate "middle" }}.log &
     ffmpeg -i {{ .InputStream }} -c:a aac -b:a {{index .OutputBitrate "high" }}k -strict -2 -f hls -segment_list_size 10 -hls_time 3 -hls_flags delete_segments $stream-{{index .OutputBitrate "high" }}.m3u8 2> $stream-{{index .OutputBitrate "high" }}.log &
done
`
