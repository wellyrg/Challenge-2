package main

import "fmt"

func main() {

	for i := 0; i <= 4; i++ {
		if i < 5 {
			fmt.Printf("Nilai i = %d\n", i)
		}
	}

	for j := 0; j <= 10; j++ {
		if j < 5 {
			fmt.Printf("Nilai j = %d\n", j)
		} else if j == 5 {
			var unicode [13]int32 = [13]int32{'C', 9, 'A', 1, 'Ш', 5, 'A', 7, 'P', 12, 'B', 11, 'O'}
			for k := 0; k < len(unicode); k++ {
				if k%2 == 0 {
					fmt.Printf("character %#U start at byte position %d\n", unicode[k], k)
				}
			}
		} else if j > 5 {
			fmt.Printf("Nilai j = %d\n", j)
		}

	}

}
