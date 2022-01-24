package main

import "fmt"

var chessBoard [8][8]string

func checkOccupied(rank int) int {
	count := 0
	for _, value := range chessBoard[rank] {
		if value == "#" {
			count++
		}
	}
	return count
}
func checkOccupiedFiles(files int) int {
	count := 0

	for index1, _ := range chessBoard {
		if chessBoard[index1][file] == "#" {
			count++
		}
	}
	return count
}
func main() {
	chessBoard = [8][8]string{
		{"#", "_", "#", "_", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "#", "_", "_", "_"},
		{"_", "_", "#", "_", "_", "_", "_", "_"},
		{"#", "_", "_", "_", "#", "_", "_", "_"},
		{"#", "_", "_", "_", "_", "_", "_", "_"},
		{"_", "#", "#", "_", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_", "_", "_", "#"},
		{"_", "_", "#", "_", "_", "_", "_", "_"}}

	fmt.Println(checkOccupied(0))
}
