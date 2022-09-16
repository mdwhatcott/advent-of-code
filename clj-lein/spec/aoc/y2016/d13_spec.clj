(ns aoc.y2016.d13-spec
  (:require [speclj.core :refer :all]))

; Search Algorithms Demo: https://www.redblobgames.com/pathfinding/a-star/introduction.html
; Hamming Weight Explanation: https://en.wikipedia.org/wiki/Hamming_weight
; Hamming Weight Demo: https://go.dev/play/p/Uf9xmV2pMcs

(defn bits [n]
  (->> n
       (iterate #(bit-and % (dec %)))
       (take-while pos?)
       count))

(defn hallway? [seed [x y]]
  (and (not (neg? x))
       (not (neg? y))
       (->> seed
            (+ (* x x) (* 3 x) (* 2 x y) y (* y y))
            bits
            even?)))

(defn bfs [origin target graph]
  (loop [seen #{} frontier [[origin 0]]]
    (let [[at dist] (peek frontier)
          frontier (pop frontier)]
      (if (= at target)
        dist
        (let [neighbors (remove seen (graph at))
              steps     (for [n neighbors] [n (inc dist)])]
          (recur (conj seen at)
                 (vec (concat steps frontier))))))))

(defn cardinal-neighbors [[x y]]
  [[(dec x) y]
   [(inc x) y]
   [x (dec y)]
   [x (inc y)]])

(defn maze-neighbors [seed point]
  (filter (partial hallway? seed) (cardinal-neighbors point)))

(describe "2016 Day 13"
  (it "Part 1"
    (should= 11 (bfs [1 1] [7 4] (partial maze-neighbors 10)))
    (should= 96 (bfs [1 1] [31 39] (partial maze-neighbors 1358))))

  (it "maze neighbors"
    (should= (cardinal-neighbors [7 1]) (maze-neighbors 10 [7 1]))
    (should= [] (maze-neighbors 10 [5 3]))
    (should= [[0 1] [1 2]] (maze-neighbors 10 [1 1])))

  (it "breadth-first search"
    (should= 0 (bfs :a :a {:a []}))
    (should= 1 (bfs :a :b {:a [:b]}))
    (should= 2 (bfs :a :c {:a [:b] :b [:c]}))
    (should= 9 (bfs [1 1] [7 4] cardinal-neighbors)))

  (it "cardinal neighbors"
    (should= [[0 1] [2 1] [1 0] [1 2]] (cardinal-neighbors [1 1])))

  (it "walls vs halls"
    (should= false (hallway? 10 [0 -1]))
    (should= false (hallway? 10 [-1 0]))
    (should= true (hallway? 10 [0 0]))
    (should= true (hallway? 10 [1 1]))
    (should= false (hallway? 10 [9 6]))
    (should= false (hallway? 10 [1 0])))

  (it "counting '1' bits"
    (should= 0 (bits 0)) #_"0000"
    (should= 1 (bits 1)) #_"0001"
    (should= 1 (bits 2)) #_"0010"
    (should= 1 (bits 4)) #_"0100"
    (should= 2 (bits 3)) #_"0011"
    (should= 8 (bits 255))))