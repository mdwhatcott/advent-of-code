#!/usr/bin/env bb

(require '[clojure.string :as str])

(defn diffs [coll]
  (->> coll (partition 2 1) (map (fn [[a b]] (- b a)))))

(defn diff-until-zeros [coll]
  (->> coll (iterate diffs) (take-while (partial (complement every?) zero?))))

(defn predict-next-in [coll]
  (loop [series (reverse (diff-until-zeros coll))]
    (if (= 1 (count series))
      (last (last series))
      (let [next (+ (last (first series)) (last (second series)))]
        (recur (cons (conj (vec (second series)) next) (drop 2 series)))))))

(defn extrapolation-sum [colls]
  (->> colls (map predict-next-in) (apply +)))

(defn ints [line]
  (->> (str/split line #" ") (map #(Integer/parseInt %))))

(require '[clojure.test :refer :all])

(def sample-a [0 3 6 9 12 15])
(def sample-b [1 3 6 10 15 21])
(def sample-c [10 13 16 21 30 45])
(def samples  [sample-a sample-b sample-c])
(def inputs   (->> "09.txt" slurp str/split-lines (map ints)))

(deftest tests
  (is (= [3 3 3 3 3]  (diffs sample-a)))
  (is (= [2 3 4 5 6]  (diffs sample-b)))
  (is (= [3 3 5 9 15] (diffs sample-c)))
  
  (is (= [sample-a [3 3 3 3 3]]                    (diff-until-zeros sample-a)))
  (is (= [sample-b [2 3 4 5 6] [1 1 1 1]]          (diff-until-zeros sample-b)))
  (is (= [sample-c [3 3 5 9 15] [0 2 4 6] [2 2 2]] (diff-until-zeros sample-c)))
  
  (is (= 18 (predict-next-in sample-a)))
  (is (= 28 (predict-next-in sample-b)))
  (is (= 68 (predict-next-in sample-c)))
  
  (is (= (+ 18 28 68) (->> samples               extrapolation-sum)))  ; part 1
  (is (= (+ 5 -3 0)   (->> samples (map reverse) extrapolation-sum)))  ; part 2
  
  (is (= 1731106378   (->> inputs                extrapolation-sum)))  ; part 1
  (is (= 1087         (->> inputs  (map reverse) extrapolation-sum)))) ; part 2

(let [{:keys [fail error]} (run-tests)]
  (if (pos? (+ fail error)) (System/exit 1) (println "\nOK")))