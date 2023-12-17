#!/usr/bin/env bb

(require '[clojure.test :refer :all])
(require '[clojure.string :as str])

(def left  \( )
(def right \) )

(defn next [at direction]
  (if (= direction left) (inc at) (dec at)))

(defn ending-floor [input]
  (->> input (reductions next 0) last))

(defn first-enters-basement-at [input]
  (->> input (reductions next 0) (take-while (complement neg?)) count))

(def actual-input (->> "01.txt" slurp))

(deftest part1-answer
  (are [input expected] (= expected (ending-floor input))
    "(())"     0
    "()()"     0
    "((("      3
    "(()(()("  3
    "))((((("  3
    "())"     -1
    "))("     -1
    ")))"     -3
    ")())())" -3)
  
  (is (= 232 (ending-floor actual-input))))

(deftest part2-answer
  (is (= 5 (first-enters-basement-at "()())")))
  (is (= 1783 (first-enters-basement-at actual-input))))

(let [{:keys [fail error]} (run-tests)]
  (if (pos? (+ fail error)) (System/exit 1) (println "\nOK")))