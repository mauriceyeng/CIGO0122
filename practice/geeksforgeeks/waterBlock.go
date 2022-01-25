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
		var waterHeight int
		if value > maxHeight {
			value = maxHeight
		}
		if index > 0 && index < len(arr)-1 {
			if arr[index-1] > maxHeight && arr[index+1] > maxHeight {
				waterHeight = min(arr[index-1], arr[index+1])
			} else {
				waterHeight = maxHeight - value
			}
		}

	}
	waterCaptured += volume(waterCaptured)
}
