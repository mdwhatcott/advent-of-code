(ns aoc.y2021.d04
  (:require [aoc.data :as data]
            [clojure.string :as string]
            [clojure.set :as set]))

(defn parse-input [lines]
  {:call-order (as-> (first lines) $
                     (string/split $ #",")
                     (map data/str->int $))
   :boards     (as-> (rest lines) $
                     (partition 6 $)
                     (map #(drop 1 %) $))})

(defn- board-numbers [lines]
  (as-> lines $
        (map #(str % " ") $)
        (apply str $)
        (string/split $ #"\s")
        (remove empty? $)
        (map data/str->int $)))

(defn row-sets [numbers]
  (map set (partition 5 numbers)))

(defn nth-nums [col rows]
  (map #(nth (seq %) col) rows))

(defn col-sets [numbers]
  (let [rows (partition 5 numbers)]
    (as-> (range (count rows)) $
          (map #(nth-nums % rows) $)
          (map set $))))

(defn prepare-board [call-order lines]
  (let [numbers (board-numbers lines)]
    {:score      0
     :checksum   0
     :call-order call-order
     :called     []
     :marked     #{}
     :unmarked   (set numbers)
     :wins       (concat (row-sets numbers)
                         (col-sets numbers))}))

(defn call-next-number [board]
  (let [num   (first (:call-order board))
        board (-> board
                  (update :call-order rest)
                  (update :called conj num))]
    (if-not (contains? (:unmarked board) num)
      board
      (-> board
          (update :marked conj num)
          (update :unmarked disj num)))))

(defn- bingo? [{:keys [marked wins]}]
  (some? (first (filter #(set/subset? % marked) wins))))

(defn score [board]
  (let [board (call-next-number board)]
    (if-not (bingo? board)
      board
      (let [score  (apply + (:unmarked board))
            called (peek (:called board))]
        (assoc board :score score
                     :checksum (* score called))))))

(defn play [board]
  (as-> (iterate score board) $
        (drop-while #(and (some? (:call-order %))
                          (zero? (:score %))) $)
        (first $)))

(defn find-winner [data comparison]
  (let [parsed  (parse-input data)
        boards  (map #(prepare-board (:call-order parsed) %)
                     (:boards parsed))
        played  (map play boards)
        winners (filter #(pos? (:score %)) played)]
    (:checksum (first (sort-by #(count (:called %)) comparison winners)))))

(defn part1 [data]
  (find-winner data <))

(defn part2 [data]
  (find-winner data >))
