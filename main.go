package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the To-Do List CLI")
	reader := bufio.NewReader(os.Stdin)
	listMap := make(map[int]string)
	nextID := 1

	for {
		fmt.Printf(">>> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("\nbye.")
			os.Exit(0)
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		cmd := parts[0]
		param := ""
		if len(parts) > 1 {
			param = strings.Join(parts[1:], " ")
		}

		switch cmd {
		case "add":
			listMap[nextID] = param
			fmt.Println("task created with id:", nextID)
			nextID++

		case "list":
			for id, task := range listMap {
				fmt.Println("id", id, "task:", task)
			}
		case "done":
			if param == "" {
				fmt.Println("please provide an id")
			} else {
				id, err := strconv.Atoi(param)
				if err != nil {
					fmt.Println("invalid id:", param)
				} else {
					delete(listMap, id)
					fmt.Println("task with id", id, "removed")
				}
			}
		case "exit":
			fmt.Println("bye.")
			os.Exit(0)
		default:
			fmt.Println("unknown command:", cmd)
		}
	}
}
