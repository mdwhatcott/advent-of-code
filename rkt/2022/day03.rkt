#lang racket/base

(require racket/list
         racket/port
         racket/set
         racket/string
         rackunit)

(define sample-data
  "vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw")

(define input-data (port->string (open-input-file "day03.txt")))

(define char-A (char->integer #\A))
(define char-a (char->integer #\a))

(define (sack-sum input)
  (for/sum ([sack (string-split input "\n")])
    (let* ([mid (/ (string-length sack) 2)]
           [a (substring sack 0 mid)]
           [b (substring sack mid)]
           [a (list->set (string->list a))]
           [b (list->set (string->list b))]
           [ab (set->list (set-intersect a b))]
           [c (first ab)]
           [n (char->integer c)])
      (if (char-upper-case? c)
          (+ 26 1 (- n char-A))
          (+ 0 1 (- n char-a))))))

(test-equal? "part1-sample" (sack-sum sample-data) 157)
(test-equal? "part1" (sack-sum input-data) 8018)

