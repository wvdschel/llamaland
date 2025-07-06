OCI_CMD ?= docker
IMAGE_PREFIX=quay.io/llamaland
IMAGES=$(shell find . -maxdepth 1 -mindepth 1 -type d | cut -d/ -f2- | sort)
PUBLISH_TARGETS=$(patsubst %,publish-%,$(IMAGES))

# Epoch changes every 4 days, forcing cached docker images to be rebuilt
CACHE_EPOCH=$(shell expr $(shell date +%s) / 3600 / 24 / 14)

.PHONY: images publish force-rebuild

images: $(IMAGES)

publish: $(PUBLISH_TARGETS)

force-rebuild:

publish-%: % force-rebuild
	@echo PUBLISHING $<
	$(OCI_CMD) push $(IMAGE_PREFIX)/$<:latest 
	$(OCI_CMD) run -it --entrypoint /versioninfo $(IMAGE_PREFIX)/$<:latest | dos2unix | xargs -i $(OCI_CMD) push $(IMAGE_PREFIX)/$<:{}

%: %/Containerfile force-rebuild
	@echo BUILDING $@
	$(OCI_CMD) build --build-arg CACHE_EPOCH=$(CACHE_EPOCH) -t $(IMAGE_PREFIX)/$@:latest -f $< $@
	$(OCI_CMD) run -it --entrypoint /versioninfo $(IMAGE_PREFIX)/$@:latest | dos2unix | xargs -i $(OCI_CMD) tag $(IMAGE_PREFIX)/$@:latest $(IMAGE_PREFIX)/$@:{}
