package main

import (
	"client/Virus/Resolver"
	"client/Virus/Scaner"
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
