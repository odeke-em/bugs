#!/bin/bash

function si {
    for ((i=0; i<40; i++))
    do
        kill -SIGIO $1 && echo "Sent SIGIO" || break
        sleep 0.6
    done
}
si $1
