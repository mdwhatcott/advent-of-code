#!/usr/bin/make

testDay: fmt
	go test -v -count=1 advent/2019/day04
	go test -v -count=1 -run 'TestFixture2019/TestDay04' advent/2019

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