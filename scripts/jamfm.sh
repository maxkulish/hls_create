#!/usr/bin/env bash

for stream in jamfm
  do mkdir -p /mnt/hls/$stream && cd /mnt/hls/$stream
     ffmpeg -i http://cast.radiogroup.com.ua:8000/jamfm.aac -c:a aac -b:a 128k -filter:a "volume=1.2" -f hls -segment_list_size 10 -hls_time 3 -hls_flags delete_segments $stream-128.m3u8 2> $stream-128.log  &
     ffmpeg -i http://cast.radiogroup.com.ua:8000/jamfm.aac -c:a aac -b:a 96k -filter:a "volume=1.2" -f hls -segment_list_size 10 -hls_time 3 -hls_flags delete_segments $stream-96.m3u8 2> $stream-96.log &
     ffmpeg -i http://cast.radiogroup.com.ua:8000/jamfm.aac -c:a aac -b:a 56k -filter:a "volume=1.2" -f hls -segment_list_size 10 -hls_time 3 -hls_flags delete_segments $stream-56.m3u8 2> $stream-56.log &
done
