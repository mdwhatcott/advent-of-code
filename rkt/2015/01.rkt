#lang racket/base

(require racket/file
         racket/string)

(define PARENTHESIS
  (string->list (file->string "01.txt")))

(define (paren2num p)
  (if (char=? p #\() 1 -1))

(define MOVES
  (map paren2num PARENTHESIS))

(define (calculate-final-floor moves)
  (apply + moves))

(define (part1 moves)
  (calculate-final-floor moves))

(define answer1 (part1 MOVES))

(printf "part 1: ~a ~a ~n"
  (= 232 answer1) answer1)

(define (find-basement-entry moves)
  (let elevate ([entry 0]
                [floor 0])
    (if (< floor 0)
      entry
      (elevate (add1 entry)
               (+ floor (list-ref moves entry))))))

(define (part2 moves)
  (find-basement-entry MOVES))

(define answer2 (part2 MOVES))

(printf "part 2: ~a ~a ~n"
  (= 1783 answer2) answer2)
