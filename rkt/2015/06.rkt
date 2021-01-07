#!/usr/bin/env racket

#lang racket/base

(require racket/file
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

(define (light-range xyxy)
  (in-range (* (list-ref xyxy 0) (list-ref xyxy 1))
      (add1 (* (list-ref xyxy 2) (list-ref xyxy 3)))))

(define (apply-instruction lights instruction)
  lights)

(require rackunit)

(test-equal? "light-range"
  (stream->list (light-range (list 1 2 3 4))) ; (* 1 2) ... (* 3 4)
  (list 2 3 4 5 6 7 8 9 10 11 12))            ;    2    ...    12

(test-equal? "parse-action" (parse-action "turn off") 0)
(test-equal? "parse-action" (parse-action "turn on")  1)
(test-equal? "parse-action" (parse-action "toggle")   2)

(test-equal? "apply-instruction"
  (apply-instruction (hash) "turn on 0,0 through 1,1")
  (hash))