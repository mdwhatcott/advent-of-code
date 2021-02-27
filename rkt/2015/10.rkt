#!/usr/bin/env racket

#lang racket/base

(require racket/file
         racket/list
         racket/string)

(define (empty-string? s)
  (not (non-empty-string? s)))

(define (say in)
  (if (empty-string? in) ""
    (let* ([initial (string-ref in 0)]
           [count (for/sum ([c in] #:break (not (equal? c initial))) 1)])
      (string-append
        (number->string count)
        (string initial)
        (say (substring in count))))))

(define (say-n in n)
  (if (= n 0) in
    (say-n (say in) (sub1 n))))

(let* ([input "1113222113"]
       [result1 (say-n input 40)]
       [result2 (say-n result1 10)]
       [answer1 (string-length result1)]
       [answer2 (string-length result2)])
  (printf "Part 1: ~a ~a~n" (= 252594 answer1) answer1)
  (printf "Part 2: ~a ~a~n" (= 3579328 answer2) answer2))

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(require rackunit)

(test-equal? "hi" (say "1") "11")
(test-equal? "hi" (say "11") "21")
(test-equal? "hi" (say "21") "1211")
(test-equal? "hi" (say "1211") "111221")
(test-equal? "hi" (say "111221") "312211")
(test-equal? "hi" (say-n "1" 5) "312211")