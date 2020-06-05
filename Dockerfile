FROM golang:1.14-alpine as agents

WORKDIR /app

RUN apk add git

# Fetching Dependencies
RUN go get -u github.com/boltdb/bolt
RUN go get -u github.com/gorilla/mux

COPY . .

# Building Executables
# Mac Build
RUN env GOOS=darwin GOARCH=amd64 go build -o exec/vault-darwin-amd64 cmd/main.go || true
# Linux Build
RUN env GOOS=linux GOARCH=386 go build -o exec/vault-linux-386 cmd/main.go
RUN env env GOOS=linux GOARCH=amd64 go build -o exec/vault-linux-amd64 cmd/main.go || true
# Windows Build
RUN env GOOS=windows GOARCH=386 go build -o exec/vault-windows-386.exe cmd/main.go
RUN env GOOS=windows GOARCH=amd64 go build -o exec/vault-windows-amd64.exe cmd/main.go



FROM scratch

WORKDIR /app

COPY LICENSE .
COPY --from=agents /app/exec ./exec
