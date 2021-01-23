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

(define (turn from instruction)
  (if (eq? (string-ref instruction 0) #\R)
      (R from)
      (L from)))

(define (walk direction from)
  (cons (+ (car from) (car direction))
        (+ (cdr from) (cdr direction))))

(define (parse-distance d)
  (string->number (substring d 1)))

(define (walk-n n direction from)
  (if (= n 0) from
    (walk-n (sub1 n) direction (walk direction from))))

(define (travel step facing from)
  (walk-n (parse-distance step)
          (turn facing step)
          from))

(define (travel-all steps)
  (let loop ([facing N]
             [from (cons 0 0)]
             [steps steps])
    (if (empty? steps) from
      (loop (turn facing (first steps))
            (travel (first steps) facing from)
            (rest steps)))))

(define (pair2list p)
  (list (car p) (cdr p)))

(define (part1 input)
  (apply + (map abs (pair2list
    (travel-all (string-split input ", "))))))

(define INPUT (file->string "01.txt"))

(define answer1 (part1 INPUT))

(printf "part 1: ~a ~a ~n"
  (= 291 answer1) answer1)

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

(test-equal? "walk" (walk N (cons 2 2)) (cons 2 3))
(test-equal? "walk" (walk S (cons 2 2)) (cons 2 1))
(test-equal? "walk" (walk W (cons 2 2)) (cons 1 2))
(test-equal? "walk" (walk E (cons 2 2)) (cons 3 2))

(test-equal? "parse-distance" (parse-distance "R32") 32)
(test-equal? "parse-distance" (parse-distance "L6") 6)

(test-equal? "walk-n" (walk-n 10 E (cons 0 0)) (cons 10 0))

(test-equal? "travel" (travel "R10" N (cons 0 0)) (cons 10 0))

(test-equal? "travel-all"
  (travel-all (list "R5" "L5" "R5" "R3"))
  (cons 10 2))

(define (enumerate-steps raw)
  ; (for/fold ([steps (list)])
            ; ([raw-step raw])
    ; (append steps
      ; (for/list ([])))))
  (list E E E N W W S S W N N N E S S S))

(test-equal? "enumerate-steps"
  (enumerate-steps (list "R3" "L1" "L2" "L2" "R1" "R3" "R1" "R3"))
  (list E E E N W W S S W N N N E S S S))


; R3 L1 L2 L2 R1 R3 R1 R3
; E E E N W W S S W N N N E S S S