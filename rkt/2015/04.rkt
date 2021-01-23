#!/usr/bin/env racket

#lang racket/base

(require file/md5
         racket/string)

(define INPUT "ckczppom")

(define (md5-prefix-match? prefix input)
  (string-prefix? (bytes->string/utf-8 (md5 input)) prefix))

(define answer1
  (for/first ([i (in-naturals)]
              #:when (md5-prefix-match?
              "00000"
              (string-append INPUT (number->string i))))
    i))

(printf "Part 1: ~a ~a ~n"
  (= 117946 answer1) answer1)

(define answer2
  (for/first ([i (in-naturals)]
              #:when (md5-prefix-match?
              "000000"
              (string-append INPUT (number->string i))))
    i))

(printf "Part 2: ~a ~a ~n"
  (= 3938038 answer2) answer2)
