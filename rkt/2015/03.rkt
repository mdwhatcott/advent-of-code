#!/usr/bin/env racket

#lang racket/base

(require racket/file
         racket/string
         racket/list
         racket/set)

(define INPUT
  (string->list (file->string "03.txt")))

(define (move from direction)
  (cond [(char=? #\^ direction) (cons (car from) (add1 (cdr from)))]
        [(char=? #\v direction) (cons (car from) (sub1 (cdr from)))]
        [(char=? #\< direction) (cons (sub1 (car from)) (cdr from))]
        [(char=? #\> direction) (cons (add1 (car from)) (cdr from))]))

(define (visit steps houses)
  (if (= (length houses) (length steps))
      houses
      (visit steps (append houses
          (list (move (last houses)
                      (list-ref steps (sub1 (length houses)))))))))

(define answer1
  (set-count (list->set (visit INPUT (list '(0 . 0))))))

(printf "Part 1: ~a ~a ~n"
  (= 2572 answer1) answer1)

(define santa-moves
  (for/list ([i (in-naturals)]
             [m INPUT]
             #:when (even? i)) m))

(define robot-moves
  (for/list ([i (in-naturals)]
             [m INPUT]
             #:when (odd? i)) m))

(define answer2
  (set-count
    (set-union
      (list->set (visit santa-moves (list '(0 . 0))))
      (list->set (visit robot-moves (list '(0 . 0)))))))

(printf "Part 2: ~a ~a ~n"
  (= 2631 answer2) answer2)
