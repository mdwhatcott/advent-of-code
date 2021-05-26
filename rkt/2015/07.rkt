#!/usr/bin/env racket

#lang racket/base

(require racket/file
         racket/list
         racket/string)


(let* ([part1 1]
       [part2 2])
  (printf "[~a] part 1: ~a ~n" (= 1 part1) part1)
  (printf "[~a] part 2: ~a ~n" (= 2 part2) part2))

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(require rackunit)

; (test-equal? "parse-numerals"
;   (parse-numerals "turn off 499,499 through 500,500")
;   (list 499 499 500 500))
