#!/bin/bash

let GPU_COUNT=0
while vulkaninfo -j=${GPU_COUNT} &> /dev/null; do
    GPU_COUNT=$((GPU_COUNT+1))
done
rm -f VP_VULKANINFO_*.json

let IDX=0
let LAST_GPU=$(expr ${GPU_COUNT} - 1)
echo '['
while vulkaninfo -j=${IDX} &> /dev/null; do
    #echo GPU ${IDX}
    echo -n "$(cat VP_VULKANINFO_*.json)"
    if [ ${IDX} != ${LAST_GPU} ]; then
        echo ','
    fi
    rm -f VP_VULKANINFO_*.json
    IDX=$((IDX+1))
done
echo -e '\n]'
