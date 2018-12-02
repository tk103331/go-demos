package main

import (
	"archive/tar"
	"compress/gzip"
	"compress/lzw"
	"compress/zlib"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

type method int

const (
	no method = iota
	gz
	bz
	lz
	zz
)

var m = no

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
	} else if mode == "-cg" {
		fmt.Println("compress gzip " + src + " to " + dst)
		m = gz
		compress(src, dst)
	} else if mode == "-cl" {
		fmt.Println("compress lzw " + src + " to " + dst)
		m = lz
		compress(src, dst)
	} else if mode == "-cz" {
		fmt.Println("compress zlib " + src + " to " + dst)
		m = zz
		compress(src, dst)
	} else if mode == "-x" {
		fmt.Println("extract " + src + " to " + dst)
		extract(src, dst)
	} else if mode == "-xg" {
		fmt.Println("extract gzip " + src + " to " + dst)
		m = gz
		extract(src, dst)
	} else if mode == "-xl" {
		fmt.Println("extract bzip2 " + src + " to " + dst)
		m = lz
		extract(src, dst)
	} else if mode == "-xz" {
		fmt.Println("extract lzw " + src + " to " + dst)
		m = zz
		extract(src, dst)
	} else {
		printUsage()
	}

}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcompress : mytar -c src dst")
	fmt.Println("\tcompress gzip : mytar -cg src dst")
	fmt.Println("\tcompress lzw : mytar -cl src dst")
	fmt.Println("\tcompress zlib : mytar -cz src dst")
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
	var tw *tar.Writer

	switch m {
	case gz:
		gw := gzip.NewWriter(dstFile)
		tw = tar.NewWriter(gw)
	case lz:
		lw := lzw.NewWriter(dstFile, lzw.MSB, 8)
		tw = tar.NewWriter(lw)
	case zz:
		zw := zlib.NewWriter(dstFile)
		tw = tar.NewWriter(zw)
	default:
		tw = tar.NewWriter(dstFile)
	}

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
	var r io.Reader
	switch m {
	case gz:
		r, err = gzip.NewReader(f)
		check(err)
	case lz:
		r = lzw.NewReader(f, lzw.MSB, 8)
	case zz:
		r, err = zlib.NewReader(f)
		check(err)
	default:
		r = f
	}

	zr := tar.NewReader(r)

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
