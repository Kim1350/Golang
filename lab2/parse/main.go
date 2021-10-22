package main

import (
	"fmt"
	"net/http"
)

type employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	createEmployeeHanlder := http.HandlerFunc(createEmployee)
	http.Handle("/employee", createEmployeeHanlder)
	http.ListenAndServe(":8080", nil)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	r.ParseForm()
	fmt.Println("request.Form::")
	for key, value := range r.Form {
		fmt.Printf("Key:%s, Value:%s\n", key, value)
	}
	fmt.Println("\nrequest.PostForm::")
	for key, value := range r.PostForm {
		fmt.Printf("Key:%s, Value:%s\n", key, value)
	}

	fmt.Printf("\nName field in Form:%s\n", r.Form["name"])
	fmt.Printf("\nName field in PostForm:%s\n", r.PostForm["name"])
	fmt.Printf("\nHobbies field in FormValue:%s\n", r.FormValue("hobbies"))

	w.WriteHeader(200)
	return
}
