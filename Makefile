unit-test-coverage:
	ginkgo -r -v -cover && go tool cover -html=coverprofile.out
unit-test:
	ginkgo test ./...
generate-mock:
	go generate -v ./...