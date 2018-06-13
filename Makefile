ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

help: _help_

_help_:
        @echo make new - start hlsStartLinux and create new stations
        @echo make update - download new hlsStartLinux from GitHub


new:
        cd $(ROOT_DIR)
        chmod +x hlsStartLinux
        ./hlsStartLinux
		chmod -R +x scripts
		chmod +x reloader.sh


update:
        cd $(ROOT_DIR)
        rm hlsStartLinux
        wget https://github.com/maxkulish/hls_create/raw/master/bin/hlsStartLinux
        chmod +x hlsStartLinux
        chmod 0755 hlsStartLinux