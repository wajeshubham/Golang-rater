package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var name string
	var userRating int64

	//take user's name
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your full name: ")
	name, _ = reader.ReadString('\n')

	//take rating
	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Please rate our sketches betweemn 1-5: ")
	rating, _ := reader.ReadString('\n')
	userRating, _ = strconv.ParseInt(strings.TrimSpace(rating), 0, 64) // convert the string type rating into Int64

	// to keep rating between 1-5
	if userRating < 5 && userRating > 0 {
		fmt.Printf("Hello %v,\nThanks for rating our sketches with %v star rating.\n", strings.TrimSpace(name), userRating)
		fmt.Printf("Your rating was recorded in our system at %v \n\n", time.Now().Format(time.Stamp))

		if userRating == 5 {
			fmt.Println("Thank for 5 star rating!")
		} else if userRating == 4 || userRating == 3 {
			fmt.Println("We will try to improve.")
		} else if userRating < 3 {
			fmt.Println("Need serious work on our side")
		}
	} else {
		fmt.Printf("\nPlease rate our sketches between 1-5 you entered %v \n\n", userRating)

		// if user enters wrong rating it will again run the main function.
		main()
	}

}
