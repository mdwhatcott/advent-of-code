(ns y2021.d13
  (:require [aoc.aoc :as aoc]
            [clojure.test :refer :all]
            [clojure.pprint]
            [clojure.string :as str]))

(defn parse-dot [line]
  (map aoc/str->int (rest (re-matches #"(\d+),(\d+)" line))))

(defn parse-all-dots [lines]
  (set (map parse-dot (take-while (complement str/blank?) lines))))

(defn fold-up [[x y] at]
  (if (< y at) [x y] [x (- y (* 2 (- y at)))]))

(defn fold-left [dot at]
  (reverse (fold-up (reverse dot) at)))

(defn parse-fold [line]
  (let [[axis at] (rest (re-matches #"fold along (.)=(\d+)" line))]
    [(if (= axis "x") fold-left fold-up) (aoc/str->int at)]))

(defn parse-all-folds [lines]
  (map parse-fold (rest (drop-while (complement str/blank?) lines))))

(defn fold [dot [folder at]]
  (folder dot at))

(defn fold-dots [dots instruction]
  (set (map #(fold % instruction) dots)))

(defn fold-all [dots instructions]
  (loop [dots dots instructions instructions]
    (if (empty? instructions)
      dots (recur (fold-dots dots (first instructions))
                  (rest instructions)))))

(def sample-lines
  ["6,10"
   "0,14"
   "9,10"
   "0,3"
   "10,4"
   "4,11"
   "6,0"
   "6,12"
   "4,1"
   "0,13"
   "10,12"
   "3,4"
   "3,0"
   "8,4"
   "1,10"
   "2,14"
   "8,10"
   "9,0"
   ""
   "fold along y=7"
   "fold along x=5"])
(def sample-dots (parse-all-dots sample-lines))
(def sample-folds (parse-all-folds sample-lines))

(deftest parses-dots
  (is (= [6 10] (parse-dot "6,10")))
  (is (= 18 (count sample-dots))))

(deftest parses-folds
  (is (= [fold-up 7] (parse-fold "fold along y=7")))
  (is (= 2 (count sample-folds))))

(deftest folds
  (is (= [3 0] (fold [3 0] [fold-up 7])))
  (is (= [0 0] (fold [0 14] [fold-up 7])))
  (let [dots sample-dots]
    (is (= 17 (count (fold-dots dots [fold-up 7]))))))

(def real-lines (aoc/input-lines 2021 13))
(def real-dots (parse-all-dots real-lines))
(def real-folds (parse-all-folds real-lines))

(deftest part1
  (is (= 765 (count (fold-dots real-dots (first real-folds))))))

(defn render-dots [dots]
  (let [max-rows (apply max (map second dots))
        max-cols (apply max (map first dots))]
    (doseq [y (range (inc max-rows))
            x (range (inc max-cols))
            :let [out (if (contains? dots [x y]) "#" " ")]]
      (if (zero? x)
        (do (println) (print out))
        (print out)))
    (println)))

(deftest part1
  (render-dots (fold-all real-dots real-folds)))            ; RZKZLPGH

; ###  #### #  # #### #    ###   ##  #  #
; #  #    # # #     # #    #  # #  # #  #
; #  #   #  ##     #  #    #  # #    ####
; ###   #   # #   #   #    ###  # ## #  #
; # #  #    # #  #    #    #    #  # #  #
; #  # #### #  # #### #### #     ### #  #

(aoc/exit (run-tests))
