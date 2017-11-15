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
	Id string
	Ptr *os.File
}

var Files = make(chan File)

func Scan(path *os.File) {
	scanner := bufio.NewScanner(path)
	for scanner.Scan() {
		var id, payload string
		var offset int64
		data := scanner.Text()
		data = data[20 : len(data) - Params.UrlLen]
		_, err := fmt.Sscanf(data,"%56s.%019x.%s", &id, &offset, &payload);
		if err != nil {
			continue
		}
		// It's new file
		if offset == 0 {
			params := strings.Split(payload, ".")
			filename, _ := hex.DecodeString(params[0])
			sizeStr, _ := hex.DecodeString(params[1])
			size, _ := strconv.ParseInt(string(sizeStr), 10, 64)
			filePtr, _ := os.Create(fmt.Sprintf("%s_%s", id, filename))
			filePtr.Truncate(size)
			Files <- File{id, filePtr}
		}
		//payload = strings.Replace(payload, ".", "", -1)
	}
	defer close(Files)
}
