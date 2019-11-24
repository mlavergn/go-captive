###############################################
#
# Makefile
#
###############################################

.DEFAULT_GOAL := run

.PHONY: test

VERSION := 1.0.2

#
# Env
#
# ROUTER = 172.16.1.1
ROUTER = 192.168.2.1

#
# ARM router settings
#
GOOS = linux
GOARCH = mipsle
GOMIPS = softfloat

GOPATH = "${HOME}/go-captive"
	
lint:
	~/go/bin/golint src

build: clean
	GOPATH=${GOPATH} go build -o captive

run: build
	./captive

dist: clean
	GOPATH=${GOPATH} GOARCH=${GOARCH} GOOS=${GOOS} GOMIPS=${GOMIPS} go build --ldflags "-s -w" -o captive

deploy: dist
	-ssh root@${ROUTER} "mount /dev/sda1 /mnt"
	-ssh root@${ROUTER} "rm -f /mnt/root/captive"
	scp captive root@${ROUTER}:/mnt/root/captive

ssh:
	ssh root@${ROUTER}

clean:
	-rm -f captive
