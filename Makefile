#!/usr/bin/make

tcr: go

go:
	cd go && make && cd -

rkt:
	cd rkt/2015 && make && cd -

.PHONY: tcr go rkt