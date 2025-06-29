#!/bin/bash

let IDX=0

echo '['
while vulkaninfo -j=${IDX} &> /dev/null; do
    echo "$(cat VP_VULKANINFO_*.json),"
    rm -f VP_VULKANINFO_*.json
    IDX=$((IDX+1))
done
echo ']'