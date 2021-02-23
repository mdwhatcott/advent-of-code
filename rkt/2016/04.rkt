#!/usr/bin/env racket

#lang racket/base

(require racket/file
         racket/list
         racket/string
         racket/match)

(define (room-checksum line)
  (second (regexp-match #px"\\[(.+)\\]" line)))

(define (sector-id line)
  (string->number (second (regexp-match #px"-(\\d+)\\[" line))))

(define (room-name line)
  (let* ([chars (string->list line)]
         [backward (reverse chars)]
         [next (list->string (reverse (cdr backward)))])
    (if (eq? #\- (first backward)) next
      (room-name next))))

(define (char-counts name)
  (for/fold ([counts (hash)])
            ([char (string->list name)]
             #:when (not (equal? char #\-)))
    (hash-set counts char (add1 (hash-ref counts char 0)))))

(define (calculate-checksum name)
  (let* ([counts (char-counts name)]
         [by-frequency-then-alpha (lambda (i j)
           (let* ([I (hash-ref counts i)]
                  [J (hash-ref counts j)]
                  [CI (char->integer i)]
                  [CJ (char->integer j)])
             (if (= I J)
               (< CI CJ)
               (> I J))))]
         [keys (hash-keys counts)]
         [sorted (sort keys by-frequency-then-alpha)]
         [sorted-text (list->string sorted)]
         [chop (min 5 (length sorted))])
    (substring sorted-text 0 chop)))

(define (valid-room? line)
  (let* ([name (room-name line)]
         [expected (room-checksum line)]
         [actual (calculate-checksum name)])
    (equal? expected actual)))

(define A (char->integer #\a))
(define LETTER-COUNT 26)

(define (ceasar-cypher c offset)
  (let* ([n (char->integer c)]
         [alpha-id (- n A)]
         [shifted (+ alpha-id offset)]
         [alpha-id (modulo shifted LETTER-COUNT)])
  (integer->char (+ A alpha-id))))

(define (decrypt line)
  (let* ([name (room-name line)]
         [sector (sector-id line)]
         [offset (modulo sector LETTER-COUNT)]
         [characters (string->list name)]
         [shifted
           (for/list ([c characters])
             (if (equal? c #\-) #\space
               (ceasar-cypher c offset)))])
    (list->string shifted)))

(define (contains-north-pole-objects? line)
  (let ([name (decrypt line)])
    (and (string-contains? name "north")
         (string-contains? name "pole")
         (string-contains? name "object"))))

(let* ([lines (string-split (file->string "04.txt"))]
       [valid-rooms (filter valid-room? lines)]
       [part1 (for/sum ([room valid-rooms]) (sector-id room))]
       [part2 (for/first ([room valid-rooms]
                          #:when (contains-north-pole-objects? room))
                (sector-id room))])
  (printf "[~a] part 1: ~a ~n" (= 409147 part1) part1)
  (printf "[~a] part 2: ~a ~n" (= 991 part2) part2))

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(require rackunit)

(define EXAMPLE1 "aaaaa-bbb-z-y-x-123[abxyz]")
(define EXAMPLE2 "a-b-c-d-e-f-g-h-987[abcde]")
(define EXAMPLE3 "totally-real-room-200[decoy]")
(define EXAMPLEA "a-b-c-d-123[abcd]")
(define ENCRYPTED "qzmt-zixmtkozy-ivhz-343[hello]")

(test-equal? "room-checksum" (room-checksum EXAMPLE1) "abxyz")
(test-equal? "sector-id" (sector-id EXAMPLE1) 123)
(test-equal? "room-name" (room-name EXAMPLE1) "aaaaa-bbb-z-y-x")
(test-equal? "char-counts"
  (char-counts (room-name EXAMPLE1))
  '#hash((#\a . 5)
         (#\b . 3)
         (#\x . 1)
         (#\y . 1)
         (#\z . 1)))
(test-equal? "calculate-checksum"
  (calculate-checksum (room-name EXAMPLE1)) "abxyz")
(test-equal? "calculate-checksum"
  (calculate-checksum (room-name EXAMPLE2)) "abcde")
(test-equal? "calculate-checksum"
  (calculate-checksum (room-name EXAMPLEA)) "abcd")

(test-equal? "valid-room?" (valid-room? EXAMPLE1) #t)
(test-equal? "valid-room?" (valid-room? EXAMPLE3) #f)

(test-equal? "decrypt-name"
  (decrypt ENCRYPTED) "very encrypted name")
