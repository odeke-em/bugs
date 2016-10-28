package main

import (
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	key, err := ioutil.ReadFile(os.ExpandEnv("./bb.key"))
	if err != nil {
		log.Fatal(err)
	}
	_, err = ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("cannot parse ssh private key: %v", err)
	}
}
