package cmd

import (

	"io/ioutil"
	"log"
)

func readfile(fileName string) (string){
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)
	return text
}
