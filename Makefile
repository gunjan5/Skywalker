build:
	go build -a -v main.go server.go

start:
	go build -v main.go server.go
	./main

get:
	go get github.com/gorilla/mux
	go get golang.org/x/net/context
	go get github.com/go-kit/kit/transport/http
	go get github.com/gunjan5/Skywalker/services/composeService
	go get github.com/gunjan5/Skywalker/services/yellingService
	go get github.com/go-kit/kit/endpoint
	go get github.com/docker/libcompose/docker
	go get github.com/docker/libcompose/project
	



