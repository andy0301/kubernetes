# https://dl.google.com/go/go1.13.8.linux-amd64.tar.gz

FROM golang

WORKDIR /go/src
COPY index.html /go/src
COPY web_example.go /go/src

RUN cd /go/src && go build ./web_example.go

EXPOSE 8080
 # This is the first command 
ENTRYPOINT ["./web_example"]