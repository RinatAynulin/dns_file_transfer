package Collector

import (
	"sync"
	"InfoSec/FilesCollector/Scanner"
	"fmt"
	"bufio"
	"os"
	"strings"
	"encoding/hex"
	"InfoSec/Params"
)

// WG is wait group to synchronize the pool and the main routine
var Wg sync.WaitGroup

// Scan reads file paths from channel and sends it
func Collect(input string) {
	Wg.Add(1)
	inf, _ := os.Open(input)
	defer inf.Close()

	for file := range Scanner.Files {
		scanner := bufio.NewScanner(inf)
		for scanner.Scan() {
			var id, payload string
			var offset int64
			data := scanner.Text()
			data = data[20: len(data) - Params.UrlLen]
			fmt.Sscanf(data, "%56s.%019x.%s", &id, &offset, &payload);
			// It's new file
			if file.Id == id && offset != 0 {
				fmt.Println(offset, payload)
				payload = strings.Replace(payload, ".", "", -1)
				offset -= 1
				bytes, _ := hex.DecodeString(payload)
				_, err := file.Ptr.WriteAt(bytes, offset)
				fmt.Println(err)
			}
		}
		file.Ptr.Close()
	}
	Wg.Done()
}
