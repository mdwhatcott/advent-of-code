#!/usr/bin/env racket

#lang racket/base

(define (bits n)
  (let loop ([w n] [r 0])
    (if (not (positive? w)) r
      (loop (bitwise-and w (sub1 w)) (add1 r)))))

(define (is-wall? seed x y)
  (if (or (negative? x) (negative? y)) #f
    (even? (bits (+ seed (* x x) (* 3 x) (* 2 x y) y (* y y))))))

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(require rackunit)

(test-case "hamming weight"
  (check-eq? (bits 0) 0)   ; 0000
  (check-eq? (bits 1) 1)   ; 0001
  (check-eq? (bits 2) 1)   ; 0010
  (check-eq? (bits 4) 1)   ; 0100
  (check-eq? (bits 8) 1)   ; 1000
  (check-eq? (bits 3) 2)   ; 0011
  (check-eq? (bits 15) 4)  ; 1111
  (check-eq? (bits 255) 8) ; 11111111
)

(define TEST-SEED 10)

(test-case "wall vs. space"
  (check-true (is-wall? TEST-SEED 0 0))
  (check-true (is-wall? TEST-SEED 1 1))
  (check-false (is-wall? TEST-SEED 9 6))
  (check-false (is-wall? TEST-SEED -1 0))
  (check-false (is-wall? TEST-SEED 0 -1))
  (check-false (is-wall? TEST-SEED -1 -1))
)

; TODO: bfs