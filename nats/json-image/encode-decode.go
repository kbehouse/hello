package main

import (
	"bufio"
	"encoding/base64"
	"io/ioutil"
	"os"
)

func encodeFromFile(filePath string) string {
	// Open file on disk.
	f, _ := os.Open(filePath)

	// Read entire JPG into byte slice.
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	return encoded
}

// refer: https://stackoverflow.com/questions/43212213/base64-string-decode-and-save-as-file
func decode2File(b64 string, filePath string) {
	dec, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}
}

// func main() {
// 	b64 := encode()
// 	fmt.Println("ENCODED: " + b64)
// 	fmt.Printf("ENCODED size: %v\n", len(b64))

// 	decode2file(b64, "./example-decode.jpg")
// }
