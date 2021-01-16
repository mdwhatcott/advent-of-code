#!/usr/bin/env racket

#lang racket/base

(require racket/file
         racket/list
         racket/string)

(define UP    (cons 0 1))
(define DOWN  (cons 0 -1))
(define LEFT  (cons -1 0))
(define RIGHT (cons 1 0))

(define (turn-right from)
  (cond ((eq? from UP)    RIGHT)
        ((eq? from RIGHT) DOWN)
        ((eq? from DOWN)  LEFT)
        ((eq? from LEFT)  UP)))

(define (turn-left from)
  (turn-right
    (turn-right
      (turn-right from))))

(define (walk direction from)
  (cons (+ (car from) (car direction))
        (+ (cdr from) (cdr direction))))

(define (parse-turn d)
  (if (eq? (string-ref d 0) #\R)
      turn-right
      turn-left))

(define (parse-distance d)
  (string->number (substring d 1)))

(define (walk-n n direction from)
  (if (= n 0) from
    (walk-n (sub1 n) direction (walk direction from))))

(define (travel step facing from)
  (walk-n (parse-distance step)
          ((parse-turn step) facing)
          from))

(define (travel-all facing from steps)
  (cons 10 2))

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(require rackunit)

(test-equal? "turn-right" (turn-right UP)    RIGHT)
(test-equal? "turn-right" (turn-right RIGHT) DOWN)
(test-equal? "turn-right" (turn-right DOWN)  LEFT)
(test-equal? "turn-right" (turn-right LEFT)  UP)

(test-equal? "turn-left"  (turn-left  UP)    LEFT)
(test-equal? "turn-left"  (turn-left  LEFT)  DOWN)
(test-equal? "turn-left"  (turn-left  DOWN)  RIGHT)
(test-equal? "turn-left"  (turn-left  RIGHT) UP)

(test-equal? "walk" (walk UP    (cons 2 2)) (cons 2 3))
(test-equal? "walk" (walk DOWN  (cons 2 2)) (cons 2 1))
(test-equal? "walk" (walk LEFT  (cons 2 2)) (cons 1 2))
(test-equal? "walk" (walk RIGHT (cons 2 2)) (cons 3 2))

(test-equal? "parse-turn" (parse-turn "R2") turn-right)
(test-equal? "parse-turn" (parse-turn "L2") turn-left)

(test-equal? "parse-distance" (parse-distance "R32") 32)
(test-equal? "parse-distance" (parse-distance "L6") 6)

(test-equal? "walk-n" (walk-n 10 RIGHT (cons 0 0)) (cons 10 0))

(test-equal? "travel" (travel "R10" UP (cons 0 0)) (cons 10 0))

(test-equal? "travel-all"
  (travel-all UP (cons 0 0) (list "R5" "L5" "R5" "R3"))
  (cons 10 2))