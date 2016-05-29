package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/openpgp"
)

func encrypted() (io.ReadCloser, error) {
	return os.Open("encrypted.gpg")
}

func main() {
	encrypted, err := encrypted()
	if err != nil {
		log.Fatal(err)
	}
	defer encrypted.Close()

	times := 0
	readPassword := func(keys []openpgp.Key, symmetric bool) ([]byte, error) {
		if times == 0 {
			times += 1
			return []byte("invalid password"), nil
		}
		return []byte("golang"), nil
	}
	md, err := openpgp.ReadMessage(encrypted, nil, readPassword, nil)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(md.UnverifiedBody)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytes))
}
