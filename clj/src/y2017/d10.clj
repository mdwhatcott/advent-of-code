(ns y2017.d10)

(defn make-ring []
  (vec (take 256 (range))))

(defn twist [ring at length]
  ;;   4--5   pinch   4  5           4   1
  ;;  /    \  5,0,1  / \/ \  twist  / \ / \
  ;; 3      0  -->  3      0  -->  3   X   0
  ;;  \    /         \ /\ /         \ / \ /
  ;;   2--1           2  1           2   5
  (let [rotation  (take (count ring) (drop at (cycle ring)))
        slice     (take length rotation)
        reversed  (reverse slice)
        remaining (drop (count slice) rotation)
        complete  (concat reversed remaining)
        final     (take
                    (count ring)
                    (drop (- (count ring) at) (cycle complete)))]
    final))


(defn bi-product [ring]
  (* (first ring)
     (second ring)))

(defn braid [original-ring lengths rounds]
  (let [lengths-count (count lengths)
        total-twists  (* rounds lengths-count)]
    (loop [ring original-ring
           l    0
           at   0
           skip 0]
      (if (>= l total-twists)
        ring
        (let [length    (nth lengths (mod l lengths-count))
              next-ring (twist ring at length)
              next-l    (inc l)
              next-at   (mod (+ at length skip) (count ring))
              next-skip (inc skip)]
          (recur next-ring next-l next-at next-skip))))))

(defn xor16s [sparse-hash]
  (let [groups (partition 16 sparse-hash)
        xOrs   (map #(reduce bit-xor %) groups)]
    xOrs))

(defn knot-hash [content]
  (let [salt        '(17 31 73 47 23)
        input       (concat content salt)
        sparse-hash (braid (make-ring) input 64)
        dense-hash  (xor16s sparse-hash)
        checksum    (reduce str (map #(format "%02x" %) dense-hash))]
    checksum))

(require '[clojure.test :refer :all])

(deftest components
  (is (= [2 1 0 3 4] (twist [0 1 2 3 4] 0 3)))
  (is (= [4 3 0 1 2] (twist [2 1 0 3 4] 3 4)))
  (is (= [4 3 0 1 2] (twist [4 3 0 1 2] 3 1)))
  (is (= [3 4 2 1 0] (twist [4 3 0 1 2] 1 5)))
  (is (= [3 4 2 1 0] (braid [0 1 2 3 4] [3 4 1 5] 1))))

(deftest part-1
  (let [part1-lengths [31 2 85 1 80 109 35 63 98 255 0 13 105 254 128 33]
        hashed-ring   (braid (make-ring) part1-lengths 1)]
    (is (= 6952 (bi-product hashed-ring)))))

(deftest hashes
  (is (= "a2582a3a0e66e6e86e3812dcb672a272" (knot-hash (.getBytes ""))))
  (is (= "33efeb34ea91902bb2f59c9920caa6cd" (knot-hash (.getBytes "AoC 2017"))))
  (is (= "3efbe78a8d82f29979031a4aa0b16a9d" (knot-hash (.getBytes "1,2,3"))))
  (is (= "63960835bcdc130f0b66d7ff4f6a5a8e" (knot-hash (.getBytes "1,2,4")))))

(deftest part-2
  (is (= "28e7c4360520718a5dc811d3942cf1fd"
         (knot-hash (.getBytes (slurp "src/y2017/d10.txt"))))))

(run-tests)