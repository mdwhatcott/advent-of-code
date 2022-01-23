(ns y2021.d14
  (:require [aoc.aoc :as aoc]
            [clojure.test :refer :all]
            [clojure.pprint]
            [clojure.string :as str]))

(defn reactor [[a c b]]
  (let [a       (first a)
        b       (first b)
        c       (first c)
        compare [a c]
        out     [b c]]
    (fn [in] (if (not= in compare) nil out))))

(defn make-reactors [input]
  (->> input
       (map #(re-matches #"(.)(.) -> (.)" %))
       (map rest)
       (map reactor)))

(defn react-each [reactions polymer]
  (as-> reactions $
        (map #(% polymer) $)
        (remove nil? $)
        (first $)
        (or $ polymer)))

(defn react [reactions polymer]
  (->> (seq polymer)
       (partition 2 1)
       (map (partial react-each reactions))
       (cons (first polymer))
       (flatten)))

(defn most-minus-least [rounds reactions polymer]
  (let [reactor (partial react reactions)
        final   (first (drop rounds (iterate reactor polymer)))
        freqs   (vals (frequencies final))]
    (- (apply max freqs)
       (apply min freqs))))

(def real-input (aoc/input-lines 2021 14))
(def real-starting-polymer (first real-input))
(def real-reactions (make-reactors (drop 2 real-input)))

(deftest parses-input
  (is (= "ONHOOSCKBSVHBNKFKSBK" real-starting-polymer))
  (is (= 100 (count real-reactions)))
  (is (= [\B \O] ((first real-reactions) [\H \O]))))

(def sample-starting-polymer "NNCB")
(def sample-reactions
  (make-reactors ["CH -> B"
                  "HH -> N"
                  "CB -> H"
                  "NH -> C"
                  "HB -> C"
                  "HC -> B"
                  "HN -> C"
                  "NN -> C"
                  "BH -> H"
                  "NC -> B"
                  "NB -> B"
                  "BN -> B"
                  "BB -> N"
                  "BC -> B"
                  "CC -> N"
                  "CN -> C"]))

(deftest iterating-rounds-with-sample-data
  (is (= (seq "NCNBCHB") (react sample-reactions sample-starting-polymer)))
  (is (= (seq "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB")
         (first (drop 4 (iterate (partial react sample-reactions) sample-starting-polymer)))))
  (is (= {\B 1749 \C 298 \H 161 \N 865}
         (frequencies (first (drop 10 (iterate (partial react sample-reactions) sample-starting-polymer)))))))

(deftest part-1
  (is (= 1588 (most-minus-least 10 sample-reactions sample-starting-polymer)))
  (is (= 2447 (most-minus-least 10 real-reactions real-starting-polymer))))

(aoc/run-tests)

