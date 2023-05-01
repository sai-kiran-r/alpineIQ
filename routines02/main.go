package main

import (
	"fmt"
	"sync"
)

// - Insert into the "output" var all Nike + Adidas + Puma product IDs that start with the digit "1"
// - Do not change the the 3 Get functions in any way
// - Utilize go routines (i.e. you will want to be calling the Get functions and inserting into "output" concurrently)

func main() {
	var output []string
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		nike := GetNikeProductIDs()
		for _, id := range nike {
			if id[0] == '1' {
				output = append(output, id)
			}
		}
	}()

	go func() {
		defer wg.Done()
		adidas := GetAdidasProductIDs()
		for _, id := range adidas {
			if id[0] == '1' {
				output = append(output, id)
			}
		}
	}()

	go func() {
		defer wg.Done()
		puma := GetPumaProductIDs()
		for _, id := range puma {
			if id[0] == '1' {
				output = append(output, id)
			}
		}
	}()

	wg.Wait()

	fmt.Println(output)
}

func GetNikeProductIDs() []string {
	var out []string
	for i := 0; i < 100; i += 10 {
		out = append(out, fmt.Sprintf("%d_%s", i, "nike"))
	}
	return out
}

func GetAdidasProductIDs() []string {
	var out []string
	for i := 0; i < 100; i += 5 {
		out = append(out, fmt.Sprintf("%d_%s", i, "adidas"))
	}
	return out
}

func GetPumaProductIDs() []string {
	var out []string
	for i := 0; i < 100; i += 2 {
		out = append(out, fmt.Sprintf("%d_%s", i, "puma"))
	}
	return out
}
