# Golang-Vault
A simple password protected key-value trust store for go-lang applications.

# Pre-requisites :
* github.com/boltdb/bolt
* github.com/gorilla/mux

# Install dependencies
```
go get -u github.com/boltdb/bolt
go get -u github.com/gorilla/mux
```

# Generate Executable
`go build -o vault cmd/main.go`

# Usage
```
vault
  -p <Required. Password that has to be set for the vault>
  -i <Optional. Input JSON file with data that will be added to the vault>
  -o <Optional. Output filename. Defaults to db.valut>

```