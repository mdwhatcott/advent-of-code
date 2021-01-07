#!/usr/bin/env racket

#lang racket/base

(require racket/file
         racket/list
         racket/string
         racket/stream)

(define INPUT
  (string-split (file->string "06.txt") "\n"))

(define (parse-action line)
  (cond [(string-contains? line "off")    0]
        [(string-contains? line "on")     1]
        [(string-prefix?   line "toggle") 2]))

(define (parse-numerals line)
  (map add1
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
          "," " ")))))


; (define answer1
;   (part1 INPUT))

; (printf "Part 1: ~a ~a~n"
;   (= 0 answer1) answer1)

(require rackunit)

(test-equal? "parse-numerals"
  (parse-numerals "turn off 499,499 through 500,500")
  (list 500 500 501 501))

(test-equal? "parse-action (off)"    (parse-action "turn off") 0)
(test-equal? "parse-action (on)"     (parse-action "turn on")  1)
(test-equal? "parse-action (toggle)" (parse-action "toggle")   2)

; (test-equal? "part 1 acceptance"
;   (- 1000000 1000 4)
;   (count-lights-on
;     (setup-lights (hash) (list
;       "turn on 0,0 through 999,999"
;       "toggle 0,0 through 999,0"
;       "turn off 499,499 through 500,500"))))
