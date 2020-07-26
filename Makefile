#!/usr/bin/make

day  := 13
year := 19

test.day: fmt
	go test -timeout=10s -v -count=1 advent/20${year}/day${day} && \
	go test -timeout=10s -count=1 advent/20${year}/intcode && \
	go test -timeout=100s -v -count=1 -run 'TestFixture20${year}/TestDay${day}' advent/20${year}

test.2019: fmt
	go test -timeout=10s -count=1 advent/2019/...

test.2018: fmt
	go test -timeout=10s -count=1 advent/2018/...

test.2017: fmt
	go test -timeout=10s -count=1 advent/2017/...

test.2016: fmt
	go test -timeout=20s -count=1 advent/2016/...

test.2015: fmt
	go test -timeout=10s -count=1 advent/2015/...

test: test.2015 test.2016 test.2017 test.2018 test.2019

fmt:
	go fmt ./...

.PHONY: fmt test test.2015 test.2016 test.2017 test.2018 test.2019 test.Day