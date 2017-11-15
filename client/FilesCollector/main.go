package main

import (
	"os"
	"log"
	"InfoSec/FilesCollector/Collector"
	"InfoSec/FilesCollector/Scanner"
)

const concurrency = 5

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Enter input file name")
	}
	input := os.Args[1]
	file, err := os.Open(input)
	if err != nil {
		log.Fatalf("File error: %s", err)
	}
	defer file.Close()

	for i := 0; i < concurrency; i++ {
		go Collector.Collect(input)
	}

	Scanner.Scan(file)
	Collector.Wg.Wait()
}
