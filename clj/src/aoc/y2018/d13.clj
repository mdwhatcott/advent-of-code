(ns aoc.y2018.d13
  (:require [clojure.string :as string]
            [clojure.pprint :as pprint]))

(def R [1 0])
(def L [-1 0])
(def D [0 1])
(def U [0 -1])

(def cart->d
  {"<" L
   ">" R
   "v" D
   "^" U})

(defn widest-line [lines]
  (apply max (map count lines)))

(defn parse-map-lines [lines]
  (for [y (range (count lines))
        x (range (widest-line lines))
        :let [line (nth lines y)
              char (str (get line x ""))]
        :when (not= char " ")]
    [[x y] char]))

(defn parse-tracks [raw-map]
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
     :> (get cart->d char)
     :+ L}))

(defn find-carts [raw-map]
  (parse-carts (string/split-lines raw-map)))

(defn parse-initial [raw-map]
  {:tracks (parse-tracks raw-map)
   :carts  (find-carts raw-map)})

(defn sorting-order [cart]
  (let [[x y] (:= cart)] [y x]))

(defn move [cart tracks]
  (let [[x1 y1] (:= cart)
        [dx dy] (:> cart)
        target [(+ x1 dx)
                (+ y1 dy)]]
    (assoc cart := target)))

(defn tick [world]
  (loop [carts  (:carts world)
         carts2 {}]
    (let [carts (sort-by sorting-order carts)]
      (if (empty? carts)
        (assoc world :carts (vals carts2))
        (let [cart2  (move (first carts) (:tracks world))
              target (:= cart2)
              collision? (contains? carts2 target)]
          (if collision?
            (recur (rest carts) (dissoc carts2 target))
            (recur (rest carts) (assoc carts2 target cart2)))))))
  )
