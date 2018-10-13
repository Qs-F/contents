test: build
	go test -v

build:
	goimports.exe -w .
