package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Habit struct {
	Name      string
	Logs      []string
	Priority  int
	CreatedAt string
}

func newHabit(name string, priority int, createdAt string) Habit {
	return Habit{
		Name:      name,
		Priority:  priority,
		CreatedAt: createdAt,
		Logs:      []string{},
	}
}

func saveHabitsToFile(habits []Habit) {
	data, err := json.MarshalIndent(habits, "", "  ")
	if err != nil {
		fmt.Println("Serialization error:", err)
		return
	}

	err = os.WriteFile("habits.json", data, 0644)
	if err != nil {
		fmt.Println("Error while writing to the file:", err)
	}
}

func loadHabitsFromFile() []Habit {
	data, err := os.ReadFile("habits.json")
	if err != nil {
		fmt.Println("Файл не найден. Начинаем с пустого списка.")
		return []Habit{}
	}

	var loadedHabits []Habit
	err = json.Unmarshal(data, &loadedHabits)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return []Habit{}
	}

	return loadedHabits
}
