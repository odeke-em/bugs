package main

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
)

const tlsCrt = `-----BEGIN CERTIFICATE-----
MIIBgzCCASmgAwIBAgIQL7MpDWvb8tnkLd7ydYboUjAKBggqhkjOPQQDAjAqMRQw
EgYDVQQKEwtDaGFpbiwgSW5jLjESMBAGA1UEAxMJbG9jYWxob3N0MB4XDTE3MDUx
OTIyMzEzMloXDTE4MDUxOTIyMzEzMlowKjEUMBIGA1UEChMLQ2hhaW4sIEluYy4x
EjAQBgNVBAMTCWxvY2FsaG9zdDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABGOC
qYVtlTOSffS8L7IzdmNYc0zI1T6xqJ5ld+OPHV5hKude/ujHC98OLfwSY+GyPEPD
ZtNVVUVGX6OD56+qeuWjMTAvMA4GA1UdDwEB/wQEAwIFoDAPBgNVHSUECDAGBgRV
HSUAMAwGA1UdEwEB/wQCMAAwCgYIKoZIzj0EAwIDSAAwRQIgQc8L/zq+bwDWx3ZL
+K1Bd636kp0gSbV3mWqUDyPg5EkCIQCVHbGBXj9cXwIf+I54SAD8gtjEE4oGLJIs
w3TNdjbg/A==
-----END CERTIFICATE-----`
const tlsKey = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIIk0ZuFh2cGXJKxaZv3mot9OSv3dKYOq/qqBOq9JBtK8oAoGCCqGSM49
AwEHoUQDQgAEY4KphW2VM5J99LwvsjN2Y1hzTMjVPrGonmV3448dXmEq517+6McL
3w4t/BJj4bI8Q8Nm01VVRUZfo4Pnr6p65Q==
-----END EC PRIVATE KEY-----`

func main() {
	config := &tls.Config{
		Certificates: make([]tls.Certificate, 1),
		RootCAs:      x509.NewCertPool(),
	}

	var err error
	config.Certificates[0], err = tls.X509KeyPair([]byte(tlsCrt), []byte(tlsKey))
	if err != nil {
		panic(err)
	}
	x509Cert, err := x509.ParseCertificate(config.Certificates[0].Certificate[0])
	if err != nil {
		panic(err)
	}
	config.RootCAs.AddCert(x509Cert)
	// config.NextProtos = []string{"h1"}

	server := &http.Server{
		Addr:      ":8080",
		TLSConfig: config,
		Handler:   http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
	}
	go func() {
		err := server.ListenAndServeTLS("", "")
		if err != nil {
			panic(err)
		}
	}()

	client := &http.Client{Transport: &http.Transport{TLSClientConfig: config}}
	resp, err := client.Post("https://localhost:8080/", "", nil)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
