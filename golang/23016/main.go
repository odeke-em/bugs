package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	var cert = `-----BEGIN CERTIFICATE-----
MIIGPjCCBSigAwIBAgIFBDFmk+0wCwYJKoZIhvcNAQELMFYxCzAJBgNVBAYTAlVT
MRYwFAYDVQQKEw1Nb3RoZXIgTmF0dXJlMRMwEQYDVQQLEwpFdmVyeXRoaW5nMRgw
FgYDVQQDEw9OYW1lIGNvbnN0cmFpbnQxADAiGA8yMDU1MTIwMTA2MDcwOFoYDzIw
NTYxMDA1MjM1MjU4WjCBmzELMAkGA1UEBhMCVVMxGDAWBgNVBAoTD0V4dHJlbWUg
RGlzY29yZDEOMAwGA1UECxMFQ2hhb3MxFDASBgNVBAcTC1RhbGxhaGFzc2VlMQsw
CQYDVQQIEwJGTDEcMBoGA1UECRMTMzIxMCBIb2xseSBNaWxsIFJ1bjEOMAwGA1UE
ERMFMzAwNjIxDzANBgNVBAMTBmdvdi51czEAMIIBIjANBgkqhkiG9w0BAQEFAAOC
AQ8AMIIBCgKCAQEA2lVo8m23yLGU5GBObxhGAMYJNfMxLXwmGaYDQGN/Zr82ZK+Y
5oYwDd8OpZmLxU1aRfWoLSs6JQ/LbftdvT5nns+xFD/ZpboeghlZvpA8espM5zDR
5i4Qs8pmfqbrARdKV1khTDfgFw6aGTmatdHcwfJ+U6jOaXPTI8S4olN4LsyV0sCa
JT/c1SBXOdfQjLGe2cxwIyBzvwxj07guuYzq3Kj72d9sKyWPAdTdBngtKH2jvVq9
XncFo+juD+nhQOsRPCqAaE5cdsdSF+YTuJF/ZrAe5ohY0UiyTg9FrxVe8hf3PB3r
h9yjHuC7J0D0M39+dmhmyfYrAzDp9NnXsWXVvQIDAQABo4ICyzCCAscwDgYDVR0P
AQH/BAQDAgCkMB0GA1UdJQQWMBQGCCsGAQUFBwMCBggrBgEFBQcDATAPBgNVHRMB
Af8EBTADAQH/MA4GA1UdIwQHMAWAAwECAzBiBggrBgEFBQcBAQRWMFQwIQYIKwYB
BQUHMAGGFWh0dHA6Ly90aGVjYS5uZXQvb2NzcDAvBggrBgEFBQcwAoYjaHR0cDov
L3RoZWNhLm5ldC90b3RhbGx5dGhlY2VydC5jcnQwOAYDVR0RBDEwL4YgaXJyZWxl
dmFudGluZm8vL3VzZXJAMTkyLjE2OC4xLjGCC3d3dy5kbnMuY29tMAsGA1UdEgQE
MACCADAbBgNVHSAEFDASMAgGBmeBDAECAjAGBgQqAwQFMIIBqwYDVR0eBIIBojCC
AZ6ggc4wE4ERZ29vZF9lbWFpbEBnZy5jb20wCYEHTHVsTWFpbDAPgg1wZXJtaXR0
ZWQuY29tMIGOpIGLMIGIMQswCQYDVQQGEwJVUzENMAsGA1UEChMEVUlVQzEMMAoG
A1UECxMDRUNFMRIwEAYDVQQHEwlDaGFtcGFpZ24xCzAJBgNVBAgTAklMMRYwFAYD
VQQJEw02MDEgV3JpZ2h0IFN0MQ4wDAYDVQQREwU2MTgyMDERMA8GA1UEAxMIdWl1
Yy5uZXQxADAKhwhKfeBI//8AAKGByjASgRBiYWRfZW1haWxAZ2cuY29tMAmBB0x1
bE1haWwwDIIKYmFubmVkLmNvbTCBjqSBizCBiDELMAkGA1UEBhMCVVMxDjAMBgNV
BAoTBVVtaWNoMQswCQYDVQQLEwJDUzESMBAGA1UEBxMJQW5uIEFyYm9yMQswCQYD
VQQIEwJNSTEVMBMGA1UECRMMNTAwIFN0YXRlIFN0MQ4wDAYDVQQREwU0ODEwOTES
MBAGA1UEAxMJdW1pY2gubmV0MQAwCocIwKgBAf//AAAwCwYJKoZIhvcNAQELA4IB
AQDfA7wx83CZo+SG+revr32LacHA/EQFHGI4lyHQWbBdVsDCb6N7w9fNXCC3+MRo
cWYNNWctlvyEZfDoU/mshh2vpAQcL56toMJYCX8+6GeAJCGIS4yD7ks+j7GY+kzR
OnW1A3irglDycF3blQoaGJsLp94VjyOzESh0+nNwypfF1lGZFzjBpXGsL04hKzAR
7CAADZ7qSHvB39o8VdEX0toWk791G+yys7ugx/pAxWIpmYU68lco9S860KxA66ln
7s3g2KgdqC4kU9auUM3tvHFqdnXfbBYSYC1qvihKwmeFbCmAgsKH1rm4x4OuEohb
ewclEGevcJayEHyQjYNwtXfg
-----END CERTIFICATE-----`

	var pemBlock, _ = pem.Decode([]byte(cert))
	var parsedCert, err = x509.ParseCertificate(pemBlock.Bytes)
	fmt.Println("Cert", parsedCert)
	fmt.Println("err", err)
}
