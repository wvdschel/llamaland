.PHONY: containers publish-containers

test:
	go test -v -count=1 ./...

containers:
	$(MAKE) -C containers/ images

publish-containers:
	$(MAKE) -C containers/ publish