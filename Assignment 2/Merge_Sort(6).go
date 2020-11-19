package example

import (
	"fmt"
	"time"
)

func Mergesort(numbers []int) []int {
	l := len(numbers)
	if l <= 1 {
		return numbers
	}

	m := l/2

	sortedLeft := Mergesort(numbers[0:m])
	sortedRight := Mergesort(numbers[m:l])

	return Merge(sortedLeft, sortedRight)
}

func Merge(left []int, right []int) []int {
	leftLength := len(left)
	rightLength := len(right)

	if leftLength == 0 {
		return right
	}
	if rightLength == 0 {
		return left
	}

	result := make([]int, (leftLength+rightLength))

	leftindex := 0
	rightindex := 0
	resultindex := 0
	var rightnum, leftnum int

	for leftindex < leftLength || rightindex < rightLength {
		if leftindex < leftLength  && rightindex < rightLength {
			leftnum = left[leftindex]
			rightnum = right[rightindex]

			if leftnum <= rightnum {
				result[resultindex] = leftnum
				leftindex++
			}else{
				result[resultindex] = rightnum
				rightindex++
			}

		} else if leftindex < leftLength {
			leftnum = left[leftindex]
			result[resultindex] = leftnum
			leftindex++
		} else if rightindex < rightLength {
			rightnum = right[rightindex]
			result[resultindex] = rightnum
			rightindex++
		}

		resultindex++
	}

	return result
}

func MergeSortAsync(numbers [] int, resultChan chan []int)  {
	l := len(numbers)
	if l <= 1 {
		resultChan <- numbers
		return
	}

	m := l/2

	leftchan := make(chan []int, 1)
	rightchan := make(chan []int, 1)

	go MergeSortAsync(numbers[0:m], leftchan)
	go MergeSortAsync(numbers[m:l], rightchan)
	go MergeAsync(<- leftchan, <- rightchan, resultChan)
}

func MergeAsync(left []int, right []int, resultChannel chan []int) {
	leftLength := len(left)
	rightLength := len(right)

	if leftLength == 0 {
		resultChannel <- right
		return
	}
	if rightLength == 0 {
		resultChannel <- left
		return
	}

	result := make([]int, (leftLength+rightLength))
	leftindex := 0
	rightindex := 0
	resultindex := 0
	var rightnum, leftnum int

	for leftindex < leftLength || rightindex < rightLength {
		if leftindex < leftLength  && rightindex < rightLength {
			lnum = left[leftindex]
			rnum = right[rightindex]

			if leftnum <= rightnum {
				result[resultindex] = leftnum
				leftindex++
			}else{
				result[resultindex] = rightnum
				rightindex++
			}

		} else if leftindex < leftLength {
			leftnum = left[leftindex]
			result[resultindex] = leftnum
			leftindex++
		} else if rightindex < rightLength {
			rightnum = right[rightindex]
			result[resultindex] = rightnum
			rightindex++
		}

		resultindex++
	}

	resultChannel <- result
}

func Mergesortone(){
	lim := 7
	largeArray := make([]int, lim)

	for i := 0; i < lim; i++{
		largeArray[i] = lim - i
	}

	fmt.Println(largeArray)
	fmt.Println("Normal Mergesort")
	r := Mergesort(largeArray)
	fmt.Println(r)
	


	fmt.Println("Mergesort with Goroutines")
	resultChan := make(chan []int, 1)
	MergeSortAsync(largeArray, resultChan)
	k := <- resultChan
	fmt.Println(k)
	
}