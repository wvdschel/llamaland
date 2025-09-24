.PHONY: containers publish-containers

containers:
	$(MAKE) -C containers/ images

publish-containers:
	$(MAKE) -C containers/ publish

test:
	cd llamaland; cargo test --all-features