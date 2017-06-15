# Investigating issue 20689

## Usage
* Server
```shell
$ go run server.go
```

* Client
The client can be used to invoke a specific URL or the
default URL of the server running as mentioned above.
Make sure to set debug flags `GODEBUG=http2debug=2`
```shell
$ GODEBUG=http2debug=2 go run client.go
```
or for example a custom URL:
```shell
$ GODEBUG=http2debug=2 go run client.go --url  https://storage.googleapis.com/veener-jba/metadata-test
```
