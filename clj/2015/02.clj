(def lines (clojure.string/split-lines (slurp "02.txt")))

(defn dimensions [line] ; "2x3x4"
  (->> (re-matcher       #"(\d+)x(\d+)x(\d+)" line)
       (re-find) ; ("2x3x4" "2"   "3"   "4")
       (drop 1)  ; (        "2"   "3"   "4")
       (map #(Integer/parseInt %)))) ; (2 3 4)

(defn paper-needed [line]
  (let [[l w h] (dimensions line)
        lw (* l w)
        wh (* w h)
        hl (* h l)
        slack (min lw wh hl)
        area (* 2 (+ lw wh hl))]
    (+ area slack)))

(defn ribbon-needed [line]
  (let [[l w h] (dimensions line)
        volume (* l w h)
        lw (* 2 (+ l w))
        wh (* 2 (+ w h))
        hl (* 2 (+ h l))
        min-perimeter (min lw wh hl)]
    (+ volume min-perimeter)))

(require '[clojure.test :refer :all])

(deftest day1
  (testing "components"
    (is (= 58 (paper-needed "2x3x4")))
    (is (= 43 (paper-needed "1x1x10")))
    (is (= 34 (ribbon-needed "2x3x4")))
    (is (= 14 (ribbon-needed "1x1x10"))))

  (testing "part1"
    (is (= 1606483 (reduce + (map paper-needed lines)))))

  (testing "part2"
    (is (= 3842356 (reduce + (map ribbon-needed lines))))))

(run-tests)