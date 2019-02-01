package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	".."
)

//main.exe -fileName=sample.vault -content="{\"storePassword\" : \"sample\" , \"data\" : { \"key1\" : \"value1\" , \"key2\" : \"value2\"}}"

func main() {
	arguments := os.Args[1:]
	fileName := "db.vault"
	content := []byte("{}")
	password := ""
	for a := 0; a < len(arguments); a = a + 2 {
		switch arguments[a] {
		case "-o":
			if len(arguments)-1 < a+1 {
				log.Fatal("-o argument requires output file name parameter")
			}
			fileName = arguments[a+1]
			break
		case "-i":
			if len(arguments)-1 < a+1 {
				log.Fatal("-i argument requires input json file name parameter")
			}
			inputFile, err := os.OpenFile(arguments[a+1], os.O_RDONLY, 0600)
			if err != nil {
				log.Fatal(err)
			}
			content, err = ioutil.ReadAll(inputFile)
			if err != nil {
				log.Fatal(err)
			}
			break
		case "-p":
			if len(arguments)-1 < a+1 {
				log.Fatal("Invalid Arguments")
			} else if arguments[a+1] == "" {
				log.Fatal("Password string not provided")
			}
			password = arguments[a+1]
			break
		default:
			log.Fatal("Invalid Argument = ", arguments[a])
		}
	}
	fmt.Println("FileName = ", fileName)
	fmt.Println("Content = ", content)
	fmt.Println("password = ", password)
	vault.DownloadTruststoreCLI(fileName, []byte(content), password)
}
