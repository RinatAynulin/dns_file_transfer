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
	"log"
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
			var payload string
			var offset int64
			var id uint64
			data := scanner.Text()
			data = data[20: len(data) - Params.UrlLen]
			fmt.Sscanf(data, "%016x.%016x.%s", &id, &offset, &payload);
			// It's not new file
			if file.Id == id && offset != 0 {
				payload = strings.Replace(payload, ".", "", -1)
				offset -= 1
				bytes, _ := hex.DecodeString(payload)
				if _, err := file.Ptr.WriteAt(bytes, offset); err != nil {
					log.Println(err)
				}
			}
		}
		file.Ptr.Close()
	}
	Wg.Done()
}
