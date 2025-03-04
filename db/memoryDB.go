package db

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/buskarion/todoapp-with-gin/entity"
)

var todoDB *[]entity.Todo

func BuildDB() *[]entity.Todo {
	if todoDB != nil {
		return todoDB
	}

	file, err := os.Open("../../db/resource.csv")
	if err != nil {
		log.Fatalf("Could not read from source file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var todos []entity.Todo
	for i, record := range records {
		if i == 0 {
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("Error casting ID \"%s\" to int: %v", record[0], err)
			continue
		}

		completed, err := strconv.ParseBool(record[2])
		if err != nil {
			log.Printf("Error casting filed Completed \"%s\" to boolean: %v", record[2], err)
			continue
		}

		todo := entity.Todo{
			ID:        id,
			Task:      record[1],
			Completed: completed,
		}

		todos = append(todos, todo)
	}

	fmt.Printf("Loaded %d ToDos from CSV\n", len(todos))
	return &todos
}
