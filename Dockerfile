FROM golang:1.10.1
MAINTAINER Jun Guan garfilone@gmail.com

RUN go get github.com/astaxie/beego && go get github.com/beego/bee