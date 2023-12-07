package main

import (
	"bufio"
	"fmt"
	gd "main/greedy"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	lines := make([]*gd.Line, 0)
	alocatedUsers := make(map[string]bool)
	userList := make([]*gd.Line, 0)

	n_inputs, err := strconv.Atoi(os.Args[1])
	check(err)

	filename := fmt.Sprintf("bases/final_base%d.csv", n_inputs)
	filename = "bases/paukk.csv"
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	_ = scanner.Scan() // Skip the first line

	for scanner.Scan() {
		line := scanner.Text()
		// Split the line into fields by ,
		fields := strings.Split(line, ",")
		// fmt.Println(fields)
		user := fields[0]
		priority, err := strconv.Atoi(fields[1])
		check(err)
		performance, err := strconv.ParseFloat(fields[2], 32)
		check(err)

		userLine := gd.NewLine(user, priority, float32(performance))

		lines = append(lines, userLine)
		alocatedUsers[user] = false
	}

	comparacoes := 0
	n_attempts := 1000
	initTime := time.Now()
	for i := 0; i < n_attempts; i++ {
		userList, comparacoes = gd.GetUserComp(lines, make(map[string]bool))
	}
	finalTime := time.Now()
	fmt.Println("Tempo de execucao: ", finalTime.Sub(initTime), "para", n_inputs, " elementos", "com", comparacoes, "comparacoes por execucao")
	fmt.Print("UserList: ")
	for _, user := range userList {
		fmt.Print(user, " ")
	}

	return

	// Código de mais de 1 alocação

	sair := false
	for !sair {

		userList, _ = gd.GetUserComp(lines, alocatedUsers)
		fmt.Println("Usuarios alocados por prioridade: ")
		fmt.Println("USER\t\tPRIORITY\tPERFORMANCE")
		for _, user := range userList {
			fmt.Printf("%s\t\t%d\t\t%f\n", user.User, user.Priority, user.Performance)
		}

		// Imprime a performance total
		totalPerformance := float32(0)
		for _, user := range userList {
			totalPerformance += user.Performance
		}
		fmt.Printf("Performance total: %f\n\n", totalPerformance)

		// fmt.Println("Deseja continuar? (s/n)")
		// var input string
		// fmt.Scanln(&input)
		input := "n"
		if input == "n" {
			sair = true
		}

	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
