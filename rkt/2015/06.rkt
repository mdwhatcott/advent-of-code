#!/usr/bin/env racket

#lang racket/base

(require racket/file
         racket/list
         racket/string)

(define INPUT
  (string-split (file->string "06.txt") "\n"))

(define OFF    0)
(define ON     1)
(define TOGGLE 2)

(define (on? lights light)
  (= ON (hash-ref lights light OFF)))

(define (turn-on lights light)
  (hash-set lights light ON))

(define (turn-off lights light)
  (hash-remove lights light))

(define (toggle lights light)
  (if (on? lights light)
      (turn-off lights light)
      (turn-on lights light)))

(define (apply-to-lights proc lights spots)
  (if (empty? spots)
     lights
     (apply-to-lights proc (proc lights (car spots)) (cdr spots))))

(define (turn-on-lights lights spots)
  (apply-to-lights turn-on lights spots))

(define (turn-off-lights lights spots)
  (apply-to-lights turn-off lights spots))

(define (toggle-lights lights spots)
  (apply-to-lights toggle lights spots))

(define (parse-action line)
  (cond [(string-contains? line "off")    turn-off-lights]
        [(string-contains? line "on")     turn-on-lights]
        [(string-prefix?   line "toggle") toggle-lights]))

(define (parse-numerals line)
  (map string->number
    (string-split
      (string-replace
        (string-replace
          (string-replace
            (string-replace
              (string-replace
                (string-replace line
                  "toggle " "")
                "turn " "")
              "off " "")
            "on " "")
          "through " "")
        "," " "))))

(define (numerals2range numerals)
  (light-range (first   numerals)
               (second  numerals)
               (third   numerals)
               (fourth  numerals)))

(define (light-range x1 y1 x2 y2)
  (for*/list ([y (in-range y1 (add1 y2))]
              [x (in-range x1 (add1 x2))]) (cons x y)))

(define (apply-instruction lights line)
  ((parse-action line) lights (numerals2range (parse-numerals line))))

(define (setup-lights lights instructions)
  (if (empty? instructions)
      lights
      (setup-lights (apply-instruction lights (first instructions))
                    (rest instructions))))

(define (count-lights-on lights)
  (for/sum ([(k v) lights]) v))

(define (part1 input)
  (count-lights-on (setup-lights (hash) input)))

(define answer1
  (part1 INPUT))

(printf "Part 1: ~a ~a~n"
  (= 0 answer1) answer1)

(require rackunit)

(test-equal? "parse-numerals"
  (parse-numerals "turn off 499,499 through 500,500")
  (list 499 499 500 500))

(test-equal? "light-range"
  (light-range 2 1 4 4)
  (list (cons 2 1) (cons 3 1) (cons 4 1)
        (cons 2 2) (cons 3 2) (cons 4 2)
        (cons 2 3) (cons 3 3) (cons 4 3)
        (cons 2 4) (cons 3 4) (cons 4 4)))

(test-equal? "light-range"
  (light-range 2 1 4 1)
  (list (cons 2 1) (cons 3 1) (cons 4 1)))

(define test-spots (list
  (cons 2 1)
  (cons 3 1)
  (cons 4 1)))
(define all-off #hash())
(define all-on #hash(
  ((2 . 1) . 1)
  ((3 . 1) . 1)
  ((4 . 1) . 1)))
(define one-on #hash(
  ((3 . 1) . 1)))

(test-equal? "turn-on-lights"  (turn-on-lights all-off test-spots) all-on)
(test-equal? "turn-off-lights" (turn-off-lights all-on test-spots) all-off)
(test-equal? "toggle-lights"   (toggle-lights one-on test-spots) '#hash(
  ((2 . 1) . 1)
 ;((3 . 1) . 0) ; off == removed
  ((4 . 1) . 1)))

(test-equal? "apply-instruction"
  (apply-instruction all-off "turn on 2,1 through 4,1")
  all-on)

(test-equal? "part 1 acceptance"
  (- 1000000 1000 4)
  (part1 (list
    "turn on 0,0 through 999,999"
    "toggle 0,0 through 999,0"
    "turn off 499,499 through 500,500")))
