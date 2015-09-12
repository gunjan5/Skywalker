FROM golang:1.5.1

MAINTAINER Gunjan Patel <github.com/gunjan5>

RUN mkdir /app
ADD . /app/

EXPOSE 8080
EXPOSE 8081

WORKDIR /app

RUN ls -la

