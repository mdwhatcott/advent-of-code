#!/usr/bin/env racket

#lang racket/base

(require racket/file
         racket/list
         racket/set
         racket/string)

(define ORIGIN (cons 0 0))

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

(define (turn facing direction)
  (if (eq? (string-ref direction 0) #\R)
      (R facing)
      (L facing)))

(define (all-steps so-far legs facing)
  (if (empty? legs) so-far
    (let* ([leg        (first legs)]
           [direction  (substring leg 0 1)]
           [now-facing (turn facing direction)]
           [distance   (string->number (substring leg 1))]
           [steps      (for/list ([i (in-range distance)]) now-facing)])
      (all-steps (append so-far steps) (rest legs) now-facing))))

(define (walk at step)
  (cons (+ (car at) (car step))
        (+ (cdr at) (cdr step))))

(define (walk-all at steps)
  (if (empty? steps) at
    (walk-all (walk at (first steps)) (rest steps))))

(define (distance-from-origin at)
  (+ (abs (car at)) (abs (cdr at))))

(define INPUT
  (string-split (file->string "01.txt") ", "))

(define STEPS
  (all-steps (list) INPUT N))

(define (part1 steps)
  (distance-from-origin (walk-all ORIGIN steps)))

(define answer1 (part1 STEPS))

(printf "part 1: ~a ~a ~n"
  (= 291 answer1) answer1)

(define (walk-until-criss-cross visited at steps)
  (if (set-member? visited at) at
    (walk-until-criss-cross (set-add visited at)
                            (walk at (first steps))
                            (rest steps))))
(define (part2 steps)
  (distance-from-origin
    (walk-until-criss-cross (set) ORIGIN steps)))

(define answer2 (part2 STEPS))
(printf "part 2: ~a ~a ~n"
  (= 159 answer2) answer2)

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

(test-equal? "turn L" (turn N "L14") W)
(test-equal? "turn R" (turn N "R14") E)

(test-equal? "walk" (walk ORIGIN N) (cons 0 1))

(test-equal? "all-steps"
  (all-steps (list) (list "R2" "L2" "L1" "L3") N)
  (list E E N N W S S S))

(test-equal? "walk-all"
  (walk-all ORIGIN (list E E N N W W S))
  (cons 0 1))


