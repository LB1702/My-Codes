package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	res := []string{}
	phrase := []rune(args[0])

	if len(args) != 1 {
		return
	}

	for _, i := range phrase {

		startPosition := int(i)
		count := 0
		f, err := os.Open("text.txt")

		if err != nil {
			panic("file not loaded")
		}

		defer f.Close()

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			count++

			if count >= (startPosition-32)*9+2 && count <= (startPosition-32)*9+9 {
				placeChar := scanner.Text()
				res = append(res, placeChar)
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

	}
	result(res, len(phrase))

}

func result(input []string, colonne int) {
	rows := (len(input) + colonne - 1) / colonne

	for row := 0; row < rows; row++ {
		for col := 0; col < colonne; col++ {
			i := col*rows + row
			fmt.Printf("%\b", input[i])
		}
		fmt.Println()
	}
}
