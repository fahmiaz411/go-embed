package test

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string
//go:embed version.txt
var version2 string

func TestString(t *testing.T){
	fmt.Println(version)
	fmt.Println(version2)
}

//go:embed KT.jpg
var logo []byte

func TestByte(t *testing.T){
	err := ioutil.WriteFile("logo_new.jpg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS
func TestMultipleFiles (t *testing.T){
 	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))
 	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))
 	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

//go:embed files/*.txt
var path embed.FS
func TestPathMatcher (t *testing.T){
	dir, _ :=path.ReadDir("files")
	for _, entry := range dir {
		if !entry.IsDir() {
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))

		}
	}
}