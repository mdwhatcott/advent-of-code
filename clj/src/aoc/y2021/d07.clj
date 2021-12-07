(ns aoc.y2021.d07)

(defn abs [n] (if (neg? n) (- n) n))

(defn part1-fuel-cost [distance] distance)

(defn lowest-fuel-cost [data cost]
  (let [lowest  (apply min data)
        highest (apply max data)]
    (apply min (for [i (range lowest (inc highest))]
                 (apply + (map #(cost (abs (- i %))) data))))))

(defn part1 [data]
  (lowest-fuel-cost data part1-fuel-cost))

(defn part2-fuel-cost [distance]
  (apply + (range 1 (inc distance))))

(defn part2 [data]
  (lowest-fuel-cost data (memoize part2-fuel-cost)))
