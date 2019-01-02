package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"syscall/js"
	"time"
)

var (
	out      chan int
	document js.Value
)

func init() {
	out = make(chan int)
	document = js.Global().Get("document")
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// div := document.Get("body")

	data := generateRandomNumber(25)

	rawNums := document.Call("getElementById", "raw nums")
	rawNums.Call("append", arrayToString(data, " "))
	go sleepSort(data)

	sortedNums := document.Call("getElementById", "sorted nums")
	for n := range out {
		sortedNums.Call("append", " "+strconv.Itoa(n))
	}
}

func sleepSort(nums []int) {
	blockers := make(chan struct{}, len(nums))
	for _, n := range nums {
		blockers <- struct{}{}
		go func(n int, blocker chan struct{}) {
			defer func() { <-blocker }()
			time.Sleep(time.Duration(n) * 25 * time.Millisecond)
			// fmt.Println(n)
			out <- n
		}(n, blockers)
	}

	for i := 0; i < cap(blockers); i++ {
		blockers <- struct{}{}
	}
	close(out)
}

func generateRandomNumber(size int) []int {
	randNumber := make([]int, size, size)
	for i := 0; i < size; i++ {
		randNumber[i] = rand.Intn(100)
	}

	return randNumber
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
