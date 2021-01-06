#!/usr/bin/env racket
#lang racket/base

(require racket/file)
(require racket/string)
(require racket/list)

(define LINES (file->lines "day02.txt"))

(define (parse-box line)
  (map string->number (string-split line "x")))

(define BOXES (map parse-box LINES))

(define (surface-area box)
  (+ (* 2 (first box) (second box))
     (* 2 (second box) (third box))
     (* 2 (third box) (first box))))

(define (smallest-face-area box)
  (min (* (first box) (second box))
       (* (second box) (third box))
       (* (third box) (first box))))

(define (paper-required box)
  (+ (surface-area box) (smallest-face-area box)))

(define (part1 boxes)
  (apply + (map paper-required boxes)))

(define answer1 (part1 BOXES))

(printf "Part 1: ~a ~a ~n" 
  (= 1606483 answer1) answer1)

(define (perimeter length width)
  (+ (* 2 length) (* 2 width)))

(define (smallest-perimeter box)
  (min (perimeter (first box) (second box))
       (perimeter (second box) (third box))
       (perimeter (third box) (first box))))

(define (volume box)
  (apply * box))

(define (ribbon-required box)
  (+ (smallest-perimeter box) (volume box)))

(define (part2 boxes)
  (apply + (map ribbon-required BOXES)))

(define answer2 (part2 BOXES))

(printf "Part 2: ~a ~a ~n" 
  (= 3842356 answer2) answer2)
