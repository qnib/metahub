#!/bin/ash

if [[ "X${INST}" == "X" ]];then
    echo ">> This image is not optimized."
else
    echo ">> This image is optimized for '${INST}.${SIZE}' with hyperthreading turned '${HT}'"
fi
if [[ "X${SLEEP_TIME}" != "X0" ]];then
    echo -n ">> Sleeping for '${SLEEP_TIME}'"
    sleep ${SLEEP_TIME}
    echo "  DONE!"
fi