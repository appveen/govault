##Golang-Vault

- A simple password protected key value truststore for golang applications.

- Pre-requisites :
    [1] go get -u github.com/boltdb/bolt
    [2] go get -u github.com/gorilla/mux

- Generate Executable Command
    go build cmd/main.go

- Generate Vault Using CLI
    main.exe -fileName=sample.vault -content="{\"storePassword\" : \"sample\" , \"data\" : { \"key1\" : \"value1\" , \"key2\" : \"value2\"}}"