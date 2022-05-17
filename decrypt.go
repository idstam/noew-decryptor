package main

import (
	"io/ioutil"
	"os"
)

func main() {
	encryptedData, _ := ioutil.ReadFile("Noew.SAM")
	swapBuffer := []byte{}
	outBuffer := []byte{}
	for _, b := range encryptedData {
		swapBuffer = append(swapBuffer, b)
		if len(swapBuffer) == 4 {
			for i := 3; i >= 0; i-- {
				outBuffer = append(outBuffer, swapBuffer[i])
			}
			swapBuffer = []byte{}
		}
	}
	for i := len(swapBuffer) - 1; i >= 0; i-- {
		outBuffer = append(outBuffer, swapBuffer[i])
	}
	os.WriteFile("out.ex_", outBuffer, 0644)
}
