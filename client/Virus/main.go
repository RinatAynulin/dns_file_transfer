package main

// Put client to $GOPATH/InfoSec

import (
	"fmt"
	"runtime"
	"InfoSec/Virus/Resolver"
	"InfoSec/Virus/Scanner"
)

const concurrency = 5

func main() {
	for i := 0; i < concurrency; i++ {
		go Resolver.Scan()
	}
	switch runtime.GOOS {
	case "linux":
		Scanner.ScanDirectory("/home/oleg/Рабочий стол/test")
	case "windows":
		fmt.Println("You use windows")
		return
	default:
		Scanner.ScanDirectory("/")
	}
	Resolver.Wg.Wait()
}
