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

(apply + (take (groups sample-data) 1)) ; part 1 (sample)
(apply + (take (groups sample-data) 3)) ; part 2 (sample)

(define input-data (port->string (open-input-file "input.txt")))

(apply + (take (groups input-data) 1)) ; part 1
(apply + (take (groups input-data) 3)) ; part 2
