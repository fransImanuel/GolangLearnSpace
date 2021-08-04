package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T){
	fmt.Println(version)
}

//go:embed images/doge.png
var logo []byte

func TestByte(t *testing.T){
	err := ioutil.WriteFile("doge_new2.png", logo, fs.ModePerm)
	if err!=nil {
		panic(err)
	}
	// err := ioutil.WriteFile("doge_new.png", logo, fs.ModePerm)
	// if err!=nil {
	// 	panic(err)
	// }
}

//go:embed file/a.txt
//go:embed file/B.TXT
//go:embed file/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T){
	a, _ :=files.ReadFile("file/a.txt")
	fmt.Println(string(a))

	b, _ :=files.ReadFile("file/B.TXT")
	fmt.Println(string(b))

	c, _ :=files.ReadFile("file/c.txt")
	fmt.Println(string(c))
}

//go:embed file/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T){
	dirEntries,_ := path.ReadDir("file")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ :=path.ReadFile("file/"+entry.Name())
			fmt.Println(string(file))
		}
	}

}
