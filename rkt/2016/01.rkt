#!/usr/bin/env racket

#lang racket/base

(require racket/file
         racket/list
         racket/string)

(define N (cons 0 1))
(define S (cons 0 -1))
(define W (cons -1 0))
(define E (cons 1 0))

(define (R from)
  (cond ((eq? from N) E)
        ((eq? from E) S)
        ((eq? from S) W)
        ((eq? from W) N)))

(define (L from)
  (R (R (R from))))

(define (walk direction from)
  (cons (+ (car from) (car direction))
        (+ (cdr from) (cdr direction))))

(define (parse-distance d)
  (string->number (substring d 1)))

(define INPUT (file->string "01.txt"))

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(require rackunit)

(test-equal? "R" (R N) E)
(test-equal? "R" (R E) S)
(test-equal? "R" (R S) W)
(test-equal? "R" (R W) N)

(test-equal? "L" (L N) W)
(test-equal? "L" (L W) S)
(test-equal? "L" (L S) E)
(test-equal? "L" (L E) N)
