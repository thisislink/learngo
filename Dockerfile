FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/thisislink/learngo
RUN cd /build && git clone https://github.com/thisislink/learngo.git

RUN cd /build/learngo/src && go build

EXPOSE 9000

ENTRYPOINT [ "/build/learngo/src/main" ]