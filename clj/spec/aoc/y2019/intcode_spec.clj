(ns aoc.y2019.intcode-spec
  (:require [speclj.core :refer :all]
            [aoc.y2019.intcode :refer :all]))

(describe "The IntCode Interpreter"

  (context "Single Ticks"

    (it "adds"
      (->> (tick,,, {:pointer 0, :memory [1 9 10 3, 2 3 11 0 99 30 40 50]})
           (should= {:pointer 4, :memory [1 9 10 70 2 3 11 0 99 30 40 50]})))

    (it "multiplies"
      (->> (tick,,, {:pointer 4, :memory [1,,, 9 10 70 2 3 11 0 99 30 40 50]})
           (should= {:pointer 8, :memory [3500 9 10 70 2 3 11 0 99 30 40 50]})))

    (it "halts"
      (->> (tick,,, {:halted false :pointer 8, :memory [3500 9 10 70 2 3 11 0 99 30 40 50]})
           (should= {:halted true, :pointer 9, :memory [3500 9 10 70 2 3 11 0 99 30 40 50]})))

    )

  (context "Examples, Day 2, Part 1"

    (it "[2,0,0,0,99]"
      (should= [2 0 0 0 99] (run [1 0 0 0 99])))

    (it "[2,3,0,3,99]"
      (should= [2 3 0 6 99] (run [2 3 0 3 99])))

    (it "[2,4,4,5,99,0]"
      (should= [2 4 4 5 99 9801] (run [2 4 4 5 99 0])))

    (it "[1,1,1,4,99,5,6,0,99]"
      (should= [30 1 1 4 2 5 6 0 99] (run [1 1 1 4 99 5 6 0 99]))))

  )