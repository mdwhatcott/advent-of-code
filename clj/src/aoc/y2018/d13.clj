(ns aoc.y2018.d13
  (:require [clojure.string :as string]
            [clojure.pprint :as pprint]))

(defn parse-map-line [[y line]]
  (as-> (range (count line)) $
        (for [x $
              :let [c (str (nth line x))]
              :when (not= c " ")]
          [[x y] c])))

(defn parse-map [raw-map]
  (as-> raw-map $
        (string/replace $ "v" "|")
        (string/replace $ "^" "|")
        (string/replace $ ">" "-")
        (string/replace $ "<" "-")
        (string/split-lines $)
        (interleave (range (count $)) $)
        (partition 2 $)
        (mapcat parse-map-line $)
        (into (sorted-map) $)))

(defn find-carts [raw-map]
  )
