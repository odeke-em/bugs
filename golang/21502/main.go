package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

const ecdsaCert = `-----BEGIN CERTIFICATE-----
MIIDnjCCA0SgAwIBAgIJAMcMB20DTOtBMAoGCCqGSM49BAMCMIGRMQswCQYDVQQG
EwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNj
bzEOMAwGA1UECgwFQ2hhaW4xFDASBgNVBAsMC2RldmVsb3BtZW50MS8wLQYDVQQD
DCZEZXZlbG9wbWVudCBSb290IENlcnRpZmljYXRlIEF1dGhvcml0eTAeFw0xNzA4
MTcxODMyNDdaFw00NTAxMDIxODMyNDdaMIGRMQswCQYDVQQGEwJVUzETMBEGA1UE
CAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNjbzEOMAwGA1UECgwF
Q2hhaW4xFDASBgNVBAsMC2RldmVsb3BtZW50MS8wLQYDVQQDDCZEZXZlbG9wbWVu
dCBSb290IENlcnRpZmljYXRlIEF1dGhvcml0eTCCAUswggEDBgcqhkjOPQIBMIH3
AgEBMCwGByqGSM49AQECIQD/////AAAAAQAAAAAAAAAAAAAAAP//////////////
/zBbBCD/////AAAAAQAAAAAAAAAAAAAAAP///////////////AQgWsY12Ko6k+ez
671VdpiGvGUdBrDMU7D2O848PifSYEsDFQDEnTYIhucEk2pmeOETnSa3gZ9+kARB
BGsX0fLhLEJH+Lzm5WOkQPJ3A32BLeszoPShOUXYmMKWT+NC4v4af5uO5+tKfA+e
FivOM1drMV7Oy7ZAaDe/UfUCIQD/////AAAAAP//////////vOb6racXnoTzucrC
/GMlUQIBAQNCAAQ3NftnV/SX+IBClVPjcr4tnLnKN6roTA137uYBCI25VeyjbtF1
I0AY3oWNBk03zj7unu50auT4BzfAsd6uWEg4o4GOMIGLMDsGA1UdEQQ0MDKCD2No
YWluLmxvY2FsaG9zdIIUdGVzdC5jaGFpbi5sb2NhbGhvc3SCCWxvY2FsaG9zdDAd
BgNVHQ4EFgQUdTZvsp3zDb8adDH4t2Nb29pR0DUwHwYDVR0jBBgwFoAUdTZvsp3z
Db8adDH4t2Nb29pR0DUwDAYDVR0TBAUwAwEB/zAKBggqhkjOPQQDAgNIADBFAiBW
L/d1DNYBgAWB1nYRIp4eusEyWA/ITMm9+m/DbdDokgIhAIDwR/Nkr7hdyZANUd6O
jJJJY8A9ay2Odxpocvj1Bog4
-----END CERTIFICATE-----`

func main() {
	// Reading CA cert into a cert pool
	pemCert := []byte(ecdsaCert)
	certPool := x509.NewCertPool()
	ok := certPool.AppendCertsFromPEM(pemCert)
	if !ok {
		fmt.Printf("Error reading devCA cert\n")
	} else {
		fmt.Printf("Successful decode, exiting...")
		return
	}
	
	// Copied from crypto/x509 package to reveal error message
	for len(pemCert) > 0 {
		var block *pem.Block
		block, pemCert = pem.Decode(pemCert)
		if block == nil {
			break
		}
		if block.Type != "CERTIFICATE" || len(block.Headers) != 0 {
			fmt.Printf("block.Type != CERTIFICATE")
			continue
		}
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			fmt.Printf("Error parsing certificate: \n%s", err)
			continue
		}
		fmt.Printf("Success!\nCertificate: %+v\n", cert)
	}
}
