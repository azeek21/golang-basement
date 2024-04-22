package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sync"
	"unicode/utf8"
)

type MODES struct {
	W bool
	L bool
	M bool
}

var (
	NEW_LINE = byte('\n')
	SPACE    = byte(' ')
	MODE     MODES
)

func count(src []byte, srclen int, tar byte, skipConsequent bool) int64 {
	srclen--
	count := int64(0)
	for srclen > 0 {
		srclen--
		if src[srclen] == tar {
			count++
			if skipConsequent {
				for src[srclen] == tar {
					srclen--
				}
			}
		}
	}
	return count
}

func countLine(src []byte) int64 {
	return count(src, len(src), NEW_LINE, false) + 1
}

func countWord(src []byte) int64 {
	return count(src, len(src), SPACE, true) + 1
}

func countChars(src []byte) int64 {
	return int64(utf8.RuneCount(src))
}

func Wc(fileName string, ch chan [2]string, wg *sync.WaitGroup) {
	file, err := os.Open(fileName)
	res := int64(0)
	defer wg.Done()
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	buf := make([]byte, 16*1024)

	for {
		n, err := file.Read(buf)

		if n > 0 {
			if MODE.L {
				res += countLine(buf[:n])
			} else if MODE.W {
				res += countWord(buf[:n])
			} else if MODE.M {
				res += countChars(buf[:n])
			}
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("read %d bytes: %v", n, err)
			break
		}
	}
	ch <- [2]string{fmt.Sprintf("%d", res), fileName}
}

func init() {
	wc := flag.Bool("wc", false, "count number of word in a file")
	l := flag.Bool("l", false, "count number of lines in a file")
	m := flag.Bool("m", false, "count number of characters in a file")
	flag.Parse()
	count := 0

	if *l {
		count++
	}
	if *m {
		count++
	}
	if count == 0 {
		*wc = true
		count++
	}

	if count > 1 {
		fmt.Printf("error [options]: only one option can be used at a time\n")
		flag.Usage()
		os.Exit(1)
	}

	MODE = MODES{W: *wc, L: *l, M: *m}
}
