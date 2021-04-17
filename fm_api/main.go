package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type author struct {
	Name          string
	FavouriteBook string
}

type filesStruct struct {
	Name string
	Mode string
}

var favouriteAuthors [6]author = [6]author{
	{Name: "N. K. Jemisin", FavouriteBook: "The Fifth Season"},
	{Name: "Noam Chomsky", FavouriteBook: "Who Rules the World?"},
	{Name: "Nnedi Okorafor", FavouriteBook: "Who Fears Death?"},
	{Name: "Robert Jordan", FavouriteBook: "The Dragon Reborn"},
	{Name: "Frederick Forsyth", FavouriteBook: "The Day of the Jackal"},
	{Name: "Octavia Butler", FavouriteBook: "Kindred"},
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/*", fs)

	http.HandleFunc("/", spaHandler)

	http.HandleFunc("/authors", authorHandler)
	http.HandleFunc("/l", dirHandler)

	http.ListenAndServe(":8080", nil)
}

func authorHandler(responseWriter http.ResponseWriter, request *http.Request) {
	favouriteAuthorsJSON, err := json.Marshal(favouriteAuthors)

	if err != nil {
		panic("Could not marshal json.")
	}

	fmt.Fprintf(responseWriter, string(favouriteAuthorsJSON))
}

func spaHandler(responseWriter http.ResponseWriter, request *http.Request) {
	http.ServeFile(responseWriter, request, "./static/index.html")
}

func dirHandler(responseWriter http.ResponseWriter, request *http.Request) {
	favouriteAuthorsJSON, err := json.Marshal(listDir("/home/mks/git"))

	if err != nil {
		panic("Could not marshal json.")
	}

	fmt.Fprintf(responseWriter, string(favouriteAuthorsJSON))
}

func listDir(dir string) []filesStruct {
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
