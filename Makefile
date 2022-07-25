PHONY: tidy start server

tidy:
	go mod tidy

app:
	rm -f server
	go build  -o server cmd/main.go


start:
	./server

rebuild:
	make app 
	make start

