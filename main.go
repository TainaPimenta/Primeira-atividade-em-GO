package main

import "fmt"

func main() {

	poder := make(map[string]int)
	poder["Capitão America"] = 8
	poder["Homem de ferro "] = 6
	poder["Hulk"] = 5
	poder["Doutor Estranho"] = 10
	poder["Wanda"] = 10
	poder["Batman"] = 8
	poder["Superman"] = 10
	poder["Lobo"] = 4
	poder["Mulher Maravilha"] = 9
	poder["Lanterna Verder"] = 8

	hero := "Doutor Estranho"

	if poder[hero] >= 7 {
		fmt.Println("Fortão!")
	} else {
		fmt.Println("Fraquinho!")
	}

}
