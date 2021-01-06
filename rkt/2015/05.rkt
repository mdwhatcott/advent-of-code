#!/usr/bin/env racket

#lang racket/base

(require rackunit
         racket/string
         racket/list
         racket/set
         racket/file)

(define vowels
  (list->set (string->list "aeiou")))

(define (vowel? c)
  (set-member? vowels c))

(define (vowels-only s)
  (string-join (map string (filter vowel? (string->list s))) ""))

(define (has-double-letter s)
  (for/first ([i (in-range 1 (string-length s))]
              #:when (char=? (string-ref s (sub1 i))
                             (string-ref s i)))
    #t))

(define (nice1? s)
  (and (has-double-letter s)
       (<= 3 (string-length (vowels-only s)))
       (not (string-contains? s "ab"))
       (not (string-contains? s "cd"))
       (not (string-contains? s "pq"))
       (not (string-contains? s "xy"))))

(define (count-nice1 l)
  (for/sum ([s l]
            #:when (and (nice1? s)))
    1))

(define INPUT
  (string-split (file->string "05.txt") "\n"))

(define answer1
  (count-nice1 INPUT))

(printf "Part 1: ~a ~a~n"
  (= 258 answer1) answer1)

(define (has-double-pair s)
  (for*/first ([i (in-range 0 (- (string-length s) 1))]
               [j (in-range 0 (- (string-length s) 1))]
               #:when (and (> (abs (- i j)) 1)
                           (string=? (substring s i (+ i 2))
                                     (substring s j (+ j 2)))))
    #t))

(define (string-reverse s)
  (string-join (map string (reverse (string->list s))) ""))

(define (has-anagram3 s)
  (for/first ([i (in-range 0 (- (string-length s) 2))]
              #:when (string=? (substring s i (+ i 3))
                               (string-reverse(substring s i (+ i 3))))) 
    #t))

(define (nice2? s)
  (and (has-double-pair s)
       (has-anagram3 s)))

(define (count-nice2 l)
  (for/sum ([s l]
            #:when (and (nice2? s)))
    1))

(define answer2
  (count-nice2 INPUT))

(printf "Part 2: ~a ~a~n"
  (= 53 answer2) answer2)

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(test-equal? "vowels-only 01"
  (vowels-only "abcdefghijklmnopqrstuvwxyz")
  "aeiou")

(test-false "has-double-letter 01" (has-double-letter "abc"))
(test-true  "has-double-letter 02" (has-double-letter "abbc"))

(test-true  "Example 01" (nice1? "ugknbfddgicrmopn"))
(test-true  "Example 02" (nice1? "aaa"))
(test-false "Example 03" (nice1? "jchzalrnumimnmhp"))
(test-false "Example 04" (nice1? "haegwjzuvuyypxyu"))
(test-false "Example 05" (nice1? "dvszwmarrgswjxmb"))

(test-true  "has-double-pair 01" (has-double-pair "abab"))
(test-false "has-double-pair 02" (has-double-pair "abba"))
(test-false "has-double-pair 03" (has-double-pair "aaa"))

(test-true  "has-anagram3 01" (has-anagram3 "aaa"))
(test-true  "has-anagram3 02" (has-anagram3 "aba"))
(test-false "has-anagram3 03" (has-anagram3 "abba"))

(test-true  "Example 06" (nice2? "qjhvhtzxzqqjkmpb"))
(test-true  "Example 07" (nice2? "xxyxx"))
(test-false "Example 08" (nice2? "uurcxstgmygtbstg"))
(test-false "Example 09" (nice2? "ieodomkazucvgmuy"))