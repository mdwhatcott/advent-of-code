(ns aoc.y2018.d13
  (:require [clojure.string :as string]
            [clojure.pprint :as pprint]))

(def R [1 0])
(def L [-1 0])
(def D [0 1])
(def U [0 -1])

(def S [0 0])

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

(defn advance [cart]
  (let [[x1 y1] (:= cart)
        [dx dy] (:> cart)]
    [(+ x1 dx)
     (+ y1 dy)]))

(defn turn [turn-direction current-direction]
  (condp = turn-direction
    S current-direction
    L (condp = current-direction U L, L D, D R, R U)
    R (condp = current-direction R D, L U, D L, U R)))

(defn orient [cart track]
  (let [direction (:> cart)]
    (case track
      "|", direction
      "-", direction
      "+", (turn (:+ cart) direction)
      "/", (condp = direction R U, L D, U R, D L)
      "\\" (condp = direction R D, L U, U L, D R))))

(defn intersection [cart track]
  (if-not (= track "+")
    (:+ cart)
    (condp = (:+ cart) L S, S R, R L)))

(defn move [cart tracks]
  (let [target         (advance cart)
        track          (get tracks target)
        direction      (orient cart track)
        turn-direction (intersection cart track)]
    (assoc cart := target
                :> direction
                :+ turn-direction)))

(defn carts-map [carts]
  (into {} (for [cart carts] [(:= cart) cart])))

(defn tick [world]
  (loop [all        (carts-map (:carts world))
         ordered    (sort-by sorting-order (:carts world))
         collisions []]
    (if (empty? ordered)
      (assoc world :carts (vals all)
                   :collisions collisions)
      (let [original (:= (first ordered))]
        (if-not (contains? all original)
          (recur all (rest ordered) collisions)
          (let [moved      (move (first ordered) (:tracks world))
                target     (:= moved)
                collision? (contains? all target)]

            (if collision?
              (recur (dissoc all original target)
                     (rest ordered)
                     (conj collisions target))
              (recur (-> all (assoc target moved) (dissoc original))
                     (rest ordered)
                     collisions)))))))
  )

(defn until-first-collision [world]
  (as-> (iterate tick world) $
        (drop-while #(empty? (:collisions %)) $)
        (first $)
        (:collisions $)
        (first $)))

(defn until-last-cart-remains [world]
  (as-> (iterate tick world) $
        (drop-while #(> (count (:carts %)) 1) $)
        (first $)
        (:carts $)
        (first $)
        (:= $)))

