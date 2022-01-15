(ns aoc.y2021.d10-spec
  (:require [speclj.core :refer :all]
            [aoc.y2021.d10 :as sut]
            [aoc.data :as data]))

(def real-data
  (data/read-lines 2021 10))

(def sample-data
  ["[({(<(())[]>[[{[]{<()<>>"
   "[(()[<>])]({[<{<<[]>>("
   "{([(<{}[<>[]}>{[]{[(<()>"
   "(((({<>}<{<{<>}{[]{[]{}"
   "[[<[([]))<([[{}[[()]]]"
   "[{[{({}]{}}([{[{{{}}([]"
   "{<[[]]>}<{[{[{[]{()[[[]"
   "[<(<(<(<{}))><([]([]()"
   "<{([([[(<>()){}]>(<<{{"
   "<{([{{}}[<[[[<>{}]]]>[]]"])

(describe "2021 Day 10"
  (context "Part 1"
    (it "analyzes"
      (should= [\[] (sut/analyze "["))
      (should= [] (sut/analyze "[]"))
      (should= [\]] (sut/analyze "[]]"))

      (should= [\(] (sut/analyze "("))
      (should= [] (sut/analyze "()"))
      (should= [\)] (sut/analyze "())"))

      (should= [\{] (sut/analyze "{"))
      (should= [] (sut/analyze "{}"))
      (should= [\}] (sut/analyze "{}}"))

      (should= [] (sut/analyze "([{<>}])"))
      (should= [] (sut/analyze "([{<>[]}])"))
      (should= [\( \[ \{] (sut/analyze "([{<>[]"))
      (should= [\>] (sut/analyze "([{>"))

      (should= "}>{[]{[(<()>" (apply str (sut/analyze "{([(<{}[<>[]}>{[]{[(<()>"))))

    (it "solves with sample data"
      (should= 26397 (sut/part1 sample-data)))

    (it "solves with real data"
      (should= 299793 (sut/part1 real-data)))

    )

  (context "Part 2"
    (it "completes"
      (should= "}}]])})]" (apply str (sut/completion "[({(<(())[]>[[{[]{<()<>>")))
      (should= ")}>]})" (apply str (sut/completion "[(()[<>])]({[<{<<[]>>(")))
      (should= "}}>}>))))" (apply str (sut/completion "(((({<>}<{<{<>}{[]{[]{}")))
      (should= "]]}}]}]}>" (apply str (sut/completion "{<[[]]>}<{[{[{[]{()[[[]")))
      (should= "])}>" (apply str (sut/completion "<{([{{}}[<[[[<>{}]]]>[]]"))))

    (it "scores"
      (should= 288957 (sut/score-completion "}}]])})]"))
      (should= 1480781 (sut/score-completion "}}>}>))))")))

    (it "solves with sample data"
      (should= 288957 (sut/part2 sample-data)))

    (it "solves with real data"
      (should= 3654963618 (sut/part2 real-data)))
    )
  )
