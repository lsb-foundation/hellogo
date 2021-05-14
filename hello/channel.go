package hello

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func randomNumbers(length uint) *[]int {
	retArr := make([]int, length)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for index := uint(0); index < length; index++ {
		retArr[index] = int(random.Int31())
	}
	return &retArr
}

func maxOfArray(arr *[]int) int {
	maxValue := math.MinInt64
	for _, value := range *arr {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func TestChannel() {
	intArr := randomNumbers(1_0000_0000)
	startTime := time.Now()
	ordinaryMaxValue := maxOfArray(intArr)
	fmt.Println("Ordinary method time cost: ", time.Since(startTime).Milliseconds(), "ms")
	fmt.Println("Max value by ordinary method: ", ordinaryMaxValue)

	routineChannel := make(chan int, 2)
	startTime = time.Now()
	go func() {
		slice := (*intArr)[:len(*intArr)/2]
		routineChannel <- maxOfArray(&slice)
	}()
	go func() {
		slice := (*intArr)[len(*intArr)/2:]
		routineChannel <- maxOfArray(&slice)
	}()
	routineArr := []int{0, 0}
	routineArr[0] = <-routineChannel
	routineArr[1] = <-routineChannel
	routineMaxValue := maxOfArray(&routineArr)
	fmt.Println("Goroutines time cost: ", time.Since(startTime).Milliseconds(), "ms")
	fmt.Println("Max value by goroutines: ", routineMaxValue)
}
