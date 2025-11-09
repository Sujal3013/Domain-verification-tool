package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func setupLogger() {
	logFile, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Could not open log file: %v\n", err)
		os.Exit(1)
	}
	log.SetOutput(logFile)

}

func main() {
	setupLogger()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter domain address to check : \n")
	for scanner.Scan() {
		domain := strings.TrimSpace(scanner.Text())
		if domain == "" {
			continue
		}
		if domain == "quit" {
			fmt.Println("Exiting...")
			break
		}
		CheckDomain(domain)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input domain: %v", err)
	}
}
