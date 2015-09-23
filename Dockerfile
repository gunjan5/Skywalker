FROM docker:dind

MAINTAINER Gunjan Patel <github.com/gunjan5>

#install Golang and set ENV for Go
RUN apk update && apk add curl git mercurial bzr go && rm -rf /var/cache/apk/*

ENV GOROOT /usr/lib/go
ENV GOPATH /gopath
ENV GOBIN /gopath/bin
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin


#RUN mkdir /app
#ADD . /app/

EXPOSE 8080
EXPOSE 8081

#WORKDIR /app

RUN ls -la

CMD []


