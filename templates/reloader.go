package templates

// ReloaderScript - bash script which monitor ffmpeg processes and reload them
const ReloaderScript = `#!/usr/bin/env bash

WORKDIR={{ .ProjDir }}

timestamp() {
  date +"%Y-%m-%d %H:%M"
}


for stream in {{ .Stations }}

do numStr=` + "`ps aux | grep $stream | grep -v grep | awk '{print $2}' | wc -l`" +
	`	echo "$(timestamp): $stream - $numStr"
	if [ $numStr -ne 3 ]
	then
		kill $(ps aux | grep $stream | grep -v grep | awk '{print $2}')
		echo -e "=> Reloading $stream ..."
		cd $WORKDIR && sh $stream.sh
	fi
done

printf "============================\n"

exit 0

`
