#!/usr/bin/env racket

#lang racket/base

(require racket/file
         racket/list
         racket/string
         racket/set
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

; (define (apply-instruction all instruction)
  ; (for/??))

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

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