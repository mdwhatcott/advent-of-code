(ns aoc.y2021.d06)

; Inspiration:
; Fred Overflow - https://www.youtube.com/watch?v=umHxPjNXD6Y&t=82
(defn swim
  [[_0 _1 _2 _3 _4 _5,,, _6,,,, _7 _8]]                     ; destructure the input sequence
  ,[_1 _2 _3 _4 _5 _6 (+ _7 _0) _8 _0])                     ; manually shift to create output sequence

(defn part1 [data days]
  (as-> (frequencies data) $
        (map #(get $ % 0) (range 9))
        (iterate swim $)
        (nth $ days)
        (apply + $)))

(defn part2 [data]
  (part1 data 256))

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

; Inspiration:
; https://www.reddit.com/r/adventofcode/comments/r9z49j/comment/hnj18bu/
; https://www.reddit.com/r/adventofcode/comments/r9z49j/comment/hnfaisu/
; https://www.reddit.com/r/adventofcode/comments/r9z49j/comment/hnfd2w5/
(defn swim--most-general [fishes]
  (for [i (range 9)]
    (case i
      6 (+ (nth fishes 0)                                   ; parents
           (nth fishes (inc i)))
      8 (nth fishes 0)                                      ; guppies
      (nth fishes (inc i)))))

; WARNING: exponential growth makes this function unbearably slow at days >= 200
; Only included here for completeness.
(defn- swim-recursive [days timer]
  (let [days (- days timer)]
    (if-not (pos? days)
      1 (+ (swim-recursive days 7)
           (swim-recursive days 9)))))
