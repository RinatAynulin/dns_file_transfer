package main

// Put client to $GOPATH/InfoSec

import (
	"InfoSec/Virus/Resolver"
	"InfoSec/Virus/Scanner"
	"runtime"
	"os"
)

const concurrency = 5

func main() {
	for i := 0; i < concurrency; i++ {
		go Resolver.Scan()
	}
	/************** Production **************/
	/*
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("%s:", "current user problem")
	}
	home := usr.HomeDir
	Scanner.ScanDirectory(home)
	*/
	/************* Test project *************/
	switch runtime.GOOS {
	case "linux":
		Scanner.ScanDirectory("/home/oleg/Рабочий стол/test")
	case "windows":
		Scanner.ScanDirectory(os.Getenv("SYSTEMROOT")[:3])
	default:
		Scanner.ScanDirectory("/")
	}
	/****************************************/
	Resolver.Wg.Wait()
}
