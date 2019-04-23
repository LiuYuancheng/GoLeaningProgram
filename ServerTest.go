package main

import (
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		http.NotFound(w, r)
		return
	}
	switch r.Method{
	case "GET":
		w.Write([]byte("Receive a GET request\n"))
	case "POST":
		w.Write([]byte("Receive a POST requst\n"))
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)+"\n"))
	}
	// provide json fmt to the web
	switch r.Header.Get("Accept"){
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte("{\"Hello\", \"world\"}"))
	case "application/xml":
		w.Header().Set("Content-Type", "application/xml; charset=utf-8")
		w.Write([]byte("<?xml version=\"1.0\" encoding= \"utf-8\"?><Message>Hello World <Message>"))
	default:
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("Hello world\n"))
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//w.Write([]byte("Hello world \n"))
	w.Write([]byte("{\"hello\", \"world\"}"))
}

func userHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Please enter the user name:"))
}

func projectHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Please select the project"))
}

func main(){
	// build the routing list:
	http.HandleFunc("/", helloWorld)
	//http.HandleFunc("/users/", userHandler)
	//http.HandleFunc("/projects/", projectHandler)

	http.ListenAndServe(":8000", nil)
}