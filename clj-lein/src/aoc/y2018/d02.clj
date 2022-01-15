(ns aoc.y2018.d02)

(defn has-n-repeats? [n input]
  (as-> (sort input) $
        (partition-by char $)
        (filter #(= n (count %)) $)
        (not (empty? $))))

(defn count-repeats [n inputs]
  (count (filter #(has-n-repeats? n %) inputs)))

(defn part1 [inputs]
  (* (count-repeats 2 inputs)
     (count-repeats 3 inputs)))

(def equal? (partial apply =))

(defn diff-count [a b]
  (->> (interleave a b)
       (partition 2)
       (remove equal?)
       (count)))

(defn common-chars [a b]
  (->> (interleave a b)
       (partition 2)
       (filter equal?)
       (map first)
       (apply str)))

(defn part2 [inputs]
  (first
    (for [x inputs
          y inputs
          :when (= (diff-count x y) 1)]
      (common-chars x y))))
