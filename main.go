package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const dataFile = "./data.txt"

func main() {
	// get the arguments passed to the program
	args := os.Args[1:]

	if len(args) > 2 {
		log.Fatal("too many arguments passed")
		return
	}

	// if no args passed
	if len(args) == 0 {
		getTasks()
		return
	}

	if len(args) == 1 {
		switch args[0] {
		case "list", "get":
			{
				getTasks()
			}
		case "clear":
			{
				{
					err := deleteAllFromFile(dataFile)
					if err != nil {
						log.Fatal("error deleting all from data file: ", err)
					}
					fmt.Println("all tasks cleared")
				}
			}
		default:
			{
				log.Fatal("unsupported argument: " + args[0])
			}
		}
		return
	}

	if len(args) == 2 {
		switch args[0] {
		case "add":
			{
				if args[1] == "" {
					log.Fatal("task cannot be empty")
				}
				addTask(args)
			}

		case "delete", "remove":
			{
				if args[1] == "" {
					log.Fatal("task index cannot be empty")
				}
				deleteTask(args)
			}
		default:
			{
				log.Fatal("unsupported argument: " + args[0])

			}
		}
		return
	}
}

func getTasks() {
	tasks, err := readFromFile(dataFile)

	if err != nil {
		log.Fatal("error reading data file: ", err)
	}

	for index, task := range tasks {
		fmt.Printf("[%v] %v\n", index, task)
	}
}

func addTask(args []string) {
	err := writeToFile(dataFile, args[1])
	if err != nil {
		log.Fatal("error writing to data file: ", err)
	}
}

func deleteTask(args []string) {
	if args[1] == "*" {
		err := deleteAllFromFile(dataFile)
		if err != nil {
			log.Fatal("error deleting all from data file: ", err)
		}
	} else {
		index, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal("invalid index provided: ", err)
		}
		err = deleteFromFile(dataFile, index)
		if err != nil {
			log.Fatal("error deleting from data file: ", err)
		}
	}
}

func writeToFile(filePath string, data string) error {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	_, err = f.WriteString(data + "\n")
	if err != nil {
		return err
	}

	return f.Close()
}

func readFromFile(filePath string) ([]string, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return []string{}, err
	}

	data := string(bytes)

	if len(data) == 0 {
		return []string{}, nil
	} else {
		data = strings.TrimSuffix(data, "\n")
		return strings.Split(data, "\n"), nil
	}
}

func deleteAllFromFile(filePath string) error {
	_, err := os.Create(filePath)
	return err
}

func deleteFromFile(filePath string, index int) error {
	lines, err := readFromFile(filePath)
	if err != nil {
		return err
	}

	newLines, err := removeAtIndex(lines, index)
	if err != nil {
		return err
	}

	data := strings.Join(newLines, "\n")

	err = os.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		return err
	}

	return nil
}

func removeAtIndex(data []string, index int) ([]string, error) {
	if index < 0 || index >= len(data) {
		return nil, errors.New("index out of range: " + strconv.Itoa(index))
	}

	return append(data[:index], data[index+1:]...), nil
}
