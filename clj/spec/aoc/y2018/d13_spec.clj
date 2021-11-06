(ns aoc.y2018.d13-spec
  (:require [speclj.core :refer :all]
            [aoc.y2018.d13 :as sut]
            [aoc.y2018.d13 :refer [L R U D]]
            [clojure.string :as string]
            [clojure.string :refer [trim]]
            [clojure.pprint :as pprint]))

(def sample-tracks-1 (slurp "data/2018/d13/sample-tracks-1.txt"))
(def sample-tracks-2 (slurp "data/2018/d13/sample-tracks-2.txt"))
(def sample-tracks-3a (slurp "data/2018/d13/sample-tracks-3-a.txt"))
(def sample-tracks-3b (slurp "data/2018/d13/sample-tracks-3-b.txt"))
(def sample-tracks-3c (slurp "data/2018/d13/sample-tracks-3-c.txt"))
(def sample-tracks-3d (slurp "data/2018/d13/sample-tracks-3-d.txt"))
(def sample-tracks-3e (slurp "data/2018/d13/sample-tracks-3-e.txt"))
(def sample-tracks-4 (slurp "data/2018/d13/sample-tracks-4.txt"))

(describe "2018 Day 13"
  (context "Part 1"
    (it "parses the tracks"
      (->> (sut/parse-tracks sample-tracks-1)
           (should= {[0 0] "/",, [1 0] "-", [2 0] "-", [3 0] "-", [4 0] "-", [5 0] "\\",
                     [0 1] "|",,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,, [5 1] "|",
                     [0 2] "|",,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,, [5 2] "|",
                     [0 3] "\\", [1 3] "-", [2 3] "-", [3 3] "-", [4 3] "-", [5 3] "/"})))

    (it "finds the starting position of the carts"
      (->> (sut/find-carts sample-tracks-1)
           ; := (current location)
           ; :> (current direction)
           ; :+ (next intersection turn)
           (should= [{:= [2 0] :> R :+ L}
                     {:= [5 1] :> D :+ L}
                     {:= [0 2] :> U :+ L}
                     {:= [3 3] :> L :+ L}])))

    (it "moves all carts along the track"
      (let [world    (sut/parse-initial sample-tracks-3a)
            expected (sut/parse-initial sample-tracks-3c)
            tick1    (sut/tick world)]
        (->> (:tracks tick1) (should= (:tracks expected)))
        (->> (:carts tick1) (should= (:carts expected)))))

    (it "removes colliding carts from the world"
      (let [world  (sut/parse-initial sample-tracks-3a)
            result (sut/tick (sut/tick world))]
        (should (empty? (:carts result)))))

    (context "turning corners"
      (it "R -> U"
        (let [cart {:= [1 1] :> R} tracks {[2 1] "\\"}]
          (should= U (:> (sut/move cart tracks)))))

      (it "L -> D"
        (let [cart {:= [1 1] :> L} tracks {[0 1] "\\"}]
          (should= D (:> (sut/move cart tracks)))))

      (it "D -> R"
        (let [cart {:= [1 1] :> D} tracks {[1 2] "\\"}]
          (should= R (:> (sut/move cart tracks)))))

      (it "U -> L"
        (let [cart {:= [1 1] :> U} tracks {[1 0] "\\"}]
          (should= L (:> (sut/move cart tracks)))))

      (it "U -> R"
        (let [cart   {:= [1 1] :> U}
              tracks {[1 0] "/"}]
          (should= R (:> (sut/move cart tracks)))))

      (it "D -> L"
        (let [cart   {:= [1 1] :> D}
              tracks {[1 2] "/"}]
          (should= L (:> (sut/move cart tracks)))))

      (it "L -> U"
        (let [cart   {:= [1 1] :> L}
              tracks {[0 1] "/"}]
          (should= U (:> (sut/move cart tracks)))))

      (it "R -> D"
        (let [cart   {:= [1 1] :> R}
              tracks {[2 1] "/"}]
          (should= D (:> (sut/move cart tracks)))))
      )

    (context "intersections")

    #_(it "registers-collisions"
        (let [world     (sut/parse-initial sample-tracks-4)
              collision (as-> (iterate sut/tick world) $
                              (take 2 $)
                              (last $)
                              (do (println $) $)
                              (:collisions $))]
          (should= [7 3] collision)))
    )

  (context "Part 2"
    )
  )