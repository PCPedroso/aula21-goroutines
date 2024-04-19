package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup
var linesChan chan string

func main() {
	linesChan = make(chan string)

	go readFile("./data.txt")

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go processLines()
	}

	wg.Wait()
}

func readFile(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Erro ao abrir arquivo", err)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linesChan <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	close(linesChan)
	wg.Done()
}

func processLines() {
	for line := range linesChan {
		processedLine := processLineAxync(line)
		fmt.Println(processedLine)
		time.Sleep(time.Millisecond * 500)
	}

	wg.Done()
}

func processLineAxync(line string) string {
	return line + " (Processado)"
}
