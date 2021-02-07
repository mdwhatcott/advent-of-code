#!/usr/bin/env racket

#lang racket/base

(require racket/file
         racket/list
         racket/string)

(define (valid triangle)
  (let ([a (first  triangle)]
        [b (second triangle)]
        [c (third  triangle)])
    (and (< a (+ b c))
         (< b (+ c a))
         (< c (+ a b)))))

(define (count-valid triangles)
  (length (filter valid triangles)))

(define (read-ints)
  (map string->number
    (string-split
      (string-replace
        (file->string "03.txt") "\n" " "))))

(define (parse-triangles-1 numbers)
  (if (empty? numbers) empty
    (let ([a (first numbers)] [b (second numbers)] [c (third numbers)])
      (cons (list a b c) (parse-triangles-1 (list-tail numbers 3))))))

(define (parse-triangles-2 numbers)
  (if (empty? numbers) empty
    (let* ([a1 (first numbers)]  [b1 (fourth numbers)] [c1 (seventh numbers)]
           [a2 (second numbers)] [b2 (fifth numbers)]  [c2 (eighth numbers)]
           [a3 (third numbers)]  [b3 (sixth numbers)]  [c3 (ninth numbers)]
           [t1 (list a1 b1 c1)]  [t2 (list a2 b2 c2)]  [t3 (list a3 b3 c3)])
      (append (list t1 t2 t3) (parse-triangles-2 (list-tail numbers 9))))))

(let* ([numbers (read-ints)]
       [triangles-part-1 (parse-triangles-1 numbers)]
       [triangles-part-2 (parse-triangles-2 numbers)]
       [part1 (count-valid triangles-part-1)]
       [part2 (count-valid triangles-part-2)])
  (printf "[~a] part 1: ~a ~n" (= 983 part1) part1)
  (printf "[~a] part 2: ~a ~n" (= 1836 part2) part2))

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(require rackunit)

(test-equal? "valid-triangles" (valid (list 1 1 1)) #t)
(test-equal? "valid-triangles" (valid (list 5 10 25)) #f)

(test-equal? "count-valid" (count-valid (list (list 1 1 1) (list 5 10 25))) 1)