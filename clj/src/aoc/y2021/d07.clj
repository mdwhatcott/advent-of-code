(ns aoc.y2021.d07)

(defn abs [n] (if (neg? n) (- n) n))

(defn lowest-fuel-cost [data cost]
  (let [lo (apply min data)
        hi (apply max data)]
    (apply min (for [i (range lo (inc hi))]
                 (apply + (map #(cost (abs (- i %))) data))))))

(defn part1-fuel-cost [distance] distance)
(defn part2-fuel-cost [distance] (/ (* distance (inc distance)) 2))

(defn part1 [data] (lowest-fuel-cost data part1-fuel-cost))
(defn part2 [data] (lowest-fuel-cost data part2-fuel-cost))

; Triangle number formula in part2-fuel-cost courtesy of KnavesRadiant on Reddit:
; https://www.reddit.com/r/adventofcode/comments/rar7ty/2021_day_7_solutions/hnlrphb/
; https://en.wikipedia.org/wiki/Triangular_number
; Formula: n(n+1) / 2