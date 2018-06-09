#!/usr/bin/env bash

for stream in business
  do mkdir -p /mnt/hls/$stream && cd /mnt/hls/$stream
     ffmpeg -i http://fex:Cssq8gEC@m.brg.ua:8080/business -c:a aac -b:a 320k  -f hls -segment_list_size 10 -hls_time 3 -hls_flags delete_segments $stream-320.m3u8 2> $stream-320.log  &
     ffmpeg -i http://fex:Cssq8gEC@m.brg.ua:8080/business -c:a aac -b:a 192k  -f hls -segment_list_size 10 -hls_time 3 -hls_flags delete_segments $stream-192.m3u8 2> $stream-192.log &
     ffmpeg -i http://fex:Cssq8gEC@m.brg.ua:8080/business -c:a aac -b:a 56k  -f hls -segment_list_size 10 -hls_time 3 -hls_flags delete_segments $stream-56.m3u8 2> $stream-56.log &
done
