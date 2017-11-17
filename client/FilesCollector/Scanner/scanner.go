package Scanner

import (
	"bufio"
	"fmt"
	"strings"
	"encoding/hex"
	"strconv"
	"os"
	"InfoSec/Params"
)

type File struct {
	Id uint64
	Ptr *os.File
}

var Files = make(chan File)

func Scan(path *os.File) {
	scanner := bufio.NewScanner(path)
	for scanner.Scan() {
		var payload string
		var id uint64
		var offset int64
		data := scanner.Text()
		data = data[20 : len(data) - Params.UrlLen]
		_, err := fmt.Sscanf(data,"%016x.%016x.%s", &id, &offset, &payload);
		if err != nil {
			continue
		}
		// It's new file
		if offset == 0 {
			params := strings.Split(payload, ".")
			filename, _ := hex.DecodeString(params[0])
			sizeStr, _ := hex.DecodeString(params[1])
			size, _ := strconv.ParseInt(string(sizeStr), 10, 64)
			filePtr, _ := os.Create(fmt.Sprintf("%x_%s", id, filename))
			filePtr.Truncate(size)
			Files <- File{id, filePtr}
		}
	}
	defer close(Files)
}
