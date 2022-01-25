package main

var arr = []int{3, 0, 0, 2, 0, 4}
var maxHeight int

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func volume(h int) int {
	return h
}

func main() {
	maxHeight = min(arr[0], arr[len(arr)-1])
	waterCaptured := 0
	for index, value := range arr {
		if index > 0 && index < len(arr) {
			heightLeft := arr[index-1]
			heightRight := arr[index+1]
			if heightLeft > maxHeight {
				heightLeft = maxHeight

			}
			if heightRight > maxHeight {
				heightRight = maxHeight

			}
			waterCaptured += volume(min(heightLeft, heightRight))
		}
	}
}
