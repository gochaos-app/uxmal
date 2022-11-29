dev:
	echo "Compiling..."
	go build -o uxmal . 
prod:
	echo "Compiling..."
	go build -ldflags="-s -w" -o uxmal .

compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=amd64 go build -ldflags="-s -w" -o bin/uxmal-freebsd-amd64 .
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/uxmal-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o bin/uxmal-linux-arm64 .
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o bin/uxmal-darwin .
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o bin/uxmal-darwin-m1 .
