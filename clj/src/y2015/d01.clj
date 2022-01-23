(ns y2015.d01
  (:require [aoc.aoc :as aoc]))

(def input (slurp "src/y2015/d01.txt"))

(defn count-char [char input]
  (count (filter #(= char %) input)))

(defn walk [input]
  (- (count-char \( input)
     (count-char \) input)))

;; super inefficient!
;; (walks incrementally larger sub-
;; strings until we arrive at -1)
(defn walk-until [low-point input]
  (->> (range 1 (inc (count input)))
       (map #(subs input 0 %))
       (map walk)
       (take-while #(> % low-point))
       (count)
       (inc)))

(def up \()
(defn step [char]
  (if (= char up) 1 -1))

(defn part2 [input cursor floor]
  (if (= floor -1)
    cursor
    (let [next-input  (drop 1 input)
          next-cursor (inc cursor)
          next-floor  (+ floor (step (first input)))]
      (part2 next-input next-cursor next-floor))))

(require '[clojure.test :refer :all])

(deftest day1
  (testing "part1"
    (is (= 232 (walk input))))

  (testing "part2-slow"
    (is (= 1783 (walk-until -1 input))))

  (testing "part2-fast")
  (is (= 1783 (part2 (seq input) 0 0))))

(aoc/run-tests)