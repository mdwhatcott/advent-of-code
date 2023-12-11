#!/usr/bin/env bb

(require '[clojure.test :refer :all])
(require '[clojure.string :as str])

(def actual-input
  (->> "01.txt" slurp str/split-lines))
(def sample-input-part1 [
  "1abc2"
  "pqr3stu8vwx"
  "a1b2c3d4e5f"
  "treb7uchet"])
(def sample-input-part2 [
  "two1nine"
  "eightwothree"
  "abcone2threexyz"
  "xtwone3four"
  "4nineeightseven2"
  "zoneight234"
  "7pqrstsixteen"])

(def part1-graph {
  \1 1 \2 2 \3 3
  \4 4 \5 5 \6 6
  \7 7 \8 8 \9 9})

(defn insert-word [graph [letters number]]
  (assoc-in graph letters number))

(defn make-word-graph [words2digits] 
  (reduce insert-word part1-graph words2digits))

(def part2-graph
  (make-word-graph {
    "one"   1 
    "two"   2 
    "three" 3 
    "four"  4 
    "five"  5 
    "six"   6 
    "seven" 7 
    "eight" 8 
    "nine"  9}))

(defn lookup-digit [graph prefix]
  (let [lookup (get graph (first prefix))]
    (if (or (nil? lookup) (number? lookup))
        lookup
        (recur lookup (subs prefix 1)))))

(defn make-substrings [s]
  (map (partial subs s) (range (count s))))

(defn inventory-digits [graph s]
  (remove nil? (map (partial lookup-digit graph) (make-substrings s))))

(defn calibration-value [digits]
  (+ (* (first digits) 10) (last digits)))

(defn solve [graph lines]
  (->> lines
       (map (partial inventory-digits graph)) 
       (map calibration-value) 
       (flatten) 
       (apply +)))

(deftest part1-answer 
  (is (= 142 (solve part1-graph sample-input-part1)))
  (is (= 55538 (solve part1-graph actual-input))))

(deftest part2-answer
  (is (= 281 (solve part2-graph sample-input-part2)))
  (is (= 54875 (solve part2-graph actual-input))))

(deftest test-inventory-digits
  (is (= [2 1 9] (inventory-digits part2-graph "two1nine")))
  (is (= [8 2 3] (inventory-digits part2-graph "eightwothree")))
  (is (= [1 2 3] (inventory-digits part2-graph "abcone2threexyz")))
  (is (= [2 1 3 4] (inventory-digits part2-graph "xtwone3four")))
  (is (= [4 9 8 7 2] (inventory-digits part2-graph "4nineeightseven2")))
  (is (= [7 6] (inventory-digits part2-graph "7pqrstsixteen"))))

(deftest test-make-substrings
  (is (= ["abc" "bc" "c"] (make-substrings "abc"))))

(let [{:keys [fail error]} (run-tests)]
  (when (pos? (+ fail error))
    (System/exit 1)))
