Chapter 2 Examples and Exercises
================================

## Examples
- ftoc
- boiling
- cf
- echo4
- fib
- tempconv
- tempconv0
- popcount

## Exercises
- Exercise 2.1:
  - Add types, constants and functions to `tempconv` for processing
    temperatures in the Kelvin scale.
  - Packages:
    - cfk
- Exercise 2.2
  - Write a general-purpose unit-conversion program analogous to `cf` that
    reads numbers from its command-line arguments or from the standard
    input if there are no arguments, and converts each number into units like
    temperature in Celsius and Fahrenheit, length in feet and meters, weight in
    pounds and kilograms, and the like.
  - Packages:
    - unitconv
    - conv
- Exercise 2.3
  - Rewrite `PopCount` to use a loop instead of a single expression. Compare
    the performance of the two variants.
  - Packages:
    - popcountLoop
    - popcountTest
- Exercise 2.4
  - Write a version of `PopCount` that counts bits by shifting its argument
    through 64 bit postitions, testing the rightmost bit each time. Compare its
    performance to the table-lookup version.
  - Packages:
    - popcountShift64
    - popcountTest
    - popcountInitTest: Just a pretty print version of popcount that gave me a
      better understanding of the algorithm.
