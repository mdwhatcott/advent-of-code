#!/usr/bin/make

day  := 03
year := 20

test.day: fmt
	go test -timeout=10s -v advent/20${year}/day${day} && \
	go test -timeout=100s -v -run 'TestFixture20${year}/TestDay${day}' advent/20${year}

test.2020: fmt
	go test -timeout=10s advent/2020/...

test.2019: fmt
	go test -timeout=10s advent/2019/...

test.2018: fmt
	go test -timeout=10s advent/2018/...

test.2017: fmt
	go test -timeout=10s advent/2017/...

test.2016: fmt
	go test -timeout=20s advent/2016/...

test.2015: fmt
	go test -timeout=10s advent/2015/...

test: test.2015 test.2016 test.2017 test.2018 test.2019

fmt:
	go fmt ./...

.PHONY: fmt test test.2015 test.2016 test.2017 test.2018 test.2019 test.2020 test.Day