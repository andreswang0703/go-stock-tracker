.PHONY: build zip clean
build:
	GOOS=linux GOARCH=arm64 go build -o bootstrap main.go

zip: build
	zip deployment.zip bootstrap

clean:
	rm -f bootstrap deployment.zip