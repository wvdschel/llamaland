.PHONY: containers publish-containers

containers:
	$(MAKE) -C containers/ images

publish-containers:
	$(MAKE) -C containers/ publish