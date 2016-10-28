package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"
)

func main() {
	// Verifying with a custom list of root certificates.

	const rootPEM = `
-----BEGIN CERTIFICATE-----
MIIDlTCCAn+gAwIBAgIIVvpPzLyqk+0wCwYJKoZIhvcNAQELMGoxaDAJBgNVBAYT
AlVTMBQGA1UECAwNTWFzc2FjaHVzZXR0czAOBgNVBAcMB05ld2J1cnkwFgYDVQQK
DA9CYWQgVExTIExpbWl0ZWQwHQYDVQQDDBZCYWQgVExTIExpbWl0ZWQgUlNBIENB
MB4XDTE2MDEwMTAwMDAwMFoXDTI2MDEwMTAwMDAwMFowajFoMAkGA1UEBhMCVVMw
FAYDVQQIDA1NYXNzYWNodXNldHRzMA4GA1UEBwwHTmV3YnVyeTAWBgNVBAoMD0Jh
ZCBUTFMgTGltaXRlZDAdBgNVBAMMFkJhZCBUTFMgTGltaXRlZCBSU0EgQ0EwggEi
MA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDSHu3OR1RS0D2xLKGK2Ts5eLoO
/P+IXst5WPdaD9UwGI8edfAy3U8wcMFDoXNhBQM+ZW69Z5uOZVxs704+j5cgCEAT
LbtyIrF2X8BixXFzrJFd+kpojURheyxML20GbZsznJgKzYvGqFqWa/1lYwy/v0SP
RNGPEkjFXb/tItDwrDxcuDzY6zjNlW5MwqvS11P1H8eg0idUrANY2MzT8+oyH3Sn
JLCsmulnmj1b6IZZDN4i8rKXEbH14jIsANHIgTqvS+kJf3Z1PqHAOUqVGlO3SDZd
KIqZ8olS6ty9/pco6cxvX2Te9m1z5f1fSrdxAtx7lHM3pdvs9DhML+8FAewDAgMB
AAGjQzBBMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFGEbxkZbhgwiZRAMx7Vs
VCRXl/tkMA8GA1UdDwEB/wQFAwMHBgAwCwYJKoZIhvcNAQELA4IBAQBKv0TJoRhd
wg7dPOFDVuKaLtuVzXEeUWfsA86iW4wjXFO/npI+1exSBX92MhsWk5Gjn9dO/Hq4
EZ1pMJ8hFdrOXoEHlvhnZSavtoy25ZvEoxJ9XWYPqWCmwdfB3xhT4hoEaIlu5Azf
Fw/QV5oFV8SYgwClQ+fTStxdW7CBKEX55KPUn4FOOXV5TfbLOJj3w/1V2pBTKn2f
2safgWyIpNw7OyvYVICdW5/NvD+VTBp+4PfWkTfRD5LEAxqvaGXupBaI2qGYVibJ
WQ77yy6bOvcJh4heqtIJuYg5F3vhvSGo4i5Bkx+daRKFzFwsoiexgRNTdlPCEGsQ
15WBlk3X/9bt
-----END CERTIFICATE-----`

	const certPEM = `
-----BEGIN CERTIFICATE-----
MIIDzzCCArmgAwIBAgIIVvpPzDSsZqswCwYJKoZIhvcNAQELMGoxaDAJBgNVBAYT
AlVTMBQGA1UECAwNTWFzc2FjaHVzZXR0czAOBgNVBAcMB05ld2J1cnkwFgYDVQQK
DA9CYWQgVExTIExpbWl0ZWQwHQYDVQQDDBZCYWQgVExTIExpbWl0ZWQgUlNBIENB
MB4XDTE2MDEwMTAwMDAwMFoXDTE5MDEwMTAwMDAwMFowajFoMAkGA1UEBhMCVVMw
FAYDVQQIDA1NYXNzYWNodXNldHRzMA4GA1UEBwwHTmV3YnVyeTAWBgNVBAoMD0Jh
ZCBUTFMgTGltaXRlZDAdBgNVBAMMFmRvbWFpbi1tYXRjaC5iYWR0bHMuaW8wggEi
MA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCkpFshPkiPs33jreAgIoGfNXlP
qTexrpdS5TkepZOUALEtgd/FZHgA5UFrYQCtq3kN3u2+SDwMGvyU9mFQl2lp2rD8
7wW+42zAgchVNPqaDfSuDv6smP7hbN0DBBNXX+r7JQcgRxHECHqdQ4Jyzv8cD3O2
7eKyBVxnOpECT2YgogTRoJLq63l6hxx4smBYJsNpr4lkV32t1uNnZQU14rk9p5/F
EXv8cN+qDQ8s6k1A2U20kSo915CKY7JMVZSB7lVY5jtMWsPiQM1lUQK497R6rrZM
YDM4ydb18NTUrxRbtdQIpVB03jM9bwOId9J3Zc1dE4FmypbXFOFPfUoTRwT9AgMB
AAGjfTB7MB8GA1UdIwQYMBaAFGEbxkZbhgwiZRAMx7VsVCRXl/tkMAkGA1UdEwQC
MAAwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMB0GA1UdDgQWBBT5jkQE
Z8/qr1Ae18kStR+mUOoaOTAPBgNVHQ8BAf8EBQMDB6AAMAsGCSqGSIb3DQEBCwOC
AQEAX1GqUgXPYgckfnrh0R1z69GVvrl+sBf92mP6oA4Y0J7RQr1sIbz+9y21p7/X
c7gLKONX+xCqZTmgZs8DkYkYgWr4JOqF61zwFN/BHw2IhI7SGlcDqKMu9kRm7YkY
onaKX/1UzVlv92+HK5kwYhnEXGHY8MVL+/TivLuZnjXm5Fj56mIkLGD//IrIR/HD
tOPi88q5gahMgW9YqM6NZ2VNw4bamaFAq/lY3htY9Mw0oXYvqHEsvqAl+NjwWV7M
iVk60h/tU4E+qJ5Z56CrjNDEjECPGW7iYyx6wmzY2i+3bLbPVQNO/mhUHQYlEblR
6VAoNOlZme5zPJanghfiOT0D8A==
-----END CERTIFICATE-----`

	// First, create the set of root certificates. For this example we only
	// have one. It's also possible to omit this in order to use the
	// default root set of the current operating system.
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	if !ok {
		panic("failed to parse root certificate")
	}

	block, _ := pem.Decode([]byte(certPEM))
	if block == nil {
		panic("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic("failed to parse certificate: " + err.Error())
	}

	opts := x509.VerifyOptions{
		CurrentTime: time.Unix(1471879557, 0),
		DNSName:     "domain-match.badtls.io",
		Roots:       roots,
	}

	if _, err := cert.Verify(opts); err != nil {
		panic("failed to verify certificate: " + err.Error())
	}
	fmt.Println("success, certificate valid")
}
