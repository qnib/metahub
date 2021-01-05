#!/bin/ash

if [[ "X${INST}" == "X" ]];then
    echo ">> This image is not optimized."
else
    echo ">> This image is optimized for '${INST}.${SIZE}' with hyperthreading turned '${HT}'"
fi