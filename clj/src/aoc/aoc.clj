(ns aoc.aoc
  (:require
    [clojure.string :as str]
    [clojure.test :as test]))

(defn str->int [s]
  (Integer/parseInt s))

(defn input-string [year day]
  (let [day (if (< day 10) (str "0" day) day)]
    (slurp (str "src/y" year "/d" day ".txt"))))

(defn input-lines [year day]
  (str/split-lines (input-string year day)))

(defn input-ints
  ([year day] (input-ints year day #"\s"))
  ([year day sep]
   (map str->int (str/split (input-string year day) sep))))

(defn exit [result]
  (let [exit-code (+ (:fail result) (:error result))]
    (println (if (zero? exit-code) "OK" "FAIL"))
    (System/exit exit-code)))