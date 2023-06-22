package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	fmt.Println("Server starting at port 8080")
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandlerFunc)
	err:= http.ListenAndServe(":8080",nil); 
	if err !=nil{
		log.Fatal(err)
	}
}

func helloHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL,r.Method,r.URL.Path,"helo")
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET"{
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello World")
}

func formHandler(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm();
	if err != nil{
		fmt.Println("Parse form error", err)
	}

	fmt.Fprintf(w, "POST request successful")

	name := r.FormValue("name")
	address:=r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}