package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func ProcessLogs(inputFiles []string, outputFile string) error {
	var wg sync.WaitGroup
	errorChan := make(chan string)
	writerDone := make(chan error)
	go func() {
		outFile, err := os.Create(outputFile)
		if err != nil {
			writerDone <- err
			return
		}
		defer outFile.Close()
		writer := bufio.NewWriter(outFile)
		defer writer.Flush()
		for line := range errorChan {
			_, err := writer.WriteString(line + "\n")
			if err != nil {
				writerDone <- err
				return
			}
		}
		writerDone <- nil
	}()
	for _, file := range inputFiles {
		wg.Add(1)
		go func(fileName string) {
			defer wg.Done()
			f, err := os.Open(fileName)
			if err != nil {
				log.Printf("Failed to open %s: %v", fileName, err)
				return
			}
			defer f.Close()
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				line := scanner.Text()
				if strings.Contains(line, "ERROR") {
					errorChan <- line
				}
			}
			if err := scanner.Err(); err != nil {
				log.Printf("Scanner error in %s: %v", fileName, err)
			}
		}(file)
	}
	go func() {
		wg.Wait()
		close(errorChan)
	}()
	if err := <-writerDone; err != nil {
		return fmt.Errorf("error writing to output file: %w", err)
	}
	return nil
}
func main() {
	inputFiles := []string{"server1.txt", "server2.txt", "server3.txt"}
	err := ProcessLogs(inputFiles, "errors.log")
	if err != nil {
		log.Fatal(err)
	}
}
