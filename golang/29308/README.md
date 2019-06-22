## 29308

Investigation and repros for https://golang.org/issues/29308

### Remedies

Take a look at https://github.com/golang/go/issues/29308#issuecomment-504617974

#### Before proposed patch

```shell
server: 2019/06/21 19:12:54 Listening at "[::]:49373"
client: 2019/06/21 19:12:55 URL: http://[::]:49373/hello
server: 2019/06/21 19:12:55 Request from: [::1]:49376

client: 2019/06/21 19:12:55 Blob: Server responding ASAP
client: 2019/06/21 19:12:55 Pausing for 1.5s

client: 2019/06/21 19:12:57 URL: http://[::]:49373/hello
server: 2019/06/21 19:12:57 Request from: [::1]:49376

server: 2019/06/21 19:12:57 Now going to lag for 3s
server: 2019/06/21 19:12:57 Pausing for 2s before reviving client with pid: 21447
server: 2019/06/21 19:13:27 Request from: [::1]:49377

server: 2019/06/21 19:13:27 Now going to lag for 3s
client: 2019/06/21 19:13:33 Failed to make request("http://[::]:49373/hello"): Get http://[::]:49373/hello: read tcp [::1]:49377->[::1]:49373: read: connection reset by peer
```

#### After proposed patch

```shell
server: 2019/06/21 19:11:00 Listening at "[::]:49346"
client: 2019/06/21 19:11:01 URL: http://[::]:49346/hello
server: 2019/06/21 19:11:01 Request from: [::1]:49349

client: 2019/06/21 19:11:01 Blob: Server responding ASAP
client: 2019/06/21 19:11:01 Pausing for 1.5s

client: 2019/06/21 19:11:02 URL: http://[::]:49346/hello
server: 2019/06/21 19:11:02 Request from: [::1]:49349

server: 2019/06/21 19:11:02 Now going to lag for 3s
server: 2019/06/21 19:11:02 Pausing for 2s before reviving client with pid: 21387
server: 2019/06/21 19:11:32 Request from: [::1]:49360

server: 2019/06/21 19:11:32 Now going to lag for 3s
server: 2019/06/21 19:11:39 Request from: [::1]:49362
```
