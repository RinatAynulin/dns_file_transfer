package main

// Put client to $GOPATH/InfoSec

import (
	"fmt"
	"runtime"
	"InfoSec/Virus/Resolver"
	"InfoSec/Virus/Scaner"
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
