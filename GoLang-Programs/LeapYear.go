package main

import "fmt"

func leap(n int) {
	if n%4 == 0 {
		if n%100 != 0 {
			fmt.Println("Leap Year")
		} else {
			fmt.Println("Not Leap Year")
		}
	} else {
		fmt.Println("Not Leap Year")
	}
}
func main() {
	leap(2020) // Leap Year
  leap(2021) //Not Leap Year
}
