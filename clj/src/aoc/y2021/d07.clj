(ns aoc.y2021.d07)

(defn abs [n] (if (neg? n) (- n) n))

(defn part1-fuel-cost [distance] distance)
(defn part2-fuel-cost [distance] (apply + (range 1 (inc distance))))

(defn lowest-fuel-cost [data cost]
  (let [lo (apply min data)
        hi (apply max data)]
    (apply min (for [i (range lo (inc hi))]
                 (apply + (map #(cost (abs (- i %))) data))))))

(defn part1 [data] (lowest-fuel-cost data part1-fuel-cost))
(defn part2 [data] (lowest-fuel-cost data (memoize part2-fuel-cost)))
