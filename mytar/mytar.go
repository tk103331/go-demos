package main

import (
	"archive/tar"
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
	fmt.Println("\tcompress Store: mytar -cs src dst")
	fmt.Println("\tcompress Deflate: mytar -cd src dst")
	fmt.Println("\textract  : mytar -x src dst")
}

func compress(src string, dst string) {
	srcFile, err := os.Open(src)
	check(err)
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
	}
	tw := tar.NewWriter(dstFile)

	write(tw, srcFile, "")

	tw.Flush()
	tw.Close()
}

func write(tw *tar.Writer, file *os.File, path string) {
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
			write(tw, f, name+"/")
		}
	} else {
		fih, err := tar.FileInfoHeader(stat, "")
		check(err)
		fih.Name = name
		err = tw.WriteHeader(fih)
		check(err)
		io.Copy(tw, file)
	}
}

func extract(src string, dst string) {
	f, err := os.Open(src)
	check(err)
	zr := tar.NewReader(f)

	for err == nil {
		zh, err := zr.Next()
		if err == io.EOF {
			break
		}
		check(err)
		name := dst + zh.Name
		dir := path.Dir(name)
		_, err = os.Stat(dir)
		if err != os.ErrNotExist {
			os.MkdirAll(dir, os.ModeDir)
		}
		f, err := os.Create(name)
		check(err)
		io.Copy(f, zr)
		f.Close()
		os.Chtimes(name, zh.ModTime, zh.ModTime)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
