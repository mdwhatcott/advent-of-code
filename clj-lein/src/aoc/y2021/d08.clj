(ns aoc.y2021.d08
  (:require [clojure.string :as string]
            [clojure.set :refer [superset?]]))

(defn count-unique-digits [lines]
  (as-> lines $
        (map #(string/split % #"\s\|\s") $)
        (map last $)
        (string/join " " $)
        (string/split $ #"\s")
        (map count $)
        (frequencies $)
        (+ (get $ 2)                                        ; 1
           (get $ 4)                                        ; 4
           (get $ 3)                                        ; 7
           (get $ 7)                                        ; 8
           )))

(defn part1 [data]
  (count-unique-digits data))

(defn chars [n line] (= n (count line)))

(defn decipher-digits [line]
  (let [options (as-> line $
                      (string/split $ #"\s\|\s") (first $)
                      (string/split $ #"\s") (map set $))]
    (let [one   (first (filter #(chars 2 %) options))
          seven (first (filter #(chars 3 %) options))
          four  (first (filter #(chars 4 %) options))
          eight (first (filter #(chars 7 %) options))
          nine  (first (filter #(and (chars 6 %1) (superset? %1 four)) options))
          three (first (filter #(and (chars 5 %1) (superset? %1 seven)) options))
          five  (first (filter #(and (chars 5 %1) (superset? nine %1) (not= %1 three)) options))
          six   (first (filter #(and (chars 6 %1) (superset? %1 five) (not= %1 nine)) options))
          zero  (first (filter #(and (chars 6 %1) (not= %1 six) (not= %1 nine)) options))
          two   (first (filter #(and (chars 5 %1) (not= %1 five) (not= %1 three)) options))]
      {zero 0 one 1 two 2 three 3 four 4 five 5 six 6 seven 7 eight 8 nine 9})))

(defn reveal-code [line]
  (let [digits  (decipher-digits line)
        options (as-> line $
                      (string/split $ #"\s\|\s") (last $)
                      (string/split $ #"\s") (map set $))]
    (as-> options $
          (map digits $)
          (apply str $)
          (Integer/parseInt $))))

(defn part2 [data]
  (->> data
       (map reveal-code)
       (apply +)))
