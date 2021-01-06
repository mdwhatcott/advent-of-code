;; Credit: https://www.reddit.com/user/bhrgunatha/
;; https://www.reddit.com/r/adventofcode/comments/k7ndux/2020_day_06_solutions/gesaa4r/

#lang racket/base

(require racket/file)
(require racket/string)
(require racket/set)

(define INPUT
  (string-split (file->string "day06.txt") "\n\n"))

(define (string-set s)
  (list->set (string->list s)))

(define alphabet-set
  (string-set "abcdefghijklmnopqrstuvwxyz"))

(define (part-01 input)
  (for/sum ([group (in-list input)])
    (answer-count group)))

(define (answer-count group)
  (set-count (string-set (string-replace group "\n" ""))))

(define (part-02 input)
  (for/sum ([line (in-list input)])
    (for/fold ([s alphabet-set]
               #:result (set-count s))
              ([p (string-split line "\n")])
      (set-intersect s (string-set p)))))

(= (part-01 INPUT) 6775)
(= (part-02 INPUT) 3356)