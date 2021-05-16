
.PHONY: test
test: clean
	@go test -coverprofile=coverage.out ./...

test/coverage: test
	@go tool cover -html=coverage.out

clean:
	@rm -f coverage.out restoros

cleanall: clean
	@rm -rf $(HOME)/.restoros

run: clean
	go build
	./restoros $(RESTOROS_ARGS)