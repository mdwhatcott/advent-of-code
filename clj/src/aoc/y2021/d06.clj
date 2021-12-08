(ns aoc.y2021.d06)

; Inspiration:
; Fred Overflow - https://www.youtube.com/watch?v=umHxPjNXD6Y&t=82
(defn swim
  [, [_0 _1 _2 _3 _4 _5,,, _6,,,, _7 _8]]                   ; destructure the input sequence
  [,, _1 _2 _3 _4 _5 _6 (+ _7 _0) _8 _0])                   ; manually shift to create output sequence

(defn part1 [data days]
  (as-> (frequencies data) $
        (for [i (range 9)]
          (get $ i 0))
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
(defn- swim--most-general [fishes]
  (for [i (range 9)]
    (case i
      6 (+ (nth fishes 0)                                   ; parents
           (nth fishes (inc i)))
      8 (nth fishes 0)                                      ; guppies
      (nth fishes (inc i)))))

; WARNING: exponential growth makes this function unbearably slow at days >= 200
; Only included here for completeness.
(defn- swim--recursive [days timer]
  (let [days (- days timer)]
    (if-not (pos? days)
      1 (+ (swim--recursive days 7)
           (swim--recursive days 9)))))

; WARNING: gets too big really fast. A literal translation of the instructions,
; each fish is an element in an exponentially expanding sequence. Just here to
; show where I started.
(defn- swim--append [timers]
  (let [parents (filter #{0} timers)
        parents (for [_ parents] 8)
        nexts   (map dec timers)
        nexts   (map #(if (neg? %1) 6 %1) nexts)]
    (concat nexts parents)))
