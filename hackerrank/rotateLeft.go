package main

import (
    "fmt"

)

/*
 * Complete the 'rotLeft' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER d
 */

func rotLeft(a []int32, d int32) []int32 {
    
    // Write your code here
    newArr := make([]int32, len(a))
    for i, k := range a{
        newPos := ((int(int32(i) - d) % len(a)) + len(a)) % len(a)
        newArr[newPos] = k
    }
    
    return newArr
}
