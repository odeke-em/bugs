package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	tlsConfig, err := tlsConfigWithHTTP2()
	if err != nil {
		log.Fatal(err)
	}
	srv := http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(15 * time.Second)
			fmt.Fprintf(w, fmt.Sprintf("PID: %d\n", os.Getpid()))
		}),
		TLSConfig: tlsConfig,
	}
	term := make(chan os.Signal, 1)
	done := make(chan struct{})
	signal.Notify(term, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-term
		log.Println("Shutting down http server...")
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Println(err)
		}
		close(done)
	}()
	errCh := make(chan error)
	go func() {
		if err := srv.ListenAndServeTLS("", ""); err != nil {
			if err != http.ErrServerClosed {
				errCh <- err
			}
		}
	}()
	select {
	case err = <-errCh:
		log.Fatal(err)
	case <-done:
		log.Println("done")
	}
}

func tlsConfig() (*tls.Config, error) {
	cert, err := createSelfSignedCertificate()
	if err != nil {
		return nil, fmt.Errorf("failed to create self signed certificate: %s", err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{*cert},
		NextProtos: []string{
			"http/1.1",
		},
	}, nil
}

func tlsConfigWithHTTP2() (*tls.Config, error) {
	cert, err := createSelfSignedCertificate()
	if err != nil {
		return nil, fmt.Errorf("failed to create self signed certificate: %s", err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{*cert},
		NextProtos: []string{
			"h2",
		},
	}, nil
}

func createSelfSignedCertificate() (*tls.Certificate, error) {
	rsaKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %v", err)
	}
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, fmt.Errorf("failed to generate certificate serial number: %v", err)
	}
	hostname, err := os.Hostname()
	if err != nil {
		return nil, fmt.Errorf("failed to get hostname: %v", err)
	}
	tmpl := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName:   hostname,
			Country:      []string{"FR"},
			Organization: []string{"ORG"},
			Locality:     []string{"PAR"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}
	derBytes, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, &rsaKey.PublicKey, rsaKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create certificate: %v", err)
	}
	var (
		certPEMBlock bytes.Buffer
		keyPEMBlock  bytes.Buffer
	)
	if err = pem.Encode(&certPEMBlock, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		return nil, fmt.Errorf("failed to encode certificate DER bytes as PEM: %v", err)
	}
	if err = pem.Encode(&keyPEMBlock, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(rsaKey),
	}); err != nil {
		return nil, fmt.Errorf("failed to encode key DER bytes as PEM: %v", err)
	}
	crt, err := tls.X509KeyPair(certPEMBlock.Bytes(), keyPEMBlock.Bytes())
	if err != nil {
		return nil, fmt.Errorf("failed to generate tls certificate from PEM Block bytes: %v", err)
	}
	return &crt, nil
}
