(ns aoc.y2018.d13
  (:require [clojure.string :as string]
            [clojure.pprint :as pprint]))

(defn widest-line [lines]
  (apply max (map count lines)))

(defn parse-map-lines [lines]
  (for [y (range (count lines))
        x (range (widest-line lines))
        :let [line (nth lines y)
              char (str (get line x ""))]
        :when (not= char " ")]
    [[x y] char]))

(defn parse-map [raw-map]
  (as-> raw-map $
        (string/replace $ "v" "|")
        (string/replace $ "^" "|")
        (string/replace $ ">" "-")
        (string/replace $ "<" "-")
        (string/split-lines $)
        (parse-map-lines $)
        (into {} $)))

(defn parse-carts [lines]
  (for [y (range (count lines))
        x (range (widest-line lines))
        :let [line (nth lines y)
              char (str (get line x ""))]
        :when (#{"<" ">" "v" "^"} char)]
    {:= [x y]
     :> char
     :+ "<"}))

(defn find-carts [raw-map]
  (parse-carts (string/split-lines raw-map)))
