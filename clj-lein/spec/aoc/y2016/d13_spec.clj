(ns aoc.y2016.d13-spec
  (:require [speclj.core :refer :all]))

; Search Algorithms Demo: https://www.redblobgames.com/pathfinding/a-star/introduction.html
; Hamming Weight Explanation: https://en.wikipedia.org/wiki/Hamming_weight
; Hamming Weight Demo: https://go.dev/play/p/Uf9xmV2pMcs

(defn bits [n]
  (->> n
       (iterate #(bit-and % (dec %)))
       (take-while pos?)
       count))

(defn hallway? [seed [x y]]
  (and (not (neg? x))
       (not (neg? y))
       (->> seed (+ (* x x) (* 3 x) (* 2 x y) y (* y y)) bits even?)))

(defn cardinal-neighbors [[x y]]
  [[(dec x) y]
   [(inc x) y]
   [x (dec y)]
   [x (inc y)]])

(defn maze-neighbors [seed point]
  (filter (partial hallway? seed) (cardinal-neighbors point)))

(defn bfs-step [graph {:keys [explored frontier]}]
  (let [current   (peek frontier)
        frontier  (pop frontier)
        neighbors (remove #(contains? explored %) (graph current))]
    {:frontier (reduce conj (vec neighbors) frontier)
     :explored (reduce #(assoc %1 %2 current) explored neighbors)}))

(defn trace [target explored]
  (loop [path [target]]
    (let [step (explored (peek path))]
      (if (nil? step)
        (reverse path)
        (recur (conj path step))))))

(defn emit [seed maxX maxY explored]
  (newline)
  (doseq [y (range maxY)
          x (range maxX)]
    (when (zero? x) (newline))
    (if-not (hallway? seed [x y])
      (print "#")
      (if (contains? explored [x y])
        (print "•")
        (print " "))))
  (newline))

(defn bfs [origin target graph emit]
  (loop [state {:frontier [origin] :explored {origin nil}}]
    (let [explored (:explored state)]
      (if (contains? explored target)
        (let [path (trace target explored)]
          (emit (set path))
          path)
        (do
          #_(emit explored)
          (recur (bfs-step graph state)))))))

(describe "2016 Day 13"
  (it "counts '1' bits"
    (should= 0 (bits 0)) #_"0000"
    (should= 1 (bits 1)) #_"0001"
    (should= 1 (bits 2)) #_"0010"
    (should= 1 (bits 4)) #_"0100"
    (should= 2 (bits 3)) #_"0011"
    (should= 8 (bits 255)))

  (it "walls vs halls"
    (should= false (hallway? 10 [0 -1]))
    (should= false (hallway? 10 [-1 0]))
    (should= true (hallway? 10 [0 0]))
    (should= true (hallway? 10 [1 1]))
    (should= false (hallway? 10 [9 6]))
    (should= false (hallway? 10 [1 0])))

  (it "cardinal neighbors"
    (should= [[0 1] [2 1] [1 0] [1 2]] (cardinal-neighbors [1 1])))

  (it "breadth-first search"
    (let [expected {:frontier [:b] :explored {:a nil :b :a}}
          input    {:frontier [:a] :explored {:a nil}}]
      (should= expected (bfs-step {:a [:b]} input)))

    (let [expected {:explored {:a nil :b :a} :frontier []}
          input    {:explored {:a nil :b :a} :frontier [:b]}]
      (should= expected (bfs-step {:a [:b] :b [:a]} input)))

    (should= [:a :b :c] (trace :c {:c :b :b :a :a nil}))
    (should= [:a :b :c] (bfs :a :c {:a [:b] :b [:c]} (fn [_])))
    (should= 9 (dec (count (bfs [1 1] [7 4] cardinal-neighbors (fn [_]))))))

  (it "maze neighbors"
    (should= [] (maze-neighbors 10 [5 3]))
    (should= (cardinal-neighbors [3 2]) (maze-neighbors 10 [3 2]))
    (should= [[0 1] [1 2]] (maze-neighbors 10 [1 1])))

  (it "part 1"
    (should= 11 (dec (count (bfs [1 1] [7 4] (partial maze-neighbors 10) (partial emit 10 10 7)))))
    (should= 96 (dec (count (bfs [1 1] [31 39] (partial maze-neighbors 1358) (partial emit 1358 50 50)))))))

; Output of emit:
;  ##  # # ## #    ##   #  # #### ####   #  #  #  #
;  •##  ## ## #  #   ### #     #     ## ######    #
; #•••# ##  ######## # #####   #  ## ## #     ## ###
;  ##••••••••••• ###     ##### #####    #  ##  # ###
; ##### ## ####•  # ##    #  # #  # ## #### ##    ##
;  ## #  # ## #•# ######  # ##  # ####  # #######  #
;    ###      #••#  #  #### ###  #  # # ##  #
; ## ####  ### #•## # #   #  # # ## ##      #  ####
;     ###### ###• ###  ## #  ###  ## #  ## ####   #
; # #  ##  # ••••  ### #  ###      # #####  # #   #
; #        ##•## #     #    # ## # ###  # # ####  ##
;  ###### # #•##   ### ##   ## #  #  ## ##   #######
; #    ##  ##•# ###  # ### # ## #  #  ## #     #  ##
; #  # ### #••#   # ##   #  # #  #       ##  # ##  #
; ####  #  #•###### ##### #  ###  ## ### ####   #
; #     ## #•##  #   ## # ## ####  # ###    #   #  #
; # ##   #  •••• #      #     # ###   ######## #####
; # ### #######•######## ## # ### ##   ##  # # ##  #
; #  ## #   ###•##•••# ## #  #     ##    # ##      #
;  # ## #     #••••#••# ## # #### # # ##  # ##### ##
; ##   ### ## # ### #•## ###   ##  ## # #      ## #
;    #  ## #  ##  ## • #     #  ## #  #  ##  # ## #
; #####    #   ### # • ## ##### #  #  ## #####  ###
;    # ## ### #  # ##•# ###  #  ## #####          #
; #  ####  ##  # # ##••#•••••#   # #  ##### ## ## ##
; #    # #   #  ##  ##•••###•## ##  #  ## # ## #   #
; ###  ## ##  # ### # #### #•## ### #     #    # # #
; # ##  #######  #  #     ##•••• #  ###### ##  #  #
;    ###  #      #  ### # # ###••#     # ## ### #  #
; # #  ## #  ######## ### ### ##•##  # #  ### #  # #
; #  #   ####    #       #   # #•##### # #   ###
;  #   #    #  # # ##  # ###  ##•##   ## ### #### ##
; # ############ # #####    # # ••••# ##      # # #
; # #  ##  #    ##  #    ## # #  ##•#    ## # ##  #
; # #    # ## # # # # ### #  ######•# ### # ## #  #
;   ## #  # # # ##  ##  ## #     #••### ##   # ### #
;  # ## #   # ## #   # # ######  #•#  ## # # # # ###
;   # #  ###   # ## ## ##  ## #  #•••  # #  ##     #
; #  ### # # # #### ##  #     ##  ##•• #  # ####
; ## ###  ##  #   #     #  ######••##•####   ######
;     # # # #  ## # ## #####  # ##•##•## #    #   #
; ### ##  #  # #  ## # ##  #  ####•##•  ##### #   #
;    # # ###   #   ##      ###  ##••••# ##  ####  #
;  # # # #### ### # ##### #   #    ## #       #### #
; ## #    ###  ##      ##  ##   ### # # ##  # ## ###
;   ### #  # #   ##  # ### # ## # ##  ## ###   #   #
; # # # #  ## ## #####  #  ## # ## #   # # #   ##  #
; # ##  ##  #### ##     ##  ##   #### ## # ## # # #
; ## #   ###  #     ##   ### #     ## ## # ##  ##  #
;  #### #   # #  ## ### #    ##  #      ##  ## # #