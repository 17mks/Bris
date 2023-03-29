package main

import (
	"log"
	"net/http"
)

func main() {

	formModelPath := "D:\\tmp\\download"
	fs := http.FileServer(http.Dir(formModelPath))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println(":9300")
	err := http.ListenAndServe(":9300", nil)
	if err != nil {
		panic(err)
	}
	log.Println("end")
}
