package templates

// PlaylistTemplate template of m3u8 file
const PlaylistTemplate = `#EXTM3U{{ $bitrate := .OutputBitRate }}
#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH={{index $bitrate "middle" }}000,CODECS="mp4a.40.5"
{{ .ExtPath }}/{{ .Name }}/{{ .Name }}-{{index $bitrate "middle" }}.m3u8
#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH={{index $bitrate "high" }}000,CODECS="mp4a.40.5"
{{ .ExtPath }}/{{ .Name }}/{{ .Name }}-{{index $bitrate "high" }}.m3u8
#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH={{index $bitrate "low" }}000,CODECS="mp4a.40.5"
{{ .ExtPath }}/{{ .Name }}/{{ .Name }}-{{index $bitrate "low" }}.m3u8
`
