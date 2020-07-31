package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"time"

	selfSort "algorithm/sortAlgorithm"
)

func main() {
	testnums := 80000
	testArr := createArr(testnums)
	// testArr := []int{5, 1, 3, 4, 5, 7, 8, 9, 10, 32, 22, 18, 17, 25, 22, 11, 12, 22, 12, 33, 44, 44}
	var re []int

	sort := selfSort.NewSortService()
	re = sort.Sort.InsertSort(testArr)
	fmt.Println(re)

	sortValue := reflect.ValueOf(sort.Sort)
	sortType := reflect.TypeOf(sort.Sort)

	a := reflect.ValueOf(testArr)
	// aNums := reflect.ValueOf(testnums)
	in := []reflect.Value{a}
	// inNums := []reflect.Value{a, aNums}
	// ret := sortValue.Method(0).Call(in)
	// fmt.Println(sortType.Method(0).Name == "BubbleSort")
	// a := sortType.Method(0).
	// re := a.Method(testArr)
	// for _, method := range sortType.Method {
	// 	fmt.Println(method)
	// }
	// var re []int
	// re = selfSort.BubbleSort(testArr)
	for i := 0; i < sortValue.NumMethod(); i++ {
		methodName := sortType.Method(i).Name
		start := time.Now()
		switch methodName {
		// case "CountSort":
		// 	sortValue.MethodByName(methodName).Call(inNums)
		default:
			sortValue.MethodByName(methodName).Call(in)
		}
		elapsed := time.Since(start)
		fmt.Printf("method name : %-16s\t elapsed time: %6s\t\t  \n", methodName, elapsed)
	}
}

func createArr(nums int) []int {
	var arr []int = []int{}
	for i := 0; i < 20000; i++ {
		randDateString := fmt.Sprintf("%d", rand.New(rand.NewSource(time.Now().UnixNano())).Intn(nums))
		randDateInt, _ := strconv.Atoi(randDateString)
		arr = append(arr, randDateInt)
	}
	return arr

}
