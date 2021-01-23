PKGS=$(shell go list ./...)

test:
	@go test -coverprofile=coverage.out $(PKGS)

clean:
	@rm -f coverage.out restoros