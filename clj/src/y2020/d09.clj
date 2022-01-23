(ns y2020.d09
  (:require [aoc.aoc :as aoc]
            [clojure.test :as test :refer [deftest testing is]]))

(defn valid? [window]
  (let [preamble (set (drop-last window))
        sum      (last window)
        checks   (for [key preamble
                       :let [remainder (- sum key)]
                       :when (not= key remainder)]
                   (contains? preamble remainder))]
    (->> checks (filter true?) first some?)))

(defn part1 [inputs window-size]
  (->> (partition (inc window-size) 1 inputs)
       (drop-while #(valid? %))
       first
       last))

(defn expand-right [{:keys [nums to sum] :as state}]
  (assoc state :sum (+ sum (nth nums to)) :to (inc to)))

(defn contract-left [{:keys [nums from sum] :as state}]
  (assoc state :sum (- sum (nth nums from)) :from (inc from)))

(defn step [{:keys [target sum] :as state}]
  (if (< sum target)
    (expand-right state)
    (contract-left state)))

(defn part2 [inputs window-size]
  (let [target     (part1 inputs window-size)
        initial    {:nums inputs :target target :from 0 :to 0 :sum 0}
        final      (->> (iterate step initial)
                        (drop-while #(not= (:sum %) (:target %))) first)
        contiguous (->> (range (:from final) (:to final))
                        (map (partial nth inputs)))]
    (+ (apply min contiguous)
       (apply max contiguous))))

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

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

(deftest taking-steps
  (let [expected [{:nums sample-data :target 127 :from 0 :to 0 :sum 0}
                  {:nums sample-data :target 127 :from 0 :to 1 :sum 35}
                  {:nums sample-data :target 127 :from 0 :to 2 :sum 55}
                  {:nums sample-data :target 127 :from 0 :to 3 :sum 70}
                  {:nums sample-data :target 127 :from 0 :to 4 :sum 95}
                  {:nums sample-data :target 127 :from 0 :to 5 :sum 142}
                  {:nums sample-data :target 127 :from 1 :to 5 :sum 107}
                  {:nums sample-data :target 127 :from 1 :to 6 :sum 147}
                  {:nums sample-data :target 127 :from 2 :to 6 :sum 127}]]
    (is (= expected (->> (first expected)
                         (iterate step)
                         (take (count expected)))))))
(deftest part-2
  (is (= 62 (part2 sample-data 5)))
  (is (= 4011064 (part2 (aoc/input-ints 2020 9) 25))))

(aoc/run-tests)
