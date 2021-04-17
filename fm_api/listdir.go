package main

import (
	"fmt"
	"io/fs"

	//"io/fs"
	"log"
	"os"
	//"time"
)

type filesStruct struct {
	Name string
	Mode string
}

func main() {
	//var fl []fs.DirEntry
	var fl3 []filesStruct
	//listDir("/home/mks/git", "", 3, 0)
	fl3 = listDir3("/home/mks/git")
	fmt.Println(fl3)
}

func listDir(dir, depthDelim string, maxdepth, curDepth uint8) {
	// Print dirTree
	//var flist []string
	curDepth += 1
	files, err := os.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		// fm, err := os.Lstat(dir + "/" + f.Name())
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Printf("%v  %v  %v\n", f.Name(), f.IsDir(), fm.Mode().String())
		if f.IsDir() && curDepth <= maxdepth {
			fmt.Printf("%v%v\n", depthDelim, f.Name())
			listDir(dir+"/"+f.Name(), depthDelim+"| ", maxdepth, curDepth)

			//fmt.Printf("%v %v\n", f.Name(), fi.Mode().Perm())
		}
	}
	//return files
}

func listDir2(dir string) []fs.DirEntry {
	//var flist []string
	files, err := os.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	// for _, f := range files {
	// 	fm, err := os.Lstat(dir + "/" + f.Name())
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("%v  %v  %v\n", f.Name(), f.IsDir(), fm.Mode().String())

	// }
	return files
}

func listDir3(dir string) []filesStruct {
	var flist []filesStruct
	files, err := os.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		var ff filesStruct
		fm, err := os.Lstat(dir + "/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v  %v  %v\n", f.Name(), f.IsDir(), fm.Mode().String())
		ff.Name = f.Name()
		ff.Mode = fm.Mode().String()
		flist = append(flist, ff)

	}
	return flist
}
