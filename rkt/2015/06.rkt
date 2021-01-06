#!/usr/bin/env racket

#lang racket/base

(require racket/file
         racket/list
         racket/string
         racket/set
         racket/stream
         rackunit)

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

(define (parse-instruction line)
  (list (parse-action line)
        (parse-numerals line)))

(define (light-range xyxy)
  (in-range (* (list-ref xyxy 0) (list-ref xyxy 1))
      (add1 (* (list-ref xyxy 2) (list-ref xyxy 3)))))

; (define (apply-instruction all instruction)
  ; (for/??))

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(test-equal? "light-range"
  (stream->list (light-range (list 1 2 3 4))) ; (* 1 2) ... (* 3 4)
  (list 2 3 4 5 6 7 8 9 10 11 12))            ;    2    ...    12

(test-equal? "parse-instruction (off = 0)" 
  (parse-instruction "turn off 0,1 through 2,3")
  (list 0 (list 1 2 3 4)))

(test-equal? "parse-instruction (on = 1)" 
  (parse-instruction "turn on 0,1 through 2,3")
  (list 1 (list 1 2 3 4)))

(test-equal? "parse-instruction (toggle = 2)" 
  (parse-instruction "toggle 0,1 through 2,3")
  (list 2 (list 1 2 3 4)))

; (test-equal? "apply-instruction"
  ; (apply-instruction (hash) (list 1 (list 1 1 2 2)))
  ; (hash))