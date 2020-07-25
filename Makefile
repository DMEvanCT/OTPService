.PHONY: build clean deploy

build:
	#dep ensure -v
	env  go build -ldflags="-s -w" -o bin/otpservice  cmd/main/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

