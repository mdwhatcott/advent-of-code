#lang racket/base

(require racket/port)
(require racket/string)

(define part1-outcomes
  (hash "A X" (+ 3 1)
        "A Y" (+ 6 2)
        "A Z" (+ 0 3)
        "B X" (+ 0 1)
        "B Y" (+ 3 2)
        "B Z" (+ 6 3)
        "C X" (+ 6 1)
        "C Y" (+ 0 2)
        "C Z" (+ 3 3)))

(define part2-outcomes
  (hash "A X" (+ 0 3)
        "A Y" (+ 3 1)
        "A Z" (+ 6 2)
        "B X" (+ 0 1)
        "B Y" (+ 3 2)
        "B Z" (+ 6 3)
        "C X" (+ 0 2)
        "C Y" (+ 3 3)
        "C Z" (+ 6 1)))

(define sample-data "A Y
B X
C Z")

(define input-data (port->string (open-input-file "day02.txt")))

(define (play input outcomes)
  (for/sum ([round (string-split input "\n")])
    (hash-ref outcomes round)))

(require rackunit)

(test-equal? "part1-sample" 15 (play sample-data part1-outcomes))
(test-equal? "part2-sample" 12 (play sample-data part2-outcomes))

(test-equal? "part1-input" 10994 (play input-data part1-outcomes))
(test-equal? "part2-input" 12526 (play input-data part2-outcomes))