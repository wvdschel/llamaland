OCI_CMD ?= docker
IMAGE_PREFIX=quay.io/llamaland
IMAGES=$(shell find . -maxdepth 1 -mindepth 1 -type d | cut -d/ -f2- | sort)
PUBLISH_TARGETS=$(patsubst %,publish-%,$(IMAGES))

.PHONY: images publish #$(IMAGES) $(PUBLISH_TARGETS)

images: $(IMAGES)

publish: $(PUBLISH_TARGETS)

publish-%: %
	@echo PUBLISHING $<
	$(OCI_CMD) push $(IMAGE_PREFIX)/$<:latest $(IMAGE_PREFIX)/$<:$(shell $(OCI_CMD) run -it --entrypoint /versioninfo $(IMAGE_PREFIX)/$<:latest)

%: %/Containerfile
	@echo BUILDING $@
	$(OCI_CMD) build --no-cache -t $(IMAGE_PREFIX)/$@:latest -f $< $@
	$(OCI_CMD) run -it --entrypoint /versioninfo $(IMAGE_PREFIX)/$@:latest
	$(OCI_CMD) tag $(IMAGE_PREFIX)/$@:latest $(IMAGE_PREFIX)/$@:$(shell $(OCI_CMD) run -it --entrypoint /versioninfo $(IMAGE_PREFIX)/$@:latest)