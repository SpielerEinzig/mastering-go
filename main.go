package main

import "fmt"

func main() {
	////1: Arrays
	////arrays with fixed length
	// var ages [3]int = [3]int {1, 2, 3}
	// fmt.Println(ages)

	////short hand
	// ageList := [3]int {3, 2, 1}
	// fmt.Println(ageList)

	////print out length of array
	// fmt.Println("'ages' array is ", len(ages), "long")
	// fmt.Println("'ageList' array is ", len(ageList), "long")


	//2: Slices
	//slices are basically an array without a fixed length
	// var slice = []int{5, 6, 7, 8}
	// fmt.Println(slice)
	// //change item at an index (works for arrays as well)
	// slice[2] = 12
	// fmt.Println(slice)
	// //add item to the slice (works for array as well)
	// slice = append(slice, 56)
	// fmt.Println(slice)


	//3. Ranges
	//Basically getting a range from one array
	//for example index 2 to 5 within an array or slice
	//Also works for slices

	var nameList = []string {"Esdeath", "Revy", "Balalaika"}
	var rangeList = nameList[1:]
	fmt.Println(rangeList)

} 