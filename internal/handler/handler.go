package handler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"binary-tree-go/internal/models" // tree.go のパッケージ名
)

func RunCLI() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Binary Tree CLI. Type 'exit' to quit.")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		args := strings.Fields(line)
		cmd := strings.ToLower(args[0])

		switch cmd {
		case "insert":
			if len(args) != 2 {
				fmt.Println("usage: insert x")
				continue
			}
			key, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("invalid number")
				continue
			}
			err = models.AddNode(key)
			if err != nil {
				fmt.Printf("%d is already exists\n", key)
			} else {
				fmt.Printf("inserted %d\n", key)
			}
		case "get":
			if len(args) != 2 {
				fmt.Println("usage: get x")
				continue
			}
			key, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("invalid number")
				continue
			}
			found := models.FindNode(key)
			if found {
				fmt.Printf("%d found\n", key)
			} else {
				fmt.Printf("%d not found\n", key)
			}
		case "delete":
			if len(args) != 2 {
				fmt.Println("usage: delete x")
				continue
			}
			key, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("invalid number")
				continue
			}
			err = models.DeleteNode(key)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("%d deleted\n", key)
			}
		case "show":
			models.ShowTree()
		case "exit":
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("unknown command")
		}
	}
}
