# TanTan-test-go

This is a project of tantan backend developer test

    ./tantan -h

    Usage of ./tantan:
      -c string
            configuration file (default "config.json")
      -v    show version


## Installation

Install:
    
    cd Tantan-test-go

    ./build.sh

Tantan server will use the user postgres of postresql, make sure it can be connected. 

## Start

    ./start.sh

This command will start tantan server with default configuration file (config.json), server will listen on port 8088, and you can change it.

## Stop

    ./stop.sh

Stop tantan server immediately.
