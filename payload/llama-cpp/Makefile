.PHONY: latest publish

latest:
	docker build --no-cache -f Containerfile  . -t llamaherd/llama-cpp:latest
	docker tag llamaherd/llama-cpp:latest llamaherd/llama-cpp:$(shell date +%Y%m%d%H%M)-$(shell docker run -it --entrypoint /bin/cat llamaherd/llama-cpp:latest /opt/llama.cpp/commithash)

publish: latest
	docker push llamaherd/llama-cpp:latest
	docker push llamaherd/llama-cpp:$(shell date +%Y%m%d%H%M)-$(shell docker run -it --entrypoint /bin/cat llamaherd/llama-cpp:latest /opt/llama.cpp/commithash)