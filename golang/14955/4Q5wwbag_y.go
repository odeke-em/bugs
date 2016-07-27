package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"log"
	"math/big"
)

func makeIssuer(name string, key crypto.Signer) (*x509.Certificate, error) {
	template := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: name,
		},
		IsCA: true,
		BasicConstraintsValid: true,
	}
	der, err := x509.CreateCertificate(rand.Reader, template, template, key.Public(), key)
	if err != nil {
		return nil, err
	}
	cert, err := x509.ParseCertificate(der)
	if err != nil {
		return nil, err
	}
	return cert, nil
}

func main() {
	issuerKey, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		log.Fatal(err)
	}
	eeKey, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		log.Fatal(err)
	}

	issuer1, err := makeIssuer("Issuer1", issuerKey)
	if err != nil {
		log.Fatal(err)
	}
	issuer2, err := makeIssuer("Issuer2", issuerKey)
	if err != nil {
		log.Fatal(err)
	}

	eeTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject: pkix.Name{
			CommonName: "EE",
		},
	}
	eeDER, err := x509.CreateCertificate(rand.Reader, eeTemplate, issuer1, eeKey.Public(), issuerKey)
	if err != nil {
		log.Fatal(err)
	}
	ee, err := x509.ParseCertificate(eeDER)
	if err != nil {
		log.Fatal(err)
	}

	err = ee.CheckSignatureFrom(issuer1)
	if err != nil {
		log.Fatalf("Problem checking signature on EE cert from issuer1: %s", err)
	}

	err = ee.CheckSignatureFrom(issuer2)
	if err == nil {
		log.Fatalf("Successfully verified signature on EE cert from issuer2, expected failure.")
	}
}
