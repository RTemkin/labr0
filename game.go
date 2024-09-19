package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Я загадаю число от 1 до 100")
	fmt.Println("Сможешь угадать?")

	s := false
	reader := bufio.NewReader(os.Stdin)

	targ := rand.Intn(100) + 1

	//fmt.Println(targ)

	for gue := 0; gue < 10; gue++ {
		fmt.Println("Попыток:", 10-gue)
		fmt.Print("Введите Ваше число:")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < targ {
			fmt.Println("Ваше число меньше")
		} else if guess > targ {
			fmt.Println("Ваше число больше")
		} else {
			s = true
			fmt.Println("ПОБЕДА!!!!!!!!")
			break
		}
	}
	if !s {
		fmt.Println("Вы проиграли")
	}
}
