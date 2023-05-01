package main

import (
	"fmt"
)

// - Insert into the "output" var all Nike + Adidas + Puma product IDs that start with the digit "1"
// - Do not change the the 3 Get functions in any way
// - Utilize go routines (i.e. you will want to be calling the Get functions and inserting into "output" concurrently)

func main() {

	// declare single channels for the product
	singleChannel := make(chan []string)

	// fork data from respective functions to declared channels - concurrently
	go func() {
		singleChannel <- GetNikeProductIDs()
		singleChannel <- GetAdidasProductIDs()
		singleChannel <- GetPumaProductIDs()
	}()

	var output []string
	for i := 0; i < 3; i++ {
		select {
		case single := <-singleChannel:
			for _, id := range single {
				if id[0] == '1' {
					output = append(output, id)
				}
			}
		}
	}

	// Printing the output
	fmt.Println(output)
}

func GetNikeProductIDs() (out []string) {
	for i := 0; i < 100; i += 10 {
		out = append(out, fmt.Sprintf("%d_%s", i, "nike"))
	}
	//fmt.Println("nike")
	return out
}

func GetAdidasProductIDs() (out []string) {
	for i := 0; i < 100; i += 5 {
		out = append(out, fmt.Sprintf("%d_%s", i, "adidas"))
	}
	//fmt.Println("adidas")
	return out
}

func GetPumaProductIDs() (out []string) {
	for i := 0; i < 100; i += 2 {
		out = append(out, fmt.Sprintf("%d_%s", i, "puma"))
	}
	//fmt.Println("puma")
	return out
}
