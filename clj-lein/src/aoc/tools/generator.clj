(ns aoc.tools.generator
  (:require [clojure.string :as str]
            [clojure.java.io :as io]))

(defn format-namespace [s]
  (let [n (Integer/parseInt s)]
    (if (< n 10) (str "0" s) s)))

(def prod-template (slurp "src/aoc/tools/generator/prod-template.txt"))
(def spec-template (slurp "src/aoc/tools/generator/spec-template.txt"))

(defn replacer [content [key value]]
  (str/replace content (format "!%s!" key) value))

(defn template [raw data]
  (reduce replacer raw (seq data)))

(defn write-file [path content]
  (println "Creating file:" path)
  (io/make-parents path)
  (spit path content))

(defn -main [year day]
  (let [namespace    (format-namespace day)
        data         {"Y" year "NS" namespace "D" day}
        spec-path    (format "./spec/aoc/y%s/d%s_spec.clj" year namespace)
        prod-path    (format "./src/aoc/y%s/d%s.clj" year namespace)
        data-path    (format "./data/%s/d%s.txt" year namespace)
        spec-content (template spec-template data)
        prod-content (template prod-template data)]
    (println (format "Initializing %s Day %s..." year day))
    (write-file spec-path spec-content)
    (write-file prod-path prod-content)
    (write-file data-path "")))