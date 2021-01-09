PKGS=$(shell go list ./...)

test:
	@go test -coverprofile=coverage.out $(PKGS)
