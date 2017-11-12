package main

import (
	"InfoSec/Virus/Scaner"
	"InfoSec/Virus/Resolver"
	"fmt"
	"runtime"
)

const concurrency = 5

func main() {
	for i := 0; i < concurrency; i++ {
		go Resolver.Scan()
	}
	switch runtime.GOOS {
	case "linux":
		Scaner.ScanDirectory("/home/oleg/Рабочий стол/test")
	case "windows":
		fmt.Println("You use windows")
		return
	default:
		Scaner.ScanDirectory("/")
	}
	Resolver.Wg.Wait()
}
