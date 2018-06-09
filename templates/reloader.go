package templates

const ReloaderScript = `#!/usr/bin/env bash

WORKDIR=/home/fon/scripts

timestamp() {
  date +"%Y-%m-%d %H:%M"
}


for stream in business djfm powerfm renessans radioparadise shanson eurojazz jamfm loungefm

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
