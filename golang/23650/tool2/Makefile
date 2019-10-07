GOPATH = $(PWD)

all: encrypt create_update md5_res

encrypt:
	cd tools && go build encrypt_res

create_update:
	cd tools && go build create_update_zip

md5_res:
	cd tools && go build md5_res

compress:
	cd tools && go build compress_png

test:
	go build test

clean:
	rm -r 

