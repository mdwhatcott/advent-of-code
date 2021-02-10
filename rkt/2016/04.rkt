#!/usr/bin/env racket

#lang racket/base

(require racket/file
         racket/list
         racket/string
         racket/match)

(define (room-checksum line)
  (second (regexp-match #px"\\[(.+)\\]" line)))

(define (sector-id line)
  (second (regexp-match #px"-(\\d+)\\[" line)))

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(require rackunit)

(define EXAMPLE1 "aaaaa-bbb-z-y-x-123[abxyz]")

(test-equal? "room-checksum" (room-checksum EXAMPLE1) "abxyz")
(test-equal? "sector-id" (sector-id EXAMPLE1) "123")