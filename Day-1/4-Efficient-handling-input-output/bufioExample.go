package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Bacon ipsum dolor amet porchetta short ribs short loin, spare ribs t-bone kielbasa bresaola ")
	fmt.Fprint(w, "tail ribeye pastrami flank doner. Turducken shankle kevin, landjaeger rump bresaola \n")
	// Don't forget to flush!
	w.Flush()

	numberInput := "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36"

	intScanner := bufio.NewScanner(strings.NewReader(numberInput))

	// custom split by comma function
	splitByComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if i:= strings.IndexRune(string(data), ','); i>= 0 {
			return i + 1, data[0:i], nil
		}

		if atEOF {
			return len(data), data, nil
		}
		return
	}

	intScanner.Split(splitByComma)
	buf := make([]byte, 2)

	// Scan 2 bytes at the time
	intScanner.Buffer(buf, bufio.MaxScanTokenSize)
	for intScanner.Scan() {
		fmt.Printf("%s", intScanner.Text())
	}
}
