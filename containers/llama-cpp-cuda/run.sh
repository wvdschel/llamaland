#!/bin/bash
set -eo pipefail

if [ -z "${LLAMA_CTX_SIZE}" ]; then
    LLAMA_CTX_SIZE=4096
fi

if [ -z "${LLAMA_GPU_LAYERS}" ]; then
    LLAMA_GPU_LAYERS=9999
fi

if [ -z "${LLAMA_OPTS}" ]; then
    LLAMA_OPTS="--cont-batching -np 1 --ctx-size ${LLAMA_CTX_SIZE} -ngl ${LLAMA_GPU_LAYERS} --threads $(nproc) --ubatch-size 4096"
fi

CMD="/opt/llama.cpp/bin/llama-server --host 0.0.0.0 --port 8080 --model /models/${MODEL_GGUF_FILENAME} ${LLAMA_OPTS}"

echo running ${CMD}
exec ${CMD}