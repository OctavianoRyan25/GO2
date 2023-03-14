package main

import "fmt"

func main() {
	for i := 0; i <= 4; i++ {
		fmt.Println("Nilai i = ",i)
	}
	
	for j := 0; j <= 4; j++ {
		fmt.Println("Nilai j = ",j)
	}
	var unicode = []rune{'C','A','Ш','A','P','B','O'}
	
	increment := 0;
	for i := 0; i < len(unicode); i++ {
		fmt.Printf("Character %U value %c starts at byte position %d\n", unicode[i], unicode[i], increment)
		increment = increment+2;
	}

	for j := 6; j <= 10; j++ {
		fmt.Println("Nilai j = ",j)
	}
}
