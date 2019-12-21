#!/usr/bin/make

day := '05'

testDay: fmt
	go test -count=1 advent/2019/day${day} && \
	go test -count=1 advent/2019/intcode && \
	go test -count=1 -run 'TestFixture2019/TestDay${day}' advent/2019

test2019: fmt
	go test -count=1 advent/2019/...

test2018: fmt
	go test -count=1 advent/2018/...

test2017: fmt
	go test -count=1 advent/2017/...

test2016: fmt
	go test -count=1 advent/2016/...

test2015: fmt
	go test -count=1 advent/2015/...

test: test2015 test2016 test2017 test2018 test2019

fmt:
	go fmt ./...

.PHONY: fmt test test2015 test2016 test2017 test2018 test2019 testDay