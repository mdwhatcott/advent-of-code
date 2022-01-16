(ns y2021.d11
  (:require [aoc.aoc :as aoc]
            [clojure.test :refer :all]
            [clojure.pprint]))

(defn square-width [coll]
  (int (Math/sqrt (count coll))))

(defn parse-board [nums]
  (let [width (square-width nums)]
    (->> (for [y (range width)
               x (range width)
               :let [i (+ x (* y width))
                     v (nth nums i)]]
           {[x y] v})
         (into {}))))

(defn neighbors8 [[x y]]
  (for [Y (range (dec y) (inc (inc y)))
        X (range (dec x) (inc (inc x)))
        :when (not= [X Y] [x y])] [X Y]))

(defn neighbors [board point]
  (sort (filter board (neighbors8 point))))

(defn parse-raw [input]
  (->> input (remove #{\newline})
       (map str) (map aoc/str->int)
       parse-board))

(defn initialize [input]
  {:flashes 0 :new-flashes 0 :board (parse-raw input)})

(defn charge [board]
  (into {} (for [[point value] board]
             [point (inc value)])))

(defn reset [board]
  (into {} (for [[point value] board]
             [point (or value 0)])))

(defn at-capacity? [[_point value]]
  (and (some? value) (>= value 10)))

(defn inc-points [board points]
  (if (empty? points)
    board
    (if (nil? (board (first points)))
      (recur board (rest points))
      (recur (update board (first points) inc) (rest points)))))

(defn mark-flashers [board points]
  (if (empty? points)
    board
    (recur (assoc board (first points) nil) (rest points))))

(defn flash [board]
  (let [to-flash (map first (filter at-capacity? board))]
    (if (empty? to-flash)
      board
      (let [neighbors (partial neighbors board)
            neighbors (mapcat neighbors to-flash)
            marked    (mark-flashers board to-flash)
            boosted   (inc-points marked neighbors)]
        (recur boosted)))))

(defn pretty-board [board]
  (with-out-str
    (clojure.pprint/pprint
      (let [width (square-width board)]
        (->> (for [y (range width)
                   x (range width)]
               (get board [x y]))
             (partition width))))))

(defn step [state]
  (let [board   (:board state)
        charged (charge board)
        flashed (flash charged)
        final   (reset flashed)
        flashes (count (filter nil? (vals flashed)))]
    #_(println (pretty-board final))
    (-> state
        (assoc :board final :new-flashes flashes)
        (update :flashes + flashes))))

(defn count-flashes [input steps]
  (->> (initialize input)
       (iterate step)
       (drop steps)
       first
       :flashes))

(defn all-flashed? [state]
  (= (:new-flashes state)
     (count (:board state))))

(defn first-synchronized-flash [input]
  (->> (initialize input)
       (iterate step)
       (take-while (complement all-flashed?))
       count))

(def sample-data
  (str "5483143223\n"
       "2745854711\n"
       "5264556173\n"
       "6141336146\n"
       "6357385478\n"
       "4167524645\n"
       "2176841721\n"
       "6882881134\n"
       "4846848554\n"
       "5283751526"))

(def real-data (aoc/input-string 2021 11))

(deftest part-1
  (is (= 1691 (count-flashes real-data 100))))

(deftest part-2
  (is (= 216 (first-synchronized-flash real-data))))

(deftest it-finds-the-first-synchronized-flash
  (is (= 195 (first-synchronized-flash sample-data))))

(deftest it-counts-flashes-over-many-steps
  (is (= 35 (count-flashes sample-data 2)))
  (is (= 204 (count-flashes sample-data 10)))
  (is (= 1656 (count-flashes sample-data 100))))

(deftest it-takes-a-step-with-multiple-flashes
  (is (= {:flashes     9
          :new-flashes 9
          :board       (parse-raw (str "34543\n40004\n50005\n40004\n34543"))}
         (step {:flashes     0
                :new-flashes 0
                :board       (parse-raw (str "11111\n19991\n19191\n19991\n11111"))}))))

(deftest it-takes-a-step-with-a-single-flash
  (is (= {:flashes     1
          :new-flashes 1
          :board       (parse-raw (str "333\n"
                                       "303\n"
                                       "333"))}
         (step {:flashes     0
                :new-flashes 0
                :board       (parse-raw (str "111\n"
                                             "191\n"
                                             "111"))}))))
(deftest it-takes-a-step
  (is (= {:flashes     0
          :new-flashes 0
          :board       (parse-raw (str "45654\n"
                                       "51115\n"
                                       "61116\n"
                                       "51115\n"
                                       "45654"))}
         (step {:flashes     0
                :new-flashes 0
                :board       (parse-raw (str "34543\n"
                                             "40004\n"
                                             "50005\n"
                                             "40004\n"
                                             "34543"))}))))

(deftest it-initializes-state
  (is (= {:board       (parse-board (range 1 10))
          :flashes     0
          :new-flashes 0}
         (initialize "123\n456\n789"))))

(deftest it-parses-raw-input
  (is (= (parse-board (range 1 10))
         (parse-raw "123\n456\n789"))))

(deftest it-identifies-neighbors
  (is (= [[0 0] [1 0] [2 0]
          [0 1],,,,,, [2 1]
          [0 2] [1 2] [2 2]]
         (neighbors8 [1 1])))
  (is (= [[0 1] [1 0] [1 1]]
         (neighbors (parse-board (range 1 10)) [0 0]))))

(deftest it-parses-board
  (is (= {[0 0] 1 [1 0] 2 [2 0] 3
          [0 1] 4 [1 1] 5 [2 1] 6
          [0 2] 7 [1 2] 8 [2 2] 9}
         (parse-board [1 2 3
                       4 5 6
                       7 8 9]))))

(aoc/exit (run-tests))