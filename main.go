package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"log"
)

func index(w http.ResponseWriter, h *http.Request) {
	log.Println("Print Hello World")
	m := make(map[string]string, 1)
	m["data"] = "Hello World"
	d, err := json.Marshal(m)
	if err != nil {
		log.Println("解析错误", err.Error)
	}
	_, err = w.Write(d)
	if err != nil {
		log.Fatalln(err.Error)
	}

}

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe("localhost:8800", nil)
	fmt.Println(err)
}
