package main

//https://www.hackerrank.com/challenges/picking-numbers/problem?isFullScreen=false

func pickingNumbers(a []int32) int32 {
	freq := make(map[int32]int32)
	for _, num := range a {
		freq[num]++
	}

	var max int32 = 0
	for num := range freq {
		if freq[num]+freq[num+1] > max {
			max = freq[num] + freq[num+1]
		}
	}

	return max
}

func main() {

	var a []int32
	a = append(a, 11, 10, 10, 10, 10, 10)

	result := pickingNumbers(a)

	println(result)
}
