(ns y2020.d09
  (:require [aoc.aoc :as aoc]
            [clojure.test :as test :refer [deftest is]]))

(defn valid? [window]
  (let [preamble (set (drop-last window))
        sum      (last window)
        checks   (for [key preamble
                       :let [remainder (- sum key)]
                       :when (not= key remainder)]
                   (contains? preamble remainder))]
    (->> checks (filter true?) first some?)))

(defn part1 [inputs window-size]
  (->> inputs
       (partition (inc window-size) 1)
       (drop-while #(valid? %))
       first
       last))

(def sample-data
  [35 20 15 25 47 40 62 55 65 95 102
   117 150 182 127 219 299 277 309 576])

(deftest valid-sums
  (is (not (valid? [1 2 10])))
  (is (valid? [1 2 3]))
  (is (not (valid? [1 2 4]))))

(deftest part-1
  (is (= 127 (part1 sample-data 5)))
  (is (= 23278925 (part1 (aoc/input-ints 2020 9) 25))))

(aoc/exit (test/run-tests))
