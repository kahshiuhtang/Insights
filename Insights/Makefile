BINARY_NAME=insights
.DEFAULT_GOAL :=  run

build:
	GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin cmd/Insights/main.go
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux cmd/Insights/main.go
	GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-windows cmd/Insights/main.go

run: build
	./bin/${BINARY_NAME}-linux


clean:
	go clean
	rm ./bin/${BINARY_NAME}-darwin
	rm ./bin/${BINARY_NAME}-linux
	rm ./bin/${BINARY_NAME}-windows

test:

vet:
	go vet

lint: 
	

