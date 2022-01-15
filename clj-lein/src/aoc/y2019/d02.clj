(ns aoc.y2019.d02
  (:require
    [aoc.y2019.intcode :as intcode]))

(defn run [memory noun verb]
  (let [tweaked  (assoc memory, 1 noun, 2 verb)
        executed (intcode/run-simple tweaked)]
    (first executed)))

(defn part1 [memory] (run memory 12 2))
(defn part2 [memory]
  (first
    (for [noun (range 100)
          verb (range 100)
          :let [output (run memory noun verb)]
          :when (= output 19690720)]
      (+ verb (* 100 noun)))))
