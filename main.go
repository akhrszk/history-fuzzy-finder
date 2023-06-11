package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"uiapp/app"
)

func main() {
	history := readHistoryFile()
	app := app.NewApp(history)

	go func() {
		cmd := <-app.Cmd
		os.Stdout.Write([]byte(cmd))
		app.Stop()
	}()

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func readHistoryFile() []string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	file, err := os.Open(filepath.Join(home, ".zsh_history"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ";")
		if len(line) == 2 {
			lines = append(lines, line[1])
		}
	}
	return reverse(lines)
}

func reverse(slice []string) []string {
	for i := len(slice) / 2; i >= 0; i-- {
		opp := len(slice) - 1 - i
		slice[i], slice[opp] = slice[opp], slice[i]
	}
	return slice
}
