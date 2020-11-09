package main

/*
*@author Isac Canedo
 */

import "fmt"

func main() {

	loop()
}

func loop() {

	for i := 1; i <= 5; i++ {
		for j := i + 1; j <= 5; j++ {

			fmt.Println("[", i, ",", j, "]")

		}

	}

}
