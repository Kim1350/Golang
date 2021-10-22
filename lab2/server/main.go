package main

import (
	"fmt"
	"net/http"
)

func main_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "api for teachers: /v1/teachers \napi for student: /v1/students")
}

func main() {
	http.HandleFunc("/", main_page)
	http.HandleFunc("/v1/teachers", teacherHandler)
	sHandler := studentHandler{}
	http.Handle("/v1/students", sHandler)
	http.ListenAndServe(":8080", nil)

}

func teacherHandler(res http.ResponseWriter, req *http.Request) {
	data := []byte("V1 of teacher's called")
	res.WriteHeader(200)
	res.Write(data)
}

type studentHandler struct{}

func (h studentHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("V1 of student's called")
	res.WriteHeader(200)
	res.Write(data)
}
