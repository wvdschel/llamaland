.PHONY: latest publish

latest:
	docker build --no-cache -f Containerfile  . -t llamaland/nvidia-smi:latest

publish: latest
	docker push llamaland/nvidia-smi:latest