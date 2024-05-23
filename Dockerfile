FROM golang:latest as builder

WORKDIR $GOPATH/src

COPY . .

ENV GO111MODULE=on
ENV GOSUMDB=off

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    GONOPROXY="" GOSUMDB="off" \
    go build -C cmd/gogen \
    -o /gogen

FROM alpine:latest

COPY --from=builder /go /go
COPY --from=builder /usr/local/go /usr/local/go
COPY --from=builder /gogen /gogen
COPY --from=builder /go/src/generate /
COPY --from=registry-new.diasoft.ru/dvt/baseimages/go-swag:24050601 /bin/swag /

ENV GOPATH /go
ENV PATH $PATH:/go/bin:/usr/local/go/bin
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
ENV GONOPROXY ""
ENV GOPROXY http://10.20.99.32:3000
ENV GOSUMDB off

RUN chmod -R 777 /root/
RUN chmod -R 777 /go/
RUN chmod -R 777 /usr/local/go/
RUN chmod u+x /generate