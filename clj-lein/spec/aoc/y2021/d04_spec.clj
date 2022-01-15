(ns aoc.y2021.d04-spec
  (:require [speclj.core :refer :all]
            [aoc.y2021.d04 :as sut]
            [aoc.data :as data]))

(def real-data
  (data/read-lines 2021 4))

(def sample-data
  ["7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1"
   ""
   "22 13 17 11  0"
   " 8  2 23  4 24"
   "21  9 14 16  7"
   " 6 10  3 18  5"
   " 1 12 20 15 19"
   ""
   " 3 15  0  2 22"
   " 9 18 13 17  5"
   "19  8  7 25 23"
   "20 11 10 24  4"
   "14 21 16 12  6"
   ""
   "14 21 17 24  4"
   "10 16 15  9 19"
   "18  8 23 26 20"
   "22 11 13  6  5"
   " 2  0 12  3  7"])

(def call-order
  [7 4 9 5 11 17 23 2 0 14 21 24 10 16 13 6 15 25 12 22 18 20 8 19 3 26 1])

(def raw-boards
  [["22 13 17 11  0"
    " 8  2 23  4 24"
    "21  9 14 16  7"
    " 6 10  3 18  5"
    " 1 12 20 15 19"]
   [" 3 15  0  2 22"
    " 9 18 13 17  5"
    "19  8  7 25 23"
    "20 11 10 24  4"
    "14 21 16 12  6"]
   ["14 21 17 24  4"
    "10 16 15  9 19"
    "18  8 23 26 20"
    "22 11 13  6  5"
    " 2  0 12  3  7"]])

(describe "2021 Day 4"
  (context "Part 1"
    (it "parses the input"
      (should= {:call-order call-order
                :boards     raw-boards}
               (sut/parse-input sample-data)))

    (context "Bingo Board"
      (it "converts raw input to board structure"
        (should= {:score      0
                  :checksum   0
                  :call-order call-order
                  :called     []
                  :marked     #{}
                  :unmarked   #{22 13 17 11 0
                                8, 2, 23 4, 24
                                21 9, 14 16 7
                                6, 10 3, 18 5
                                1, 12 20 15 19}
                  :wins       [#{22 13 17 11 0,}
                               #{8, 2, 23 4, 24}
                               #{21 9, 14 16 7,}
                               #{6, 10 3, 18 5,}
                               #{1, 12 20 15 19}
                               #{22 8, 21 6, 1,}
                               #{13 2, 9, 10 12}
                               #{17 23 14 3, 20}
                               #{11 4, 16 18 15}
                               #{0, 24 7, 5, 19}]}
                 (sut/prepare-board call-order (first raw-boards))))

      (it "tracks a called number"
        (let [board (sut/prepare-board call-order (first raw-boards))
              board (sut/call-next-number board)]
          (should-contain 7 (:marked board))
          (should-not-contain 7 (:unmarked board))
          (should-not-contain 7 (:call-order board))
          (should= [7] (:called board))))

      (it "plays a game"
        (let [board (sut/prepare-board call-order (last raw-boards))
              board (sut/play board)]
          (should= 24 (peek (:called board)))
          (should= 188 (:score board))
          (should= 4512 (:checksum board)))))

    (it "solves with sample data"
      (should= 4512 (sut/part1 sample-data)))

    (it "solves with real data"
      (should= 29440 (sut/part1 real-data)))
    )

  (context "Part 2"
    (it "solves with sample data"
      (should= 1924 (sut/part2 sample-data)))

    (it "solves with real data"
      (should= 13884 (sut/part2 real-data)))
    )
  )
