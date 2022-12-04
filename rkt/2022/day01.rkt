#lang racket/base

(require racket/list)
(require racket/port)
(require racket/string)

(define sample-data "1000
2000
3000

4000

5000
6000

7000
8000
9000

10000")

(define (groups input)
  (sort
   (for/list ([group (string-split input "\n\n")])
     (apply + (map string->number (string-split group "\n")))) >))

(require rackunit)

(test-equal? "part1-sample" 24000 (apply + (take (groups sample-data) 1)))
(test-equal? "part2-sample" 45000 (apply + (take (groups sample-data) 3)))

(define input-data (port->string (open-input-file "day01.txt")))

(test-equal? "part1" 71924  (apply + (take (groups input-data) 1)))
(test-equal? "part2" 210406 (apply + (take (groups input-data) 3)))
