package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

func main() {
	plaintext := []byte("Hello World!")
	password := []byte("helloworld")
	cipher := "aes192" // aes128 and aes256 work fine.
	packetConfig := &packet.Config{
		DefaultCipher: cipher2CipherFunc(cipher),
	}
	encrypted, err := Encrypt(plaintext, password, packetConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println("Encrypted:", string(encrypted))
	decrypted, err := Decrypt(encrypted, password, packetConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decrypted:", string(decrypted))
}

func cipher2CipherFunc(cipher string) packet.CipherFunction {
	switch cipher {
	default:
		if cipher != "" {
			fmt.Println("Invalid cipher, using default")
		}
		fallthrough
	case "aes256":
		return packet.CipherAES256
	case "aes192":
		return packet.CipherAES192
	case "aes128":
		return packet.CipherAES128
	}
}

func Encrypt(plaintext []byte, password []byte, packetConfig *packet.Config) (ciphertext []byte, err error) {
	encbuf := bytes.NewBuffer(nil)

	w, err := armor.Encode(encbuf, "PGP MESSAGE", nil)
	if err != nil {
		return
	}
	defer w.Close()

	pt, err := openpgp.SymmetricallyEncrypt(w, password, nil, packetConfig)
	if err != nil {
		return
	}
	defer pt.Close()

	_, err = pt.Write(plaintext)
	if err != nil {
		return
	}

	// Close writers to force-flush their buffer
	pt.Close()
	w.Close()
	ciphertext = encbuf.Bytes()

	return
}

func Decrypt(ciphertext []byte, password []byte, packetConfig *packet.Config) (plaintext []byte, err error) {
	decbuf := bytes.NewBuffer(ciphertext)

	armorBlock, err := armor.Decode(decbuf)
	if err != nil {
		return
	}

	failed := false
	prompt := func(keys []openpgp.Key, symmetric bool) ([]byte, error) {
		// If the given passphrase isn't correct, the function will be called again, forever.
		// This method will fail fast.
		// Ref: https://godoc.org/golang.org/x/crypto/openpgp#PromptFunction
		if failed {
			return nil, errors.New("decryption failed")
		}
		failed = true
		return password, nil
	}

	md, err := openpgp.ReadMessage(armorBlock.Body, nil, prompt, packetConfig)
	if err != nil {
		return
	}

	plaintext, err = ioutil.ReadAll(md.UnverifiedBody)
	if err != nil {
		return
	}

	return
}
