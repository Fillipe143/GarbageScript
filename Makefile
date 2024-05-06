.SILENT:

build:
	go build -o garbage main.go

run: build  
	./garbage
