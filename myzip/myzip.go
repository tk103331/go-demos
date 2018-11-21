package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

func main() {
	if len(os.Args) < 4 {
		printUsage()
		os.Exit(1)
	}

	mode := os.Args[1]
	src := os.Args[2]
	dst := os.Args[3]

	if mode == "-c" {
		fmt.Println("compress " + src + " to " + dst)
		compress(src, dst)
	} else if mode == "-x" {
		fmt.Println("extract " + src + " to " + dst)
		extract(src, dst)
	} else {
		printUsage()
	}

}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcompress : myzip -c src dst")
	fmt.Println("\textract  : myzip -x src dst")
}

func compress(src string, dst string) {
	srcFile, err := os.Open(src)
	check(err)
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
	}
	zw := zip.NewWriter(dstFile)

	write(zw, srcFile, "")

	zw.Flush()
	zw.Close()
}

func write(zw *zip.Writer, file *os.File, path string) {
	stat, err := file.Stat()
	check(err)
	name := path + stat.Name()
	if stat.IsDir() {
		fs, err := file.Readdir(0)
		check(err)
		for _, fi := range fs {
			f, err := os.Open(name + "/" + fi.Name())
			check(err)
			defer f.Close()
			write(zw, f, name+"/")
		}
	} else {
		zf, err := zw.Create(name)
		check(err)
		io.Copy(zf, file)
	}
}

func extract(src string, dst string) {
	srcFile, err := os.Open(src)
	check(err)
	defer srcFile.Close()
	stat, err := srcFile.Stat()
	check(err)
	zr, err := zip.NewReader(srcFile, stat.Size())
	check(err)
	for _, zf := range zr.File {
		if zf.FileInfo().IsDir() {
			os.MkdirAll(zf.Name, os.ModeDir)
		} else {
			dir := dst + path.Dir(zf.Name)
			_, err := os.Stat(dir)
			if err != os.ErrNotExist {
				os.MkdirAll(dir, os.ModeDir)
			}
			r, err := zf.Open()
			check(err)
			defer r.Close()
			f, err := os.Create(dst + zf.Name)
			check(err)
			io.Copy(f, r)
			f.Close()
		}

	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
