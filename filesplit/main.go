package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	var file, sizeStr string
	var count int
	flag.StringVar(&file, "file", "input.log", "input file")
	flag.StringVar(&sizeStr, "size", "", "size of item (k,m,g)")
	flag.IntVar(&count, "count", 5, "count of item")
	flag.Parse()

	size := parseSize(sizeStr)
	if size == 0 {
		return
	}
	fmt.Printf("file: %s\n", file)
	fmt.Printf("size: %d  count: %d\n", size, count)
	reader, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer reader.Close()

	br := bufio.NewReader(reader)

	if size == 0 {
		return
	}
	for i := 0; i < count; i++ {
		e := func(num int) error {
			w, e1 := os.Create(file + "." + strconv.Itoa(num))
			if e1 != nil {
				return e1
			}
			defer w.Close()
			e2 := copyBuffer(w, br, size)
			if e2 != nil {
				return e2
			}
			return nil
		}(i)
		if e != nil {
			break
		}
	}
}

func parseSize(size string) int {
	l := len(size)
	if l == 0 {
		return 0
	}
	numStr := string(size[0 : l-1])
	unit := size[l-1]
	switch unit {
	case 'k', 'K':
		num, e := strconv.Atoi(numStr)
		if e == nil {
			return num * 1024
		}
	case 'm', 'M':
		num, e := strconv.Atoi(numStr)
		if e == nil {
			return num * 1024 * 1024
		}
	case 'g', 'G':
		num, e := strconv.Atoi(numStr)
		if e == nil {
			return num * 1024 * 1024 * 1024
		}
	default:
		num, e := strconv.Atoi(size)
		if e == nil {
			return num
		}
	}
	fmt.Println("invalid file size")
	return 0
}

func copyBuffer(dst io.Writer, src io.Reader, limit int) error {
	buf := make([]byte, 32*1024)
	var written int64 = 0
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
				if written >= int64(limit) {
					break
				}
			}
			if ew != nil {
				return ew
				break
			}
			if nr != nw {
				return io.ErrShortWrite
				break
			}
		}
		if er != nil {
			return er
		}
	}
	return nil
}
