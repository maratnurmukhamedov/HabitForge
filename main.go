package main

import (
	"fmt"
	"time"
)

var habits []Habit

func main() {
	fmt.Println("Welcome to HabitForge - your habit, your forge.")

	habits = loadHabitsFromFile()

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

		var name string
		var priority int

		fmt.Print("Enter habit name: ")
		fmt.Scan(&name)

		fmt.Print("Enter priority (1-5): ")
		fmt.Scan(&priority)

		createdAt := time.Now().Format("2006-01-02") // make the realtime.
		newHabit := NewHabit(name, priority, createdAt)
		fmt.Println(newHabit)

		habits = append(habits, newHabit)
		fmt.Println("Habit added successfully!")
		saveHabitsToFile(habits)

	case 2:
		fmt.Println("Adding progress to the habit")
		for i, h := range habits {
			fmt.Printf("%d. %s\n", i+1, h.Name)
		}

		var choice int
		fmt.Print("Enter habit number: ")
		fmt.Scan(&choice)

		if choice < 1 || choice > len(habits) {
			fmt.Println("Invalid choice")
			break
		}

		currentDate := time.Now().Format("2006-01-02")
		habits[choice-1].Logs = append(habits[choice-1].Logs, currentDate)

		fmt.Printf("Logged progress for \"%s\" on %s\n", habits[choice-1].Name, currentDate)

	case 3:
		fmt.Println(" Check all habits")

		fmt.Println("All habits:")
		for i, h := range habits {
			fmt.Printf("%d. %s (Priority: %d, Created: %s)\n", i+1, h.Name, h.Priority, h.CreatedAt)
		}
	case 4:
		fmt.Println(" Showing your longest streaks")
	case 0:
		fmt.Println("Quitting")
	default:
		fmt.Println("Incorrect input")
	}

}
