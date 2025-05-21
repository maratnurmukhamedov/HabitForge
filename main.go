package main

import "fmt"

func main() {
	fmt.Println("Welcome to HabitForge - your habit, your forge.")

	var action int

	fmt.Println("Please choose your action - type number")
	fmt.Println("1. Add new habit")
	fmt.Println("2. Describe habit progress")
	fmt.Println("3. See all habits")
	fmt.Println("4. Check your streak")
	fmt.Println("0. Quit")

	fmt.Scan(&action)

	fmt.Printf("You chose %d\n", action)

	//TODO - add the logic behind the actions
	switch action {
	case 1:
		fmt.Println("Adding new habit")
	case 2:
		fmt.Println("Adding progress to the habit")
	case 3:
		fmt.Println(" Check all habits")
	case 4:
		fmt.Println(" Showing your longest streaks")
	case 0:
		fmt.Println("Quitting")
	default:
		fmt.Println("Incorrect input")
	}
}

/*
=========================
      HABITFORGE
=========================

Выбери действие:
1. Добавить привычку
2. Логировать выполнение
3. Посмотреть все привычки
4. Посмотреть streaks
0. Выход

Введите номер действия:
*/
