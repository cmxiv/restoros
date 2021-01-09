PKGS=$(shell go list ./...)

test:
	@go test $(PKGS)
