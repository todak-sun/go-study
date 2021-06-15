package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix() // 현재 날짜/시간을 정숫값으로 가져온다.
	rand.Seed(seconds)           // 난수를 발생시키기 위한 Seed.
	target := rand.Intn(100) + 1

	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)

	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left")
		fmt.Print("Make a guess : ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops. Youre guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH")
		} else {
			fmt.Println("Good job! You guessed it!")
			break
		}
	}
}
