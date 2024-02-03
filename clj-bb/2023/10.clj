#!/usr/bin/env bb

(require '[clojure.string :as str])

(def up    [ 0 -1])
(def down  [ 0  1])
(def left  [-1  0])
(def right [ 1  0])

(defn travel [])

(defn scan-world [lines]
  (into {}
    (for [y (range (count lines))
          x (range (count (first lines)))
          :let [c (get (get lines y) x)]
          :when (not (= c \.))]
      [[x y] c])))

(defn init [lines [x y] [dx dy]]
  {:world (scan-world lines)
   :start [x y]
   :xy    [x y]
   :dxy   [dx dy]
   :loop  []})

(require '[clojure.test :refer :all])

(def sample-a [
  "....."
  ".F-7."
  ".|.|."
  ".L-J."
  "....."])

(deftest tests
  (is (= {:world {[1 1] \F [2 1] \- [3 1] \7
                  [1 2] \|          [3 2] \|
                  [1 3] \L [2 3] \- [3 3] \J}
          :start [1 1]
          :xy    [1 1]
          :dxy   [1 0]
          :loop  []}
         (init sample-a [1 1] [1 0])))
)

(let [{:keys [fail error]} (run-tests)]
  (if (pos? (+ fail error)) (System/exit 1) (println "\nOK")))