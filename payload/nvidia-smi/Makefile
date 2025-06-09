.PHONY: latest publish

latest:
	docker build --no-cache -f Containerfile  . -t llamaherd/nvidia-smi:latest

publish: latest
	docker push llamaherd/nvidia-smi:latest