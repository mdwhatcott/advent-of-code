(ns aoc.perf)

(defmacro time-ns [ex]
  `(let [start# (. System (nanoTime))]
     (do ~ex (- (. System (nanoTime)) start#))))

(defn time-it [n f]
  (time-ns (dotimes [_ n] (f))))

(def template "- %s - %d x %d ops, average per-op: %dns")

(defn benchmark
  ([title f] (benchmark 10 10000 title f))
  ([n title f] (benchmark 10 n title f))
  ([x n title f]
   (if (or (<= x 0) (<= n 0))
     (throw (IllegalArgumentException. "x and n must be positive"))
     (let [times   (for [_ (range x)] (time-it n f))
           per-ops (map #(quot % n) times)
           average (quot (apply + per-ops) (count per-ops))]
       (println (format template title x n average))
       average))))
