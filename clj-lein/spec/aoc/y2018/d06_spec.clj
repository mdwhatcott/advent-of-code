(ns aoc.y2018.d06-spec
  (:require [speclj.core :refer :all]
            [aoc.y2018.d06 :as sut]
            [aoc.data :as data]))

(def sample-data
  (str "1, 1" "\n"
       "1, 6" "\n"
       "8, 3" "\n"
       "3, 4" "\n"
       "5, 5" "\n"
       "8, 9" "\n"))

(def real-data
  (data/read-str 2018 6))

(describe "2018 Day 6"
  (context "Part 1"
    (it "parses coordinates from input"
      (should= #{[1 1] [1 6] [8 3] [3 4] [5 5] [8 9]}
               (sut/parse-coordinates sample-data)))

    (it "limits the skope of the arena"
      (should= [[1 1] [8 9]]
               (sut/scope (sut/parse-coordinates sample-data))))

    (it "computes manhattan distance"
      (should= 5 (sut/manhattan [0 0] [5 0]))
      (should= 5 (sut/manhattan [0 0] [0 5]))
      (should= 10 (sut/manhattan [0 0] [5 5])))

    (it "identifies landmarks with infinite area"
      (let [landmarks (sut/parse-coordinates sample-data)]
        (should (sut/infinite? landmarks [1 1]))
        (should-not (sut/infinite? landmarks [3 4]))))

    (it "calculates areas of finite landmarks"
      (let [landmarks (sut/parse-coordinates sample-data)
            arena     (sut/arena (sut/scope landmarks))
            areas     (sut/calculate-areas arena landmarks)]
        (should= {[3 4] 9
                  [5 5] 17} areas)))

    (it "solves part 1 with sample data"
      (should= 17 (sut/part1 sample-data)))

    #_(it "solves with real data (albeit slowly @ 12+s)"
        (should= 3604 (sut/part1 real-data)))

    )

  (context "Part 2"
    (it "solves with sample data"
      (should= 16 (sut/part2 32 sample-data)))

    #_(it "solves with real data (albeit somewhat slowly @ 2.5s)"
        (should= 46563 (sut/part2 10000 real-data)))
    )
  )
