package main

import "fmt"

func main() {

	arr := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				fmt.Println(arr[i][j])
			}
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			
			if i+j==2 {
				fmt.Println(arr[i][j])
			}
		}
	}

}