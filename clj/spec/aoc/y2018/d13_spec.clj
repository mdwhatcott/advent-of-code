(ns aoc.y2018.d13-spec
  (:require [speclj.core :refer :all]
            [aoc.y2018.d13 :as sut]
            [aoc.y2018.d13 :refer [L R U D S]]
            [clojure.string :refer [trim]]))

(def sample-tracks-1 (slurp "data/2018/d13/sample-tracks-1.txt"))
(def sample-tracks-3a (slurp "data/2018/d13/sample-tracks-3-a.txt"))
(def sample-tracks-3c (slurp "data/2018/d13/sample-tracks-3-c.txt"))
(def sample-tracks-4 (slurp "data/2018/d13/sample-tracks-4.txt"))
(def sample-tracks-5 (slurp "data/2018/d13/sample-tracks-5.txt"))
(def real-tracks (slurp "data/2018/d13.txt"))

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

    (context "turning corners and going through intersections"
      (with tracks {[1 1] "-" [2 2] "|" [3 3] "/" [4 4] "\\" [5 5] "+"})

      (it "R + \\ = D" (-> {:= [3 4] :> R} (sut/move @tracks) :> (should= D)))
      (it "L + \\ = U" (-> {:= [5 4] :> L} (sut/move @tracks) :> (should= U)))
      (it "U + \\ = L" (-> {:= [4 5] :> U} (sut/move @tracks) :> (should= L)))
      (it "D + \\ = R" (-> {:= [4 3] :> D} (sut/move @tracks) :> (should= R)))

      (it "R + / = U" (-> {:= [2 3] :> R} (sut/move @tracks) :> (should= U)))
      (it "L + / = D" (-> {:= [4 3] :> L} (sut/move @tracks) :> (should= D)))
      (it "U + / = R" (-> {:= [3 4] :> U} (sut/move @tracks) :> (should= R)))
      (it "D + / = L" (-> {:= [3 2] :> D} (sut/move @tracks) :> (should= L)))

      (it "L@+ U = L" (-> {:= [5 6] :> U :+ L} (sut/move @tracks) (should= {:= [5 5] :> L :+ S})))
      (it "L@+ L = D" (-> {:= [6 5] :> L :+ L} (sut/move @tracks) (should= {:= [5 5] :> D :+ S})))
      (it "L@+ D = R" (-> {:= [5 4] :> D :+ L} (sut/move @tracks) (should= {:= [5 5] :> R :+ S})))
      (it "L@+ R = U" (-> {:= [4 5] :> R :+ L} (sut/move @tracks) (should= {:= [5 5] :> U :+ S})))

      (it "S@+ R = R" (-> {:= [4 5] :> R :+ S} (sut/move @tracks) (should= {:= [5 5] :> R :+ R})))
      (it "S@+ L = L" (-> {:= [6 5] :> L :+ S} (sut/move @tracks) (should= {:= [5 5] :> L :+ R})))
      (it "S@+ U = U" (-> {:= [5 6] :> U :+ S} (sut/move @tracks) (should= {:= [5 5] :> U :+ R})))
      (it "S@+ D = D" (-> {:= [5 4] :> D :+ S} (sut/move @tracks) (should= {:= [5 5] :> D :+ R})))

      (it "R@+ U = R" (-> {:= [5 6] :> U :+ R} (sut/move @tracks) (should= {:= [5 5] :> R :+ L})))
      (it "R@+ R = D" (-> {:= [4 5] :> R :+ R} (sut/move @tracks) (should= {:= [5 5] :> D :+ L})))
      (it "R@+ D = L" (-> {:= [5 4] :> D :+ R} (sut/move @tracks) (should= {:= [5 5] :> L :+ L})))
      (it "R@+ L = U" (-> {:= [6 5] :> L :+ R} (sut/move @tracks) (should= {:= [5 5] :> U :+ L})))
      )

    (it "solves with sample data"
      (let [world     (sut/parse-initial sample-tracks-4)
            collision (sut/until-first-collision world)]
        (should= [7 3] collision)))

    (it "solves with real data"
      (let [world     (sut/parse-initial real-tracks)
            collision (sut/until-first-collision world)]
        (should= [14 42] collision)))
    )

  (context "Part 2"
    (it "solves with sample data"
      (let [world     (sut/parse-initial sample-tracks-5)
            collision (sut/until-last-cart-remains world)]
        (should= [6 4] collision)))

    #_(it "solves with real data"
      (let [world    (sut/parse-initial real-tracks)
            survivor (sut/until-last-cart-remains world)]
        (should= [14 42] survivor)))

    )
  )