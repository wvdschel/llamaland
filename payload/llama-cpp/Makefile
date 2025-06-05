.PHONY: build publish

build:
	docker build --no-cache -f Containerfile  . -t computemaestro/llama-cpp:latest
	docker tag computemaestro/llama-cpp:latest computemaestro/llama-cpp:$(date +%Y%m%d%H%M)-$(docker run maestro/llama-cpp:latest cat /opt/llama.cpp/commithash)