.PHONY=default

main: main.go athena.go
	GOOS=linux go build -o main

default: main

main.zip: main
	zip main.zip main

