(ns aoc.y2019.intcode-spec
  (:require [speclj.core :refer :all]
            [aoc.y2019.intcode :refer :all]))

(describe "The IntCode Interpreter"

  (context "Single Ticks"

    (it "adds (day 2, part 1)"
      (->> (tick,,, {:running true :pointer 0, :memory [1 9 10 3, 2 3 11 0 99 30 40 50]})
           (should= {:running true :pointer 4, :memory [1 9 10 70 2 3 11 0 99 30 40 50]})))

    (it "multiplies (day 2, part 1)"
      (->> (tick,,, {:running true :pointer 4, :memory [1,,, 9 10 70 2 3 11 0 99 30 40 50]})
           (should= {:running true :pointer 8, :memory [3500 9 10 70 2 3 11 0 99 30 40 50]})))

    (it "halts (day 2, part 1)"
      (->> (tick,,, {:running true, :pointer 8, :memory [3500 9 10 70 2 3 11 0 99 30 40 50]})
           (should= {:running false :pointer 9, :memory [3500 9 10 70 2 3 11 0 99 30 40 50]})))

    (it "receives input (day 5, part 1)"
      (->> (tick,,, {:running true :pointer 0 :memory [3, 0 4 0 99] :inputs [5]})
           (should= {:running true :pointer 2 :memory [5, 0 4 0 99] :inputs []})))

    (it "provides output (day 5, part 1"
      (->> (tick,,, {:running true :pointer 2 :memory [5 0 4 0 99]})
           (should= {:running true :pointer 4 :memory [5 0 4 0 99] :outputs [5]})))

    (it "waits for missing input (nop, day 5, part 1)"
      (->> (tick,,, {:running true :pointer 0 :memory [3, 0 4 0 99] :inputs []})
           (should= {:running true :pointer 0 :memory [3, 0 4 0 99] :inputs []})))
    )

  (context "Examples, Day 2, Part 1"

    (it "[2,0,0,0,99]"
      (->> (run-simple,,,, [1, 0 0 0 99])
           (should= [2, 0 0 0 99])))

    (it "[2,3,0,3,99]"
      (->> (run-simple,,,, [2 3 0 3, 99])
           (should= [2 3 0 6, 99])))

    (it "[2,4,4,5,99,0]"
      (->> (run-simple,,,, [2 4 4 5 99 0,,,])
           (should= [2 4 4 5 99 9801])))

    (it "[1,1,1,4,99,5,6,0,99]"
      (->> (run-simple,,,, [1, 1 1 4 99 5 6 0 99])
           (should= [30 1 1 4 2, 5 6 0 99]))))

  (context "Examples, Day 5, Part 1"

    (it "[3,0,4,0,99] (i/o: 42)"
      (->> (run,,,, {:pointer 0 :inputs [42] :memory [3, 0 4 0 99] :outputs []})
           (should= {:pointer 5 :inputs [],, :memory [42 0 4 0 99] :outputs [42] :running false})))
    )

  ;; TODO: parameter modes

  )