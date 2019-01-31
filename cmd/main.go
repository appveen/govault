package main

import (
	"fmt"
	"os"
	"strings"

	".."
)

//main.exe -fileName=sample.vault -content="{\"storePassword\" : \"sample\" , \"data\" : { \"key1\" : \"value1\" , \"key2\" : \"value2\"}}"

func main() {
	arguments := os.Args[1:]
	fmt.Println(arguments)
	fileName := strings.Split(arguments[0], "-fileName=")[1]
	content := strings.Split(arguments[1], "-content=")[1]
	fmt.Println("FileName = ", fileName)
	fmt.Println("Content = ", content)
	vault.DownloadTruststoreCLI(fileName, []byte(content))
}
