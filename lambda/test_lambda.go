package lambda

import (
    "fmt"
    "strings"
)

func TestMap() {
    twice := func(x Any) Any {
        /** function subtyping is contravariant in args and covariant in return args contravariance: x should be lowest type in type hierarchy in the absence of which, we just downcast it to the right type.
          return covariance: y should not be cast back to interface
        */
        xi := x.(int)
        yi := 2 * xi
        return yi
    }
    xs := List{2, 3, 5, 8, 13}
    ys := Map(xs, twice)
    fmt.Printf("ys = %v\n", ys)
    fmt.Printf("42 x 2 = %v\n", twice(42))

    toUpper := func(s Any) Any {
        t := s.(string)
        return strings.ToUpper(t)
    }
    ss := List{"Foo", "BAR", "bAz", "quux"}
    ts := Map(ss, toUpper)
    fmt.Printf("ts = %v\n", ts)
}

func TestReduce() {
    sum := func(x Any, y Any) Any {
        xi := x.(int)
        yi := y.(int)
        return xi + yi
    }
    xs := List{1, 2, 3, 4, 5}
    fmt.Printf("sum of %v = %v\n", xs, Reduce(xs, 0, sum))

    product := func(x Any, y Any) Any {
        xi := x.(int)
        yi := y.(int)
        return xi * yi
    }
    fmt.Printf("product of %v = %v\n", xs, Reduce(xs, 1, product))
}

func TestFilter() {
    isOdd := func(x Any) bool {
        xi := x.(int)
        return xi%2 != 0
    }

    isVowel := func(c Any) bool {
        vowels := map[string]bool{
            "a": true,
            "e": true,
            "i": true,
            "o": true,
            "u": true,
        }
        d := c.(string)
        if _, ok := vowels[strings.ToLower(d)]; !ok {
            return false
        }

        return true
    }

    fib := List{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
    fmt.Printf("odds in fibonacci = %v\n", Filter(fib, isOdd))

    letters := List{"p", "R", "i", "m", "E"}
    fmt.Printf("vowels in 'pRimE' = %v\n", Filter(letters, isVowel))
}
