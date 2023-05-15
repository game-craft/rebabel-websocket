FROM golang:1.18-alpine

ENV TZ /usr/share/zoneinfo/Asia/Tokyo

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

ENV GO111MODULE=on

COPY . .
EXPOSE 8001

RUN go mod init docker-echo-template && go mod tidy
RUN apk upgrade --update && apk add bash && apk --no-cache add git
RUN go install github.com/cosmtrek/air@latest
RUN go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go get -u github.com/labstack/echo/v4@latest
RUN go get -u github.com/dgrijalva/jwt-go@latest