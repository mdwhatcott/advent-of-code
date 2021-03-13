#!/usr/bin/env racket

#lang racket/base

(require racket/file
         racket/list
         racket/string)

(define direction
  #hash((#\U . 0)
        (#\R . 1)
        (#\D . 2)
        (#\L . 3)))

(define (key-code keypad start lines)
  (define (move from to)
    (list-ref (hash-ref keypad from)
              (hash-ref direction to)))
  (define (moves from line)
    (if (= (string-length line) 0) from
      (moves (move from (string-ref line 0))
             (substring line 1))))
  (if (empty? lines) ""
    (let ([next (moves start (car lines))])
      (string-append
        (string next)
        (key-code keypad next (cdr lines))))))

(define keypad1
  ; 1 2 3
  ; 4 5 6
  ; 7 8 9
  #hash((#\1 . (#\1  #\2  #\4  #\1))
        (#\2 . (#\2  #\3  #\5  #\1))
        (#\3 . (#\3  #\3  #\6  #\2))
        (#\4 . (#\1  #\5  #\7  #\4))
        (#\5 . (#\2  #\6  #\8  #\4))
        (#\6 . (#\3  #\6  #\9  #\5))
        (#\7 . (#\4  #\8  #\7  #\7))
        (#\8 . (#\5  #\9  #\8  #\7))
        (#\9 . (#\6  #\9  #\9  #\8))))
;       from     up right down left

(define keypad2
  ;     1
  ;   2 3 4
  ; 5 6 7 8 9
  ;   A B C
  ;     D
  #hash((#\1 . (#\1  #\1  #\3  #\1))
        (#\2 . (#\2  #\3  #\6  #\2))
        (#\3 . (#\1  #\4  #\7  #\2))
        (#\4 . (#\4  #\4  #\8  #\3))
        (#\5 . (#\5  #\6  #\5  #\5))
        (#\6 . (#\2  #\7  #\A  #\5))
        (#\7 . (#\3  #\8  #\B  #\6))
        (#\8 . (#\4  #\9  #\C  #\7))
        (#\9 . (#\9  #\9  #\9  #\8))
        (#\A . (#\6  #\B  #\A  #\A))
        (#\B . (#\7  #\C  #\D  #\A))
        (#\C . (#\8  #\C  #\C  #\B))
        (#\D . (#\B  #\D  #\D  #\D))))

(let* ([lines (string-split (file->string "02.txt") "\n")]
       [part1 (key-code keypad1 #\5 lines)]
       [part2 (key-code keypad2 #\5 lines)])
  (printf "[~a] part 1: ~a ~n" (equal? "74921" part1) part1)
  (printf "[~a] part 2: ~a ~n" (equal? "A6B35" part2) part2))

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(require rackunit)

(test-equal? "example 1, part 1 keypad"
  (key-code keypad1 #\5 (list "ULL" "RRDDD" "LURDL" "UUUUD")) "1985")
(test-equal? "example 1, part 2 keypad"
  (key-code keypad2 #\5 (list "ULL" "RRDDD" "LURDL" "UUUUD")) "5DB3")



