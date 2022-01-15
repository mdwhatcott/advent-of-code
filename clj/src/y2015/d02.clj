(ns y2015.d02)

(def lines (clojure.string/split-lines (slurp "src/y2015/d02.txt")))

(defn dimensions [line]                                     ; "2x3x4"
  (->> (re-matcher #"(\d+)x(\d+)x(\d+)" line)
       (re-find)                                            ; ("2x3x4" "2"   "3"   "4")
       (rest)                                               ; (        "2"   "3"   "4")
       (map #(Integer/parseInt %))))                        ; (2 3 4)

(defn paper-needed [[length width height]]
  (let [lw    (* length width)
        wh    (* width height)
        hl    (* height length)
        slack (min lw wh hl)
        area  (* 2 (+ lw wh hl))]
    (+ area slack)))

(defn ribbon-needed [[length width height]]
  (let [volume        (* length width height)
        lw            (* 2 (+ length width))
        wh            (* 2 (+ width height))
        hl            (* 2 (+ height length))
        min-perimeter (min lw wh hl)]
    (+ volume min-perimeter)))

(require '[clojure.test :refer :all])

(deftest day1
  (testing "components"
    (is (= 58 (paper-needed '(2 3 4))))
    (is (= 43 (paper-needed '(1 1 10))))
    (is (= 34 (ribbon-needed '(2 3 4))))
    (is (= 14 (ribbon-needed '(1 1 10)))))

  (let [all-dimensions (map dimensions lines)]

    (testing "part 1"
      (is (= 1606483
             (reduce + (map paper-needed all-dimensions)))))

    (testing "part 2"
      (is (= 3842356
             (reduce + (map ribbon-needed all-dimensions)))))))

(run-tests)