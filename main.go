package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed test/version.txt
var version string
//go:embed test/KT.jpg
var logo []byte
//go:embed test/files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)
	
	err:= ioutil.WriteFile("logo_new.jpg",logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dir, _ := path.ReadDir("test/files")
	for _, entry := range dir {
		if !entry.IsDir(){
			file, _ := path.ReadFile("test/files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}