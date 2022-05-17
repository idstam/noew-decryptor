package main

import (
	"io/ioutil"
	"os"
)

func main() {

	encryptedData, _ := ioutil.ReadFile("file.txt")

	buffer := []byte{}
	outBuffer := []byte{}
	for _, b := range encryptedData {
		buffer = append(buffer, b)
		if len(buffer) == 4 {
			for i := 0; i < 4; i++ {
				outBuffer = append(outBuffer, buffer[i])
			}
			buffer = []byte{}
		}
	}
	for i := 0; i < len(buffer); i++ {
		outBuffer = append(outBuffer, buffer[i])
	}

	os.WriteFile("out.ex_", outBuffer, 0644)

}
