package main

import (
	"net/http"
	"io"
	"fmt"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*") 			
	w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET") 			
	w.Header().Add("Access-Control-Allow-Headers", "content-type, accept, accept-language, content-language") 
	if (req.Method == "OPTIONS") {
		w.WriteHeader(204)
		io.WriteString(w, "successful option")
	}	
	if req.Method == "GET" {
		resp, err := http.Get("https://localhost:8080/")
		if err != nil {
			fmt.Println(err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		sb := string(body)
		fmt.Printf(sb)
		io.WriteString(w, "successful get")
	} 
	if req.Method == "POST" {
		data, err := io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {return }
		fmt.Printf("%s\n", data)
		io.WriteString(w, "successful post")
	} else {
		w.WriteHeader(405)
	}	
}

func main() {
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":8080", nil)
	panic(err)
}
