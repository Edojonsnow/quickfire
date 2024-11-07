package random

import (
	"fmt"
	"time"

	"math/rand"
)


func init(){

	rand.Seed(time.Now().UnixNano())

	pool  := []int{1,2,3,4,5,6,7,8,9,10}

	result := generateRandomNumbers(pool, 5)

    fmt.Println("Random numbers:", result)
}


func generateRandomNumbers(pool []int, count int) []int {
    if count > len(pool) {
        count = len(pool)
    }

    result := make([]int, count)
    for i := 0; i < count; i++ {
        randomIndex := rand.Intn(len(pool))
        result[i] = pool[randomIndex]
        pool = append(pool[:randomIndex], pool[randomIndex+1:]...)
    }

    return result
}