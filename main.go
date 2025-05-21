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

	fmt.Printf("You chose %d", action)

	//TODO - add the logic behind the actions
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
