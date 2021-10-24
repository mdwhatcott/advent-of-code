(ns aoc.y2018.d13-spec
  (:require [speclj.core :refer :all]
            [aoc.y2018.d13 :as sut]
            [clojure.string :as string]
            [clojure.string :refer [trim]]))

(def sample-map-1 (slurp "data/2018/d13/sample-map-1.txt"))

(describe "2018 Day 13"
  (context "Part 1"
    (it "parses the tracks"
      (->> (sut/parse-map sample-map-1)
           (should= {[0 0] "/",, [1 0] "-", [2 0] "-", [3 0] "-", [4 0] "-", [5 0] "\\",
                     [0 1] "|",,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,, [5 1] "|",
                     [0 2] "|",,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,, [5 2] "|",
                     [0 3] "\\", [1 3] "-", [2 3] "-", [3 3] "-", [4 3] "-", [5 3] "/"})))

    (xit "finds the starting position of the carts"
      (->> (sut/find-carts sample-map-1)
           ; := (current location)
           ; :> (current direction)
           ; :+ (next intersection turn)
           (should= [{:= [2 0] :> :r :+ :l}
                     {:= [5 1] :> :d :+ :l}
                     {:= [0 2] :> :u :+ :l}
                     {:= [3 3] :> :l :+ :l}])))
    )

  (context "Part 2"
    )
  )