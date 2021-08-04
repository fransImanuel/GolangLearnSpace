package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed images/doge.png
var logo []byte

//go:embed file/*.txt
var path embed.FS

func main(){
	fmt.Println(version)

	err := ioutil.WriteFile("doge_new3.png", logo, fs.ModePerm)
	if err!=nil {
		panic(err)
	}

	dirEntries,_ := path.ReadDir("file")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ :=path.ReadFile("file/"+entry.Name())
			fmt.Println(string(file))
		}
	}
}