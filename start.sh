#!/bin/bash

_ps=`ps --no-headers -FC tantan`

if [ "$_ps" == "" ]
then
    ./tantan > tantan.log 2>&1 &
else
    echo "Already running..."
    echo $_ps
    exit 1
fi
