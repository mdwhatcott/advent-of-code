(ns aoc.y2018.d05
  (:require [clojure.string :as string]))

(defn abs [n]
  (max n (- n)))

(def lower-vs-upper
  (- (int \a)
     (int \A)))

(defn unstable? [a b]
  (= lower-vs-upper
     (abs (- (int (or a 0))
             (int (or b 0))))))

(defn react [output i]
  (if (unstable? (peek output) i)
    (pop output)
    (conj output i)))

(defn part1 [input]
  (count (reduce react [] input)))

(defn exclude-pair [input i]
  (let [lower (str (char i))
        upper (str (char (- i lower-vs-upper)))]
    (-> input
        (string/replace lower "")
        (string/replace upper ""))))

(defn part2 [input]
  (->> (range (int \a) (inc (int \z)))
       (map #(exclude-pair input %))
       (remove #(= (count %) (count input)))
       (map #(part1 %))
       (apply min)))

(defn part1-recur-fast [INPUT]
  (loop [input  (rest INPUT)
         output [(first INPUT)]]
    (if (empty? input)
      (count output)
      (let [i (first input)]
        (if (unstable? (peek output) i)
          (recur (rest input) (pop output))
          (recur (rest input) (conj output i)))))))
