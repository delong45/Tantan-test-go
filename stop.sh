#!/bin/bash

pid=`ps --no-headers -FC tantan | awk '{print $2}'`

kill $pid

if [ $? -ne 0 ]
then
    echo "Failed to stop Tantan server"
    exit 1
fi
