(ns aoc.y2021.d05
  (:require [aoc.data :as data]))

(defn parse-endpoints [lines diagonals?]
  (as-> (map #(re-matches #"(\d+),(\d+) -> (\d+),(\d+)" %) lines) $
        (map rest $)
        (flatten $)
        (map data/str->int $)
        (partition 4 $)
        (for [[x1 y1 x2 y2] $
              :when (or diagonals? (= x1 x2) (= y1 y2))]
          {:from [x1 y1] :to [x2 y2]})))

(defn abs [n]
  (if (neg? n) (- n) n))

(defn draw-line [{:keys [from to]}]
  (let [[x1 y1] from
        [x2 y2] to
        run  (- x2 x1)
        rise (- y2 y1)
        n    (if (zero? rise) (abs run) (abs rise))
        run  (cond (zero? run) 0 (pos? run) 1 (neg? run) -1)
        rise (cond (zero? rise) 0 (pos? rise) 1 (neg? rise) -1)]
    (for [a (range (inc n))]
      [(+ x1 (* run a))
       (+ y1 (* rise a))])))

(defn count-overlaps [data diagonals?]
  (as-> (parse-endpoints data diagonals?) $
        (mapcat draw-line $)
        (frequencies $)
        (remove #(= 1 (second %)) $)
        (count $)))

(defn part1 [data]
  (count-overlaps data false))

(defn part2 [data]
  (count-overlaps data true))
