package main

import (
    lmd "github.com/mayukh42/project-lambda/lambda"
)

func main() {
    /** Map tests
     * Multiply elements of an array by 2
     * Convert words in a list to uppercase
     */
    lmd.TestMap()

    /** Reduce Tests
     * Find sum of list of integers
     * Find product of list of integers
     */
    lmd.TestReduce()

    /** Filter Tests
     * Find odd numbers in a list of integers
     * Find vowels in a list of strings
     */
    lmd.TestFilter()
}
