package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Структура для представления URL
type Item struct {
	Name string
	Date time.Time
	Tags string
	Link string
}

func main() {
	var items []Item

	fmt.Println("Коллекция URL")
	fmt.Println("Введите 'A' для добавления, 'D' для удаления, 'L' для вывода списка, 'Q' для выхода")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Введите ваш выбор: ")
		scanner.Scan()
		input := scanner.Text()

		switch strings.ToUpper(input) {
		case "A":
			addItem(&items)
		case "D":
			deleteItem(&items)
		case "L":
			listItems(items)
		case "Q":
			fmt.Println("Выход...")
			os.Exit(0)
		default:
			fmt.Println("Неверная опция. Пожалуйста, попробуйте еще раз.")
		}
	}
}

// Функция для добавления нового URL
func addItem(items *[]Item) {
	fmt.Println("Добавление нового URL:")

	name := getInput("Введите название: ")
	tags := getInput("Введите теги (через запятую): ")
	link := getInput("Введите URL: ")

	newItem := Item{
		Name: name,
		Date: time.Now(),
		Tags: tags,
		Link: link,
	}

	*items = append(*items, newItem)

	fmt.Println("URL успешно добавлен!")
}

// Функция для удаления URL
func deleteItem(items *[]Item) {
	if len(*items) == 0 {
		fmt.Println("Нет URL для удаления.")
		return
	}

	fmt.Println("Выберите индекс для удаления:")
	listItems(*items)

	index := getIndex(len(*items))
	*items = append((*items)[:index], (*items)[index+1:]...)

	fmt.Println("URL успешно удален!")
}

// Функция для вывода списка URL
func listItems(items []Item) {
	if len(items) == 0 {
		fmt.Println("Нет URL в коллекции.")
		return
	}

	fmt.Println("Коллекция URL:")
	for i, item := range items {
		fmt.Printf("%d. %s - Теги: %s - Ссылка: %s\n", i+1, item.Name, item.Tags, item.Link)
	}
}

// Функция для получения ввода от пользователя
func getInput(prompt string) string {
	fmt.Print(prompt + " ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// Функция для получения индекса элемента для удаления
func getIndex(maxIndex int) int {
	var index int

	for {
		indexInput := getInput("Введите индекс: ")
		_, err := fmt.Sscanf(indexInput, "%d", &index)

		if err == nil && index > 0 && index <= maxIndex {
			break
		}

		fmt.Println("Неверный индекс. Пожалуйста, попробуйте еще раз.")
	}

	return index - 1
}

