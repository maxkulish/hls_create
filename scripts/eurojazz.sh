#!/usr/bin/env bash

for stream in eurojazz
  do mkdir -p /mnt/hls/$stream && cd /mnt/hls/$stream
     ffmpeg -i http://thesoundofjazz.net:8000/TheSoundOfJazzHD -c:a aac -b:a 320k -filter:a "volume=1.2" -f hls -segment_list_size 10 -hls_time 3 -hls_flags delete_segments $stream-320.m3u8 2> $stream-320.log  &
     ffmpeg -i http://thesoundofjazz.net:8000/TheSoundOfJazzHD -c:a aac -b:a 192k -filter:a "volume=1.2" -f hls -segment_list_size 10 -hls_time 3 -hls_flags delete_segments $stream-192.m3u8 2> $stream-192.log &
     ffmpeg -i http://thesoundofjazz.net:8000/TheSoundOfJazzHD -c:a aac -b:a 56k -filter:a "volume=1.2" -f hls -segment_list_size 10 -hls_time 3 -hls_flags delete_segments $stream-56.m3u8 2> $stream-56.log &
done
