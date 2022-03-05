compile:
	GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/ip-lookup src/iplookup/main.go

clean:
	rm ip-lookup.zip
	rm -rf bin