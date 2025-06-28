.PHONY: latest publish

latest:
	docker build --no-cache -f Containerfile  . -t llamaland/llama-cpp:latest
	docker tag llamaland/llama-cpp:latest llamaland/llama-cpp:$(shell date +%Y%m%d%H%M)-$(shell docker run -it --entrypoint /bin/cat llamaland/llama-cpp:latest /opt/llama.cpp/commithash)

publish: latest
	docker push llamaland/llama-cpp:latest
	docker push llamaland/llama-cpp:$(shell date +%Y%m%d%H%M)-$(shell docker run -it --entrypoint /bin/cat llamaland/llama-cpp:latest /opt/llama.cpp/commithash)