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

(define (light-range xyxy)
  (stream->list
    (in-range (* (list-ref xyxy 0) (list-ref xyxy 1))
        (add1 (* (list-ref xyxy 2) (list-ref xyxy 3))))))

(define (turn-off lights light)
  (hash-set lights light 0))

(define (turn-on lights light)
  (hash-set lights light 1))

(define (toggle lights light)
  (hash-set lights light (if (= 0 (hash-ref lights light 0)) 1 0)))

(define (hash-set-range proc lights range)
  (if (empty? range)
      lights
      (hash-set-range proc
                      (proc lights (car range))
                      (cdr range))))

(define (apply-instruction lights action range)
  (cond [(= action 0) (hash-set-range turn-off lights range)]
        [(= action 1) (hash-set-range turn-on  lights range)]
        [(= action 2) (hash-set-range toggle   lights range)]))

(define (setup-lights lights instructions)
  (if (empty? instructions)
      lights
      (setup-lights
        (apply-instruction
          lights
          (parse-action (car instructions))
          (light-range (parse-numerals (car instructions))))
        (cdr instructions))))

(define (count-lights-on lights)
  (for/sum ([(k v) lights]) v))

(define (part1 input)
  (count-lights-on (setup-lights (hash) input)))

; (define answer1
;   (part1 INPUT))

; (printf "Part 1: ~a ~a~n"
;   (= 0 answer1) answer1)

(require rackunit)

(test-equal? "light-range"
  (light-range (list 1 2 3 4))     ; (* 1 2) ... (* 3 4)
  (list 2 3 4 5 6 7 8 9 10 11 12)) ;    2    ...    12

(test-equal? "light-range"
  (light-range (list 500 500 501 501))
  (list 251001 252005))

(test-equal? "parse-numerals"
  (parse-numerals "turn off 499,499 through 500,500")
  (list 500 500 501 501))


(test-equal? "parse-action (off)"    (parse-action "turn off") 0)
(test-equal? "parse-action (on)"     (parse-action "turn on")  1)
(test-equal? "parse-action (toggle)" (parse-action "toggle")   2)

(test-equal? "apply-instruction (off)"
  (apply-instruction (hash) 0 (list 4 5 6))
  #hash((4 . 0) (5 . 0) (6 . 0)))

(test-equal? "apply-instruction (on)"
  (apply-instruction (hash) 1 (list 4 5 6))
  #hash((4 . 1) (5 . 1) (6 . 1)))

(test-equal? "apply-instruction (toggle)"
  (apply-instruction #hash((4 . 1) (5 . 0)) 2 (list 4 5 6))
  #hash((4 . 0) (5 . 1) (6 . 1)))

(test-equal? "part 1 acceptance"
  (- 1000000 1000 4)
  (count-lights-on
    (setup-lights (hash) (list
      "turn on 0,0 through 999,999"
      "toggle 0,0 through 999,0"
      "turn off 499,499 through 500,500"))))
