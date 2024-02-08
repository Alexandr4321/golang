package main

import (
	"bufio"
	"fmt"
	"os"
    
	"unicode"
)

func main() {
    fmt.Print("Введите предложение: ")
    
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        sentence := scanner.Text()
        countLetters(sentence)
    } else {
        fmt.Println("Ошибка ввода:", scanner.Err())
    }
}

func countLetters(sentence string) {
    letterCount := make(map[rune]int)

    for _, char := range sentence {
        if unicode.IsLetter(char) {
            char = unicode.ToLower(char)
            letterCount[char]++
        }
    }

    totalLetters := 0
    for _, count := range letterCount {
        totalLetters += count
    }

    for letter, count := range letterCount {
        frequency := float64(count) / float64(totalLetters)
        fmt.Printf("%c - %d %.2f%%\n", letter, count, frequency*100)
    }
}
