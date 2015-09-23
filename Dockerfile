FROM docker:dind

MAINTAINER Gunjan Patel <github.com/gunjan5>

#install Golang and set ENV for Go
RUN apk update && apk add curl rsync bash git mercurial bzr go && rm -rf /var/cache/apk/*

ENV GOROOT /usr/lib/go
ENV GOPATH /gopath
ENV GOBIN /gopath/bin
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin

RUN git clone https://github.com/GoogleCloudPlatform/kubernetes
RUN cd kubernetes/
RUN KUBE_STATIC_OVERRIDES=kubectl hack/build-go.sh cmd/kubectl
RUN cd /kubernetes/_output/local/bin/linux/amd64/
RUN mv kubectl /usr/local/bin/
RUN cd

#RUN mkdir /app
#ADD . /app/

EXPOSE 8080
EXPOSE 8081

#WORKDIR /app

RUN ls -la
RUN kubectl

CMD []


